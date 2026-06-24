package rpc

import (
	"context"
	"encoding/json"
)

type GetTokenAccountsByOwnerResponse JsonRpcResponse[GetTokenAccountsByOwner]

type GetTokenAccountsByOwner ValueWithContext[GetProgramAccounts]

// GetTokenAccountsByOwnerConfig is a option config for `GetTokenAccountsByOwnerWithFilter`
type GetTokenAccountsByOwnerConfig struct {
	Commitment Commitment      `json:"commitment,omitempty"`
	Encoding   AccountEncoding `json:"encoding,omitempty"`
	DataSlice  *DataSlice      `json:"dataSlice,omitempty"`
}

// GetTokenAccountsByOwnerConfigFilter either mint or programId
type GetTokenAccountsByOwnerConfigFilter struct {
	Mint      string `json:"mint,omitempty"`
	ProgramId string `json:"programId,omitempty"`
}

type TokenAccountFlatResult struct {
	Value []TokenAccountFlat `json:"value"`
}

type TokenAccountFlat struct {
	Mint           string  `json:"mint"`
	Owner          string  `json:"owner"`
	IsNative       bool    `json:"isNative"`
	State          string  `json:"state"`
	Amount         string  `json:"amount"`
	Decimals       int     `json:"decimals"`
	UiAmount       float64 `json:"uiAmount"`
	UiAmountString string  `json:"uiAmountString"`
	Pubkey         string  `json:"pubkey"`
}

func (fv *TokenAccountFlat) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	account := raw["account"].(map[string]interface{})
	dataField := account["data"].(map[string]interface{})
	parsed := dataField["parsed"].(map[string]interface{})
	info := parsed["info"].(map[string]interface{})
	tokenAmount := info["tokenAmount"].(map[string]interface{})

	fv.Mint = info["mint"].(string)
	fv.Owner = info["owner"].(string)
	fv.IsNative = info["isNative"].(bool)
	fv.State = info["state"].(string)
	fv.Amount = tokenAmount["amount"].(string)
	fv.Decimals = int(tokenAmount["decimals"].(float64))
	fv.UiAmount = tokenAmount["uiAmount"].(float64)
	fv.UiAmountString = tokenAmount["uiAmountString"].(string)
	fv.Pubkey = raw["pubkey"].(string)

	return nil
}

func (c *RpcClient) GetTokenAccountsByOwner(ctx context.Context, base58Addr string) (JsonRpcResponse[ValueWithContext[GetProgramAccounts]], error) {
	return call[JsonRpcResponse[ValueWithContext[GetProgramAccounts]]](c, ctx, "getTokenAccountsByOwner", base58Addr)
}

func (c *RpcClient) GetTokenAccountsByOwnerWithFilter(ctx context.Context, base58Addr string, filter GetTokenAccountsByOwnerConfigFilter) (JsonRpcResponse[ValueWithContext[GetProgramAccounts]], error) {
	return call[JsonRpcResponse[ValueWithContext[GetProgramAccounts]]](c, ctx, "getTokenAccountsByOwner", base58Addr, filter)
}

func (c *RpcClient) GetTokenAccountsByOwnerWithConfig(ctx context.Context, base58Addr string, filter GetTokenAccountsByOwnerConfigFilter, cfg GetTokenAccountsByOwnerConfig) (JsonRpcResponse[ValueWithContext[GetProgramAccounts]], error) {
	return call[JsonRpcResponse[ValueWithContext[GetProgramAccounts]]](c, ctx, "getTokenAccountsByOwner", base58Addr, filter, cfg)
}

func (c *RpcClient) GetAllTokenAccountsByOwnerWithConfig(ctx context.Context, base58Addr string, filter GetTokenAccountsByOwnerConfigFilter, cfg GetTokenAccountsByOwnerConfig) (JsonRpcResponse[ValueWithContext[[]TokenAccountFlat]], error) {
	return call[JsonRpcResponse[ValueWithContext[[]TokenAccountFlat]]](c, ctx, "getTokenAccountsByOwner", base58Addr, filter, cfg)
}
