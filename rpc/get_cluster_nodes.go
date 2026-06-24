package rpc

import "context"

type GetClusterNodesResponse JsonRpcResponse[GetClusterNodes]

type GetClusterNodes []GetClusterNode

type GetClusterNode struct {
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
func (c *RpcClient) GetClusterNodes(ctx context.Context) (JsonRpcResponse[GetClusterNodes], error) {
	return call[JsonRpcResponse[GetClusterNodes]](c, ctx, "getClusterNodes")
}
