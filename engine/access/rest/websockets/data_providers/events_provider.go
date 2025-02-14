package data_providers

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"

	"github.com/onflow/flow-go/engine/access/rest/common"
	"github.com/onflow/flow-go/engine/access/rest/common/parser"
	"github.com/onflow/flow-go/engine/access/rest/http/request"
	"github.com/onflow/flow-go/engine/access/rest/websockets/data_providers/models"
	wsmodels "github.com/onflow/flow-go/engine/access/rest/websockets/models"
	"github.com/onflow/flow-go/engine/access/state_stream"
	"github.com/onflow/flow-go/engine/access/state_stream/backend"
	"github.com/onflow/flow-go/engine/access/subscription"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module/counters"
)

// eventsArguments contains the arguments required for subscribing to events
type eventsArguments struct {
	StartBlockID     flow.Identifier          // ID of the block to start subscription from
	StartBlockHeight uint64                   // Height of the block to start subscription from
	Filter           state_stream.EventFilter // Filter applied to events for a given subscription
}

// EventsDataProvider is responsible for providing events
type EventsDataProvider struct {
	*baseDataProvider

	logger         zerolog.Logger
	stateStreamApi state_stream.API

	heartbeatInterval uint64
}

var _ DataProvider = (*EventsDataProvider)(nil)

// NewEventsDataProvider creates a new instance of EventsDataProvider.
func NewEventsDataProvider(
	ctx context.Context,
	logger zerolog.Logger,
	stateStreamApi state_stream.API,
	subscriptionID string,
	topic string,
	arguments wsmodels.Arguments,
	send chan<- interface{},
	chain flow.Chain,
	eventFilterConfig state_stream.EventFilterConfig,
	heartbeatInterval uint64,
) (*EventsDataProvider, error) {
	p := &EventsDataProvider{
		logger:            logger.With().Str("component", "events-data-provider").Logger(),
		stateStreamApi:    stateStreamApi,
		heartbeatInterval: heartbeatInterval,
	}

	if p.stateStreamApi == nil {
		logger.Debug().Msg("state stream api is nil. cannot stream events")
		return nil, fmt.Errorf("this acess node is configured with no state stream api. cannot stream events")
	}

	// Initialize arguments passed to the provider.
	eventArgs, err := parseEventsArguments(arguments, chain, eventFilterConfig)
	if err != nil {
		return nil, fmt.Errorf("invalid arguments for events data provider: %w", err)
	}

	subCtx, cancel := context.WithCancel(ctx)

	p.baseDataProvider = newBaseDataProvider(
		subscriptionID,
		topic,
		arguments,
		cancel,
		send,
		p.createSubscription(subCtx, eventArgs), // Set up a subscription to events based on arguments.
	)

	return p, nil
}

// Run starts processing the subscription for events and handles responses.
//
// Expected errors during normal operations:
//   - context.Canceled: if the operation is canceled, during an unsubscribe action.
func (p *EventsDataProvider) Run() error {
	return subscription.HandleSubscription(p.subscription, p.handleResponse())
}

// handleResponse processes events and sends the formatted response.
//
// No errors are expected during normal operations.
func (p *EventsDataProvider) handleResponse() func(eventsResponse *backend.EventsResponse) error {
	blocksSinceLastMessage := uint64(0)
	messageIndex := counters.NewMonotonicCounter(0)

	return func(eventsResponse *backend.EventsResponse) error {
		// check if there are any events in the response. if not, do not send a message unless the last
		// response was more than HeartbeatInterval blocks ago
		if len(eventsResponse.Events) == 0 {
			blocksSinceLastMessage++
			if blocksSinceLastMessage < p.heartbeatInterval {
				return nil
			}
		}
		blocksSinceLastMessage = 0

		index := messageIndex.Value()
		if ok := messageIndex.Set(messageIndex.Value() + 1); !ok {
			return fmt.Errorf("message index already incremented to: %d", messageIndex.Value())
		}

		eventsPayload := models.NewEventResponse(eventsResponse, index)
		response := models.BaseDataProvidersResponse{
			SubscriptionID: p.ID(),
			Topic:          p.Topic(),
			Payload:        &eventsPayload,
		}

		p.send <- &response

		return nil
	}
}

// createSubscription creates a new subscription using the specified input arguments.
func (p *EventsDataProvider) createSubscription(ctx context.Context, args eventsArguments) subscription.Subscription {
	if args.StartBlockID != flow.ZeroID {
		return p.stateStreamApi.SubscribeEventsFromStartBlockID(ctx, args.StartBlockID, args.Filter)
	}

	if args.StartBlockHeight != request.EmptyHeight {
		return p.stateStreamApi.SubscribeEventsFromStartHeight(ctx, args.StartBlockHeight, args.Filter)
	}

	return p.stateStreamApi.SubscribeEventsFromLatest(ctx, args.Filter)
}

// parseEventsArguments validates and initializes the events arguments.
func parseEventsArguments(
	arguments wsmodels.Arguments,
	chain flow.Chain,
	eventFilterConfig state_stream.EventFilterConfig,
) (eventsArguments, error) {
	var args eventsArguments

	// Parse block arguments
	startBlockID, startBlockHeight, err := ParseStartBlock(arguments)
	if err != nil {
		return args, err
	}
	args.StartBlockID = startBlockID
	args.StartBlockHeight = startBlockHeight

	// Parse 'event_types' as a JSON array
	var eventTypes parser.EventTypes
	if eventTypesIn, ok := arguments["event_types"]; ok && eventTypesIn != "" {
		result, err := common.ParseInterfaceToString(eventTypesIn)
		if err != nil {
			return args, fmt.Errorf("'event_types' must be an array of string")
		}

		err = eventTypes.Parse(result)
		if err != nil {
			return args, fmt.Errorf("invalid 'event_types': %w", err)
		}
	}

	// Parse 'addresses' as []string{}
	var addresses []string
	if addressesIn, ok := arguments["addresses"]; ok && addressesIn != "" {
		addresses, err = common.ParseInterfaceToString(addressesIn)
		if err != nil {
			return args, fmt.Errorf("'addresses' must be an array of string")
		}
	}

	// Parse 'contracts' as []string{}
	var contracts []string
	if contractsIn, ok := arguments["contracts"]; ok && contractsIn != "" {
		contracts, err = common.ParseInterfaceToString(contractsIn)
		if err != nil {
			return args, fmt.Errorf("'contracts' must be an array of string")
		}
	}

	// Initialize the event filter with the parsed arguments
	args.Filter, err = state_stream.NewEventFilter(eventFilterConfig, chain, eventTypes.Flow(), addresses, contracts)
	if err != nil {
		return args, fmt.Errorf("failed to create event filter: %w", err)
	}

	return args, nil
}
