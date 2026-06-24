package client

import (
	"context"

	"github.com/blocto/solana-go-sdk/rpc"
)

type ClusterNode struct {
	FeatureSet      uint64 `json:"featureSet,omitempty"`
	Gossip          string `json:"gossip,omitempty"`
	Pubkey          string `json:"pubkey"`
	Pubsub          string `json:"pubsub,omitempty"`
	RPC             string `json:"rpc,omitempty"`
	ServeRepair     string `json:"serveRepair,omitempty"`
	ShredVersion    int    `json:"shredVersion,omitempty"`
	TPU             string `json:"tpu,omitempty"`
	TPUForwards     string `json:"tpuForwards,omitempty"`
	TPUForwardsQuic string `json:"tpuForwardsQuic,omitempty"`
	TPUQuic         string `json:"tpuQuic,omitempty"`
	TPUVote         string `json:"tpuVote,omitempty"`
	TVU             string `json:"tvu,omitempty"`
	Version         string `json:"version,omitempty"`
}

// GetClusterNodes returns information about all the nodes participating in the cluster
func (c *Client) GetClusterNodes(ctx context.Context) ([]ClusterNode, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.GetClusterNodes], error) {
			return c.RpcClient.GetClusterNodes(ctx)
		},
		convertGetClusterNodes,
	)
}

func convertGetClusterNodes(v rpc.GetClusterNodes) ([]ClusterNode, error) {
	output := make([]ClusterNode, 0, len(v))
	for _, info := range v {
		output = append(output, ClusterNode{
			FeatureSet:      info.FeatureSet,
			Gossip:          info.Gossip,
			Pubkey:          info.Pubkey,
			Pubsub:          info.Pubsub,
			RPC:             info.RPC,
			ServeRepair:     info.ServeRepair,
			ShredVersion:    info.ShredVersion,
			TPU:             info.TPU,
			TPUForwards:     info.TPUForwards,
			TPUForwardsQuic: info.TPUForwardsQuic,
			TPUQuic:         info.TPUQuic,
			TPUVote:         info.TPUVote,
			TVU:             info.TVU,
			Version:         info.Version,
		})
	}
	return output, nil
}
