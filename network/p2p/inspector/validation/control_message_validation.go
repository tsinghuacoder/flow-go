package validation

import (
	"fmt"
	"math/rand"

	"github.com/hashicorp/go-multierror"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	pubsub_pb "github.com/libp2p/go-libp2p-pubsub/pb"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/rs/zerolog"

	"github.com/onflow/flow-go/engine/common/worker"
	"github.com/onflow/flow-go/module/component"
	"github.com/onflow/flow-go/module/irrecoverable"
	"github.com/onflow/flow-go/module/mempool/queue"
	"github.com/onflow/flow-go/module/metrics"
	"github.com/onflow/flow-go/network/channels"
	"github.com/onflow/flow-go/network/p2p"
	"github.com/onflow/flow-go/utils/logging"
)

const (
	// DefaultNumberOfWorkers default number of workers for the inspector component.
	DefaultNumberOfWorkers = 5
	// DefaultControlMsgValidationInspectorQueueCacheSize is the default size of the inspect message queue.
	DefaultControlMsgValidationInspectorQueueCacheSize = 100
)

// InspectMsgReq represents a short digest of an RPC control message. It is used for further message inspection by component workers.
type InspectMsgReq struct {
	// Nonce adds random value so that when msg req is stored on hero store a unique ID can be created from the struct fields.
	Nonce uint64
	// Peer sender of the message.
	Peer peer.ID
	// TopicIDS list of topic IDs in the control message.
	TopicIDS []string
	// Count the amount of control messages.
	Count            uint64
	validationConfig *CtrlMsgValidationConfig
}

// ControlMsgValidationInspectorConfig validation configuration for each type of RPC control message.
type ControlMsgValidationInspectorConfig struct {
	// NumberOfWorkers number of component workers to start for processing RPC messages.
	NumberOfWorkers int
	// InspectMsgStoreOpts options used to configure the underlying herocache message store.
	InspectMsgStoreOpts []queue.HeroStoreConfigOption
	// GraftValidationCfg validation configuration for GRAFT control messages.
	GraftValidationCfg *CtrlMsgValidationConfig
	// PruneValidationCfg validation configuration for PRUNE control messages.
	PruneValidationCfg *CtrlMsgValidationConfig
}

func (conf *ControlMsgValidationInspectorConfig) config(controlMsg p2p.ControlMessageType) (*CtrlMsgValidationConfig, bool) {
	switch controlMsg {
	case p2p.CtrlMsgGraft:
		return conf.GraftValidationCfg, true
	case p2p.CtrlMsgPrune:
		return conf.PruneValidationCfg, true
	default:
		return nil, false
	}
}

// configs returns all control message validation configs in a list.
func (conf *ControlMsgValidationInspectorConfig) configs() CtrlMsgValidationConfigs {
	return CtrlMsgValidationConfigs{conf.GraftValidationCfg, conf.PruneValidationCfg}
}

// ControlMsgValidationInspector RPC message inspector that inspects control messages and performs some validation on them,
// when some validation rule is broken feedback is given via the Peer scoring notifier.
type ControlMsgValidationInspector struct {
	component.Component
	logger zerolog.Logger
	// config control message validation configurations.
	config *ControlMsgValidationInspectorConfig
	// distributor used to disseminate invalid RPC message notifications.
	distributor p2p.GossipSubInspectorNotificationDistributor
	// workerPool queue that stores *InspectMsgReq that will be processed by component workers.
	workerPool *worker.Pool[*InspectMsgReq]
}

var _ component.Component = (*ControlMsgValidationInspector)(nil)

// NewInspectMsgReq returns a new *InspectMsgReq.
func NewInspectMsgReq(from peer.ID, validationConfig *CtrlMsgValidationConfig, topicIDS []string, count uint64) *InspectMsgReq {
	return &InspectMsgReq{Nonce: rand.Uint64(), Peer: from, validationConfig: validationConfig, TopicIDS: topicIDS, Count: count}
}

