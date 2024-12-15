package multi_client

import (
	"cosmossdk.io/depinject"
	cosmosclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/pokt-network/poktroll/pkg/client"
	"github.com/pokt-network/poktroll/pkg/client/query"
)

var _ MultiQueryClient = (*queryClient)(nil)

// TODO_IN_THIS_COMMIT: godoc...
func NewMultiQueryClient(deps depinject.Config, cometWebsocketURL string) (MultiQueryClient, error) {
	accountQuerier, err := query.NewAccountQuerier(deps)
	if err != nil {
		return nil, err
	}

	bankQuerier, err := query.NewBankQuerier(deps)
	if err != nil {
		return nil, err
	}

	blockQuerier, err := cosmosclient.NewClientFromNode(cometWebsocketURL)
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
//// TODO_IN_THIS_COMMIT: godoc...
//func (qc *queryClient) GetSharedParams(ctx context.Context) (*sharedtypes.Params, error) {
//	return qc.SharedQueryClient.GetParams(ctx)
//}
//
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
// func (qc *queryClient) GetApplicationParams(ctx context.Context) (*apptypes.Params, error) {
// 	return qc.ApplicationQueryClient.GetParams(ctx)
// }