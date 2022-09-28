package fvm

import (
	"github.com/onflow/cadence/runtime/common"
	"github.com/onflow/cadence/runtime/interpreter"

	"github.com/onflow/flow-go/fvm/environment"
	"github.com/onflow/flow-go/fvm/programs"
	"github.com/onflow/flow-go/fvm/state"
	"github.com/onflow/flow-go/model/flow"
)

// TODO(patrick): rm after emulator is updated
type Environment = environment.Environment

type EnvironmentParams struct {
	Chain flow.Chain

	// NOTE: The ServiceAccountEnabled option is used by the playground
	// https://github.com/onflow/flow-playground-api/blob/1ad967055f31db8f1ce88e008960e5fc14a9fbd1/compute/computer.go#L76
	ServiceAccountEnabled bool

	environment.RuntimeParams

	environment.TracerParams
	environment.ProgramLoggerParams

	environment.EventEmitterParams

	environment.BlockInfoParams
	environment.TransactionInfoParams

	environment.ContractUpdaterParams
}

func DefaultEnvironmentParams() EnvironmentParams {
	return EnvironmentParams{
		Chain:                 flow.Mainnet.Chain(),
		ServiceAccountEnabled: true,

		RuntimeParams:         environment.DefaultRuntimeParams(),
		TracerParams:          environment.DefaultTracerParams(),
		ProgramLoggerParams:   environment.DefaultProgramLoggerParams(),
		EventEmitterParams:    environment.DefaultEventEmitterParams(),
		BlockInfoParams:       environment.DefaultBlockInfoParams(),
		TransactionInfoParams: environment.DefaultTransactionInfoParams(),
		ContractUpdaterParams: environment.DefaultContractUpdaterParams(),
	}
}

var _ environment.Environment = &facadeEnvironment{}

// facadeEnvironment exposes various fvm business logic as a single interface.
type facadeEnvironment struct {
	*environment.Runtime

	*environment.Tracer
	environment.Meter

	*environment.ProgramLogger
	environment.EventEmitter

	*environment.UnsafeRandomGenerator
	*environment.CryptoLibrary

	*environment.BlockInfo
	*environment.AccountInfo
	environment.TransactionInfo

	*environment.ValueStore

	*environment.SystemContracts

	*environment.UUIDGenerator

	environment.AccountCreator
	environment.AccountFreezer

	*environment.AccountKeyReader
	environment.AccountKeyUpdater

	*environment.ContractReader
	environment.ContractUpdater
	*environment.Programs

	accounts environment.Accounts
}

func newFacadeEnvironment(
	ctx Context,
	txnState *state.TransactionState,
	programs environment.TransactionPrograms,
	tracer *environment.Tracer,
	meter environment.Meter,
) *facadeEnvironment {
	accounts := environment.NewAccounts(txnState)
	logger := environment.NewProgramLogger(tracer, ctx.ProgramLoggerParams)
	runtime := environment.NewRuntime(ctx.RuntimeParams)
	systemContracts := environment.NewSystemContracts(
		ctx.Chain,
		tracer,
		logger,
		runtime)

	env := &facadeEnvironment{
		Runtime: runtime,

		Tracer: tracer,
		Meter:  meter,

		ProgramLogger: logger,
		EventEmitter:  environment.NoEventEmitter{},

		UnsafeRandomGenerator: environment.NewUnsafeRandomGenerator(
			tracer,
			ctx.BlockHeader,
		),
		CryptoLibrary: environment.NewCryptoLibrary(tracer, meter),

		BlockInfo: environment.NewBlockInfo(
			tracer,
			meter,
			ctx.BlockHeader,
			ctx.Blocks,
		),
		AccountInfo: environment.NewAccountInfo(
			tracer,
			meter,
			accounts,
			systemContracts,
			ctx.ServiceAccountEnabled,
		),
		TransactionInfo: environment.NoTransactionInfo{},

		ValueStore: environment.NewValueStore(
			tracer,
			meter,
			accounts,
		),

		SystemContracts: systemContracts,

		UUIDGenerator: environment.NewUUIDGenerator(
			tracer,
			meter,
			txnState),

		AccountCreator: environment.NoAccountCreator{},
		AccountFreezer: environment.NoAccountFreezer{},

		AccountKeyReader: environment.NewAccountKeyReader(
			tracer,
			meter,
			accounts,
		),
		AccountKeyUpdater: environment.NoAccountKeyUpdater{},

		ContractReader: environment.NewContractReader(
			tracer,
			meter,
			accounts,
		),
		ContractUpdater: environment.NoContractUpdater{},
		Programs: environment.NewPrograms(
			tracer,
			meter,
			txnState,
			accounts,
			programs),

		accounts: accounts,
	}

	env.Runtime.SetEnvironment(env)

	return env
}

func (env *facadeEnvironment) FlushPendingUpdates() (
	programs.ModifiedSetsInvalidator,
	error,
) {
	keys, err := env.ContractUpdater.Commit()
	return programs.ModifiedSetsInvalidator{
		ContractUpdateKeys: keys,
		FrozenAccounts:     env.FrozenAccounts(),
	}, err
}

func (env *facadeEnvironment) Reset() {
	env.ContractUpdater.Reset()
	env.EventEmitter.Reset()
	env.AccountFreezer.Reset()
}

// Miscellaneous cadence runtime.Interface API.
func (facadeEnvironment) ResourceOwnerChanged(
	*interpreter.Interpreter,
	*interpreter.CompositeValue,
	common.Address,
	common.Address,
) {
}
