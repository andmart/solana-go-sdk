package rpc

import (
	"context"

	"github.com/blocto/solana-go-sdk/common"
)

type GetTokenLargestAccounts struct {
	Address common.PublicKey `json:"address"`
	Amount  float64          `json:"uiAmount"`
}

// GetTransaction returns transaction details for a confirmed transaction
func (c *RpcClient) GetTokenLargestAccounts(ctx context.Context, base58Addr string) (JsonRpcResponse[ValueWithContext[[]GetTokenLargestAccounts]], error) {
	return call[JsonRpcResponse[ValueWithContext[[]GetTokenLargestAccounts]]](c, ctx, "getTokenLargestAccounts", base58Addr)
}
