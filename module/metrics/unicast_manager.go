package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/onflow/flow-go/module"
)

// UnicastManagerMetrics metrics collector for the unicast manager.
type UnicastManagerMetrics struct {
	// createStreamAttempts tracks the number of retry attempts to create a stream.
	createStreamAttempts *prometheus.HistogramVec
	// createStreamDuration tracks the overall time it takes to create a stream, this time includes
	// time spent dialing the peer and time spent connecting to the peer and creating the stream.
	createStreamDuration *prometheus.HistogramVec
	// dialPeerAttempts tracks the number of retry attempts to dial a peer during stream creation.
	dialPeerAttempts *prometheus.HistogramVec
	// dialPeerDuration tracks the time it takes to dial a peer and establish a connection.
	dialPeerDuration *prometheus.HistogramVec
	// establishStreamOnConnAttempts tracks the number of retry attempts to create the stream after peer dialing completes and a connection is established.
	establishStreamOnConnAttempts *prometheus.HistogramVec
	// establishStreamOnConnDuration tracks the time it takes to create the stream after peer dialing completes and a connection is established.
	establishStreamOnConnDuration *prometheus.HistogramVec

	prefix string
}

var _ module.UnicastManagerMetrics = (*UnicastManagerMetrics)(nil)

func NewUnicastManagerMetrics(prefix string) *UnicastManagerMetrics {
	uc := &UnicastManagerMetrics{prefix: prefix}

	uc.createStreamAttempts = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespaceNetwork,
			Subsystem: subsystemGossip,
			Name:      uc.prefix + "create_stream_attempts",
			Help:      "number of retry attempts before stream created successfully",
			Buckets:   []float64{1, 2, 3},
		}, []string{LabelSuccess},
	)

	uc.createStreamDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespaceNetwork,
			Subsystem: subsystemGossip,
			Name:      uc.prefix + "create_stream_duration",
			Help:      "the amount of time it takes to create a stream successfully",
			Buckets:   []float64{0.01, 0.1, 0.5, 1, 2, 5},
		}, []string{LabelSuccess},
	)

	uc.dialPeerAttempts = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespaceNetwork,
			Subsystem: subsystemGossip,
			Name:      uc.prefix + "dial_peer_attempts",
			Help:      "number of retry attempts before a peer is dialed successfully",
			Buckets:   []float64{1, 2, 3},
		}, []string{LabelSuccess},
	)

	uc.dialPeerDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespaceNetwork,
			Subsystem: subsystemGossip,
			Name:      uc.prefix + "dial_peer_duration",
			Help:      "the amount of time it takes to dial a peer during stream creation",
			Buckets:   []float64{0.01, 0.1, 0.5, 1, 2, 5},
		}, []string{LabelSuccess},
	)

	uc.establishStreamOnConnAttempts = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespaceNetwork,
			Subsystem: subsystemGossip,
			Name:      uc.prefix + "create_raw_stream_attempts",
			Help:      "number of retry attempts before a stream is created on the available connection between two peers",
			Buckets:   []float64{1, 2, 3},
		}, []string{LabelSuccess},
	)

	uc.establishStreamOnConnDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespaceNetwork,
			Subsystem: subsystemGossip,
			Name:      uc.prefix + "create_raw_stream_duration",
			Help:      "the amount of time it takes to create a stream on the available connection between two peers",
			Buckets:   []float64{0.01, 0.1, 0.5, 1, 2, 5},
		}, []string{LabelSuccess},
	)

	return uc
}

// OnStreamCreated tracks the overall time taken to create a stream successfully and the number of retry attempts.
func (u *UnicastManagerMetrics) OnStreamCreated(duration time.Duration, attempts int) {
	u.createStreamAttempts.WithLabelValues("true").Observe(float64(attempts))
	u.createStreamDuration.WithLabelValues("true").Observe(duration.Seconds())
}

// OnStreamCreationFailure tracks the overall time taken and number of retry attempts used when the unicast manager fails to create a stream.
func (u *UnicastManagerMetrics) OnStreamCreationFailure(duration time.Duration, attempts int) {
	u.createStreamAttempts.WithLabelValues("false").Observe(float64(attempts))
	u.createStreamDuration.WithLabelValues("false").Observe(duration.Seconds())
}

// OnPeerDialed tracks the time it takes to dial a peer during stream creation and the number of retry attempts before a peer
// is dialed successfully.
func (u *UnicastManagerMetrics) OnPeerDialed(duration time.Duration, attempts int) {
	u.dialPeerAttempts.WithLabelValues("true").Observe(float64(attempts))
	u.dialPeerDuration.WithLabelValues("true").Observe(duration.Seconds())
}

// OnPeerDialFailure tracks the amount of time taken and number of retry attempts used when the unicast manager cannot dial a peer
// to establish the initial connection between the two.
func (u *UnicastManagerMetrics) OnPeerDialFailure(duration time.Duration, attempts int) {
	u.dialPeerAttempts.WithLabelValues("false").Observe(float64(attempts))
	u.dialPeerDuration.WithLabelValues("false").Observe(duration.Seconds())
}

// OnStreamEstablished tracks the time it takes to create a stream successfully on the available open connection during stream
// creation and the number of retry attempts.
func (u *UnicastManagerMetrics) OnStreamEstablished(duration time.Duration, attempts int) {
	u.establishStreamOnConnAttempts.WithLabelValues("true").Observe(float64(attempts))
	u.establishStreamOnConnDuration.WithLabelValues("true").Observe(duration.Seconds())
}

// OnEstablishStreamFailure tracks the amount of time taken and number of retry attempts used when the unicast manager cannot establish
// a stream on the open connection between two peers.
func (u *UnicastManagerMetrics) OnEstablishStreamFailure(duration time.Duration, attempts int) {
	u.establishStreamOnConnAttempts.WithLabelValues("false").Observe(float64(attempts))
	u.establishStreamOnConnDuration.WithLabelValues("false").Observe(duration.Seconds())
}