// NewControlMsgValidationInspector returns new ControlMsgValidationInspector
func NewControlMsgValidationInspector(
	logger zerolog.Logger,
	config *ControlMsgValidationInspectorConfig,
	distributor p2p.GossipSubInspectorNotificationDistributor,
) *ControlMsgValidationInspector {
	lg := logger.With().Str("component", "gossip_sub_rpc_validation_inspector").Logger()
	c := &ControlMsgValidationInspector{
		logger:      lg,
		config:      config,
		distributor: distributor,
	}

	cfg := &queue.HeroStoreConfig{
		SizeLimit: DefaultControlMsgValidationInspectorQueueCacheSize,
		Collector: metrics.NewNoopCollector(),
	}

	for _, opt := range config.InspectMsgStoreOpts {
		opt(cfg)
	}

	store := queue.NewHeroStore(cfg.SizeLimit, logger, cfg.Collector)
	pool := worker.NewWorkerPoolBuilder[*InspectMsgReq](lg, store, c.processInspectMsgReq).Build()

	c.workerPool = pool

	builder := component.NewComponentManagerBuilder()
	// start rate limiters cleanup loop in workers
	for _, conf := range c.config.configs() {
		builder.AddWorker(func(ctx irrecoverable.SignalerContext, ready component.ReadyFunc) {
			ready()
			conf.RateLimiter.CleanupLoop(ctx)
		})
	}
	for i := 0; i < c.config.NumberOfWorkers; i++ {
		builder.AddWorker(pool.WorkerLogic())
	}
	c.Component = builder.Build()
	return c
}

// Inspect inspects the rpc received and returns an error if any validation rule is broken.
// For each control message type an initial inspection is done synchronously to check the amount
// of messages in the control message. Further inspection is done asynchronously to check rate limits
// and validate topic IDS each control message if initial validation is passed.
// All errors returned from this function can be considered benign.
func (c *ControlMsgValidationInspector) Inspect(from peer.ID, rpc *pubsub.RPC) error {
	control := rpc.GetControl()

	err := c.inspect(from, p2p.CtrlMsgGraft, control)
	if err != nil {
		return fmt.Errorf("validation failed for control message %s: %w", p2p.CtrlMsgGraft, err)
	}

	err = c.inspect(from, p2p.CtrlMsgPrune, control)
	if err != nil {
		return fmt.Errorf("validation failed for control message %s: %w", p2p.CtrlMsgPrune, err)
	}

	return nil
}

// inspect performs initial inspection of RPC control message and queues up message for further inspection if required.
// All errors returned from this function can be considered benign.
// errors returned:
//
//	ErrUpperThreshold if message Count greater than the configured upper threshold.
func (c *ControlMsgValidationInspector) inspect(from peer.ID, ctrlMsgType p2p.ControlMessageType, ctrlMsg *pubsub_pb.ControlMessage) error {
	validationConfig, ok := c.config.config(ctrlMsgType)
	if !ok {
		return fmt.Errorf("failed to get validation configuration for control message %s", ctrlMsg)
	}
	count, topicIDS := c.getCtrlMsgData(ctrlMsgType, ctrlMsg)
	lg := c.logger.With().
		Str("peer_id", from.String()).
		Str("ctrl_msg_type", string(ctrlMsgType)).
		Uint64("ctrl_msg_count", count).Logger()

	// if Count greater than upper threshold drop message and penalize
	if count > validationConfig.UpperThreshold {
		upperThresholdErr := NewUpperThresholdErr(validationConfig.ControlMsg, count, validationConfig.UpperThreshold)
		lg.Warn().
			Err(upperThresholdErr).
			Uint64("upper_threshold", upperThresholdErr.upperThreshold).
			Bool(logging.KeySuspicious, true).
			Msg("rejecting rpc message")

		err := c.distributor.DistributeInvalidControlMessageNotification(p2p.NewInvalidControlMessageNotification(from, ctrlMsgType, count, upperThresholdErr))
		if err != nil {
			lg.Error().
				Err(err).
				Bool(logging.KeySuspicious, true).
				Msg("failed to distribute invalid control message notification")
			return err
		}
		return upperThresholdErr
	}
	// queue further async inspection
	c.requestMsgInspection(NewInspectMsgReq(from, validationConfig, topicIDS, count))
	return nil
}

