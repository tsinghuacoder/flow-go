package data_providers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog"

	"github.com/onflow/flow-go/access"
	commonmodels "github.com/onflow/flow-go/engine/access/rest/common/models"
	commonparser "github.com/onflow/flow-go/engine/access/rest/common/parser"
	"github.com/onflow/flow-go/engine/access/rest/websockets/models"
	"github.com/onflow/flow-go/engine/access/subscription"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module/counters"

	"github.com/onflow/flow/protobuf/go/flow/entities"
)

// sendAndGetTransactionStatusesArguments contains the arguments required for sending tx and subscribing to transaction statuses
type sendAndGetTransactionStatusesArguments struct {
	Transaction flow.TransactionBody // The transaction body to be sent and monitored.
}

type SendAndGetTransactionStatusesDataProvider struct {
	*baseDataProvider

	logger        zerolog.Logger
	api           access.API
	linkGenerator commonmodels.LinkGenerator
}

var _ DataProvider = (*SendAndGetTransactionStatusesDataProvider)(nil)

func NewSendAndGetTransactionStatusesDataProvider(
	ctx context.Context,
	logger zerolog.Logger,
	api access.API,
	subscriptionID string,
	linkGenerator commonmodels.LinkGenerator,
	topic string,
	arguments models.Arguments,
	send chan<- interface{},
	chain flow.Chain,
) (*SendAndGetTransactionStatusesDataProvider, error) {
	p := &SendAndGetTransactionStatusesDataProvider{
		logger:        logger.With().Str("component", "send-transaction-statuses-data-provider").Logger(),
		api:           api,
		linkGenerator: linkGenerator,
	}

	// Initialize arguments passed to the provider.
	sendTxStatusesArgs, err := parseSendAndGetTransactionStatusesArguments(arguments, chain)
	if err != nil {
		return nil, fmt.Errorf("invalid arguments for send tx statuses data provider: %w", err)
	}

	subCtx, cancel := context.WithCancel(ctx)

	p.baseDataProvider = newBaseDataProvider(
		subscriptionID,
		topic,
		arguments,
		cancel,
		send,
		p.createSubscription(subCtx, sendTxStatusesArgs), // Set up a subscription to tx statuses based on arguments.
	)

	return p, nil
}

// Run starts processing the subscription for events and handles responses.
//
// No errors are expected during normal operations.
func (p *SendAndGetTransactionStatusesDataProvider) Run() error {
	messageIndex := counters.NewMonotonicCounter(0)

	return run(
		p.closedChan,
		p.subscription,
		func(response []*access.TransactionResult) error {
			return p.sendResponse(response, &messageIndex)
		},
	)
}

func (p *SendAndGetTransactionStatusesDataProvider) sendResponse(
	txResults []*access.TransactionResult,
	messageIndex *counters.StrictMonotonicCounter,
) error {
	for i := range txResults {
		var txStatusesPayload models.TransactionStatusesResponse
		txStatusesPayload.Build(p.linkGenerator, txResults[i], messageIndex.Value())

		var response models.BaseDataProvidersResponse
		response.Build(p.ID(), p.Topic(), &txStatusesPayload)

		messageIndex.Increment()
		p.send <- &response
	}

	return nil
}

// createSubscription creates a new subscription using the specified input arguments.
func (p *SendAndGetTransactionStatusesDataProvider) createSubscription(
	ctx context.Context,
	args sendAndGetTransactionStatusesArguments,
) subscription.Subscription {
	return p.api.SendAndSubscribeTransactionStatuses(ctx, &args.Transaction, entities.EventEncodingVersion_JSON_CDC_V0)
}

// parseSendAndGetTransactionStatusesArguments validates and initializes the account statuses arguments.
func parseSendAndGetTransactionStatusesArguments(
	arguments models.Arguments,
	chain flow.Chain,
) (sendAndGetTransactionStatusesArguments, error) {
	allowedFields := []string{
		"reference_block_id",
		"script",
		"arguments",
		"gas_limit",
		"payer",
		"proposal_key",
		"authorizers",
		"payload_signatures",
		"envelope_signatures",
	}
	err := ensureAllowedFields(arguments, allowedFields)
	if err != nil {
		return sendAndGetTransactionStatusesArguments{}, err
	}

	var args sendAndGetTransactionStatusesArguments

	// Convert the arguments map to JSON
	rawJSON, err := json.Marshal(arguments)
	if err != nil {
		return args, fmt.Errorf("failed to marshal arguments: %w", err)
	}

	// Create an io.Reader from the JSON bytes
	rawReader := bytes.NewReader(rawJSON)

	var tx commonparser.Transaction
	err = tx.Parse(rawReader, chain)
	if err != nil {
		return args, fmt.Errorf("failed to parse transaction: %w", err)
	}

	args.Transaction = tx.Flow()

	return args, nil
}
