package client

import (
	"context"

	"github.com/blocto/solana-go-sdk/rpc"
)

// GetBlockTime returns the estimated production time of a block.
func (c *Client) GetBlockHeight(ctx context.Context) (uint64, error) {
	return process(
		func() (rpc.JsonRpcResponse[uint64], error) {
			return c.RpcClient.GetBlockHeight(ctx)
		},
		forward[uint64],
	)
}
