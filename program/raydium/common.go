package raydium

import (
	"github.com/blocto/solana-go-sdk/types"
	"github.com/shopspring/decimal"
)

type RaydiumSwapBaseInInstructionParam struct {
	Accounts     []types.AccountMeta
	AmountIn     decimal.Decimal
	MinAmountOut decimal.Decimal
}

type RaydiumSwapBaseOutInstructionParam struct {
	Accounts    []types.AccountMeta
	AmountOut   decimal.Decimal
	MaxAmountIn decimal.Decimal
}

type RaydiumCLLMSwapInstructionParam struct {
	Accounts            []types.AccountMeta
	Amount              decimal.Decimal
	OtherAmountThreshod decimal.Decimal
	SqrtPriceLimitX64   decimal.Decimal
	IsBaseInput         bool
	Version             int
}