// processInspectMsgReq func used by component workers to perform further inspection of control messages that will check if the messages are rate limited
// and ensure all topic IDS are valid when the amount of messages is above the configured safety threshold.
func (c *ControlMsgValidationInspector) processInspectMsgReq(req *InspectMsgReq) error {
	lg := c.logger.With().
		Str("peer_id", req.Peer.String()).
		Str("ctrl_msg_type", string(req.validationConfig.ControlMsg)).
		Uint64("ctrl_msg_count", req.Count).Logger()
	var validationErr error
	switch {
	case !req.validationConfig.RateLimiter.Allow(req.Peer, int(req.Count)): // check if Peer RPC messages are rate limited
		validationErr = NewRateLimitedControlMsgErr(req.validationConfig.ControlMsg)
	case req.Count > req.validationConfig.SafetyThreshold: // check if Peer RPC messages Count greater than safety threshold further inspect each message individually
		validationErr = c.validateTopics(req.validationConfig.ControlMsg, req.TopicIDS)
	default:
		lg.Trace().
			Uint64("upper_threshold", req.validationConfig.UpperThreshold).
			Uint64("safety_threshold", req.validationConfig.SafetyThreshold).
			Msg(fmt.Sprintf("control message %s inspection passed %d is below configured safety threshold", req.validationConfig.ControlMsg, req.Count))
		return nil
	}
	if validationErr != nil {
		lg.Error().
			Err(validationErr).
			Bool(logging.KeySuspicious, true).
			Msg(fmt.Sprintf("rpc control message async inspection failed"))
		err := c.distributor.DistributeInvalidControlMessageNotification(p2p.NewInvalidControlMessageNotification(req.Peer, req.validationConfig.ControlMsg, req.Count, validationErr))
		if err != nil {
			lg.Error().
				Err(err).
				Bool(logging.KeySuspicious, true).
				Msg("failed to distribute invalid control message notification")
		}
	}
	return nil
}

// requestMsgInspection queues up an inspect message request.
func (c *ControlMsgValidationInspector) requestMsgInspection(req *InspectMsgReq) {
	c.workerPool.Submit(req)
}

// getCtrlMsgData returns the amount of specified control message type in the rpc ControlMessage as well as the topic ID for each message.
func (c *ControlMsgValidationInspector) getCtrlMsgData(ctrlMsgType p2p.ControlMessageType, ctrlMsg *pubsub_pb.ControlMessage) (uint64, []string) {
	topicIDS := make([]string, 0)
	count := 0
	switch ctrlMsgType {
	case p2p.CtrlMsgGraft:
		grafts := ctrlMsg.GetGraft()
		for _, graft := range grafts {
			topicIDS = append(topicIDS, graft.GetTopicID())
		}
		count = len(grafts)
	case p2p.CtrlMsgPrune:
		prunes := ctrlMsg.GetPrune()
		for _, prune := range prunes {
			topicIDS = append(topicIDS, prune.GetTopicID())
		}
		count = len(prunes)
	}
	return uint64(count), topicIDS
}

// validateTopics ensures the topic is a valid flow topic/channel and the node has a subscription to that topic.
// All errors returned from this function can be considered benign.
func (c *ControlMsgValidationInspector) validateTopics(ctrlMsg p2p.ControlMessageType, topics []string) error {
	var errs *multierror.Error
	for _, t := range topics {
		topic := channels.Topic(t)
		channel, ok := channels.ChannelFromTopic(topic)
		if !ok {
			errs = multierror.Append(errs, NewMalformedTopicErr(ctrlMsg, topic))
			continue
		}

		if !channels.ChannelExists(channel) {
			errs = multierror.Append(errs, NewUnknownTopicChannelErr(ctrlMsg, topic))
		}
	}
	return errs.ErrorOrNil()
}
