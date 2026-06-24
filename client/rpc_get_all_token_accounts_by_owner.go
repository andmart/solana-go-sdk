package client

import (
	"context"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/rpc"
)

func (c *Client) GetAllTokenAccountsByOwner(ctx context.Context, owner string) ([]rpc.TokenAccountFlat, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[[]rpc.TokenAccountFlat]], error) {
			return c.RpcClient.GetAllTokenAccountsByOwnerWithConfig(
				ctx,
				owner,
				rpc.GetTokenAccountsByOwnerConfigFilter{
					ProgramId: common.TokenProgramID.String(),
				},
				rpc.GetTokenAccountsByOwnerConfig{
					Encoding: rpc.AccountEncodingJsonParsed,
				},
			)
		},
		convertGetAllTokenAccountsByOwner,
	)
}

func convertGetAllTokenAccountsByOwner(v rpc.ValueWithContext[[]rpc.TokenAccountFlat]) ([]rpc.TokenAccountFlat, error) {
	return v.Value, nil
}
