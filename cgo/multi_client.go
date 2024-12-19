package main

import (
	"context"

	"cosmossdk.io/depinject"
	cosmosclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/pokt-network/poktroll/app"
	"github.com/pokt-network/poktroll/pkg/client"
	"github.com/pokt-network/poktroll/pkg/client/query"
	sharedtypes "github.com/pokt-network/poktroll/x/shared/types"
	"github.com/spf13/pflag"
)

var _ MultiQueryClient = (*queryClient)(nil)

// TODO_IN_THIS_COMMIT: godoc...
// Required dependencies:
//   - cosmosclient.Context (gogogrpc.ClientConn)
//   - client.BlockQueryClient
func NewMultiQueryClient(deps depinject.Config, queryNodeRPCURL string) (MultiQueryClient, error) {
	// TODO_IMPROVE: This should be parameterized.
	homedir := app.DefaultNodeHome
	clientCtx := cosmosclient.Context{}.
		WithCodec(cdc).
		WithTxConfig(TxConfig).
		WithHomeDir(homedir).
		WithAccountRetriever(authtypes.AccountRetriever{}).
		WithInterfaceRegistry(InterfaceRegistry)

	flagSet := pflag.NewFlagSet("test", pflag.ContinueOnError)
	// TODO_IMPROVE: It would be nice if the value could be set correctly based
	// on whether the test using it is running in tilt or not.
	flagSet.String(flags.FlagNode, queryNodeRPCURL, "")
	flagSet.String(flags.FlagHome, "", homedir)
	flagSet.String(flags.FlagChainID, app.Name, "use poktroll chain-id")
	err := flagSet.Parse([]string{})
	if err != nil {
		return nil, err
	}

	clientCtx, err = cosmosclient.ReadPersistentCommandFlags(clientCtx, flagSet)
	if err != nil {
		return nil, err
	}

	deps = depinject.Configs(deps, depinject.Supply(clientCtx))

	accountQuerier, err := query.NewAccountQuerier(deps)
	if err != nil {
		return nil, err
	}

	bankQuerier, err := query.NewBankQuerier(deps)
	if err != nil {
		return nil, err
	}

	blockQuerier, err := cosmosclient.NewClientFromNode(queryNodeRPCURL)
	if err != nil {
		return nil, err
	}

	sharedQuerier, err := query.NewSharedQuerier(deps)
	if err != nil {
		return nil, err
	}

	applicationQuerier, err := query.NewApplicationQuerier(deps)
	if err != nil {
		return nil, err
	}

	supplierQuerier, err := query.NewSupplierQuerier(deps)
	if err != nil {
		return nil, err
	}

	sessionQuerier, err := query.NewSessionQuerier(deps)
	if err != nil {
		return nil, err
	}

	serviceQuerier, err := query.NewServiceQuerier(deps)
	if err != nil {
		return nil, err
	}

	proofQuerier, err := query.NewProofQuerier(deps)
	if err != nil {
		return nil, err
	}

	// TODO_OPTIMIZE: lazily initialize these, so that they're only constructed when needed.
	return &queryClient{
		AccountQueryClient:     accountQuerier,
		BankQueryClient:        bankQuerier,
		BlockQueryClient:       blockQuerier,
		SharedQueryClient:      sharedQuerier,
		ApplicationQueryClient: applicationQuerier,
		SupplierQueryClient:    supplierQuerier,
		SessionQueryClient:     sessionQuerier,
		ServiceQueryClient:     serviceQuerier,
		ProofQueryClient:       proofQuerier,
	}, nil
}

// TODO_IN_THIS_COMMIT: godoc...
type queryClient struct {
	client.AccountQueryClient
	client.BankQueryClient
	client.BlockQueryClient
	client.SharedQueryClient
	client.ApplicationQueryClient
	client.SupplierQueryClient
	client.SessionQueryClient
	client.ServiceQueryClient
	client.ProofQueryClient
}

// TODO_BLOCKED(@bryanchriswhite poktroll#543): add once available.
//
// TODO_IN_THIS_COMMIT: godoc...
func (qc *queryClient) GetSharedParams(ctx context.Context) (*sharedtypes.Params, error) {
	return qc.SharedQueryClient.GetParams(ctx)
}

//// TODO_IN_THIS_COMMIT: godoc...
//func (qc *queryClient) GetSupplierParams(ctx context.Context) (*suppliertypes.Params, error) {
//	return qc.SupplierQueryClient.GetParams(ctx)
//}
//
//// TODO_IN_THIS_COMMIT: godoc...
//func (qc *queryClient) GetServiceParams(ctx context.Context) (*servicetypes.Params, error) {
//	return qc.ServiceQueryClient.GetParams(ctx)
//}
//
//// TODO_IN_THIS_COMMIT: godoc...
//func (qc *queryClient) GetSessionParams(ctx context.Context) (*sessiontypes.Params, error) {
//	return qc.SessionQueryClient.GetParams(ctx)
//}
//
//// TODO_IN_THIS_COMMIT: godoc...
//func (qc *queryClient) GetProofParams(ctx context.Context) (*prooftypes.Params, error) {
//	params, err := qc.ProofQueryClient.GetParams(ctx)
//	return params.(*prooftypes.Params), err
//}
//
//// TODO_IN_THIS_COMMIT: godoc...
//func (qc *queryClient) GetApplicationParams(ctx context.Context) (*apptypes.Params, error) {
//	return qc.ApplicationQueryClient.GetParams(ctx)
//}
