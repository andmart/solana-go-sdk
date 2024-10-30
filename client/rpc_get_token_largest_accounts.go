package client

import (
	"context"

	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/rpc"

)

type TokenLargestAccounts struct {
	Address common.PublicKey
	Amount  float64
}

func (c *Client) GetTokenLargestAccounts(ctx context.Context, addr string) ([]TokenLargestAccounts, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.ValueWithContext[[]rpc.GetTokenLargestAccounts]], error) {
			return c.RpcClient.GetTokenLargestAccounts(ctx, addr)
		},
		convertGetTokenLargestAccounts,
	)
}

func convertGetTokenLargestAccounts(tla rpc.ValueWithContext[[]rpc.GetTokenLargestAccounts]) ([]TokenLargestAccounts, error) {
	var retval []TokenLargestAccounts
	for _, item := range tla.Value {
		t := TokenLargestAccounts{Address: item.Address, Amount: item.Amount}
		retval = append(retval, t)
	}
	return retval, nil
}
