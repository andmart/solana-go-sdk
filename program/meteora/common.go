package meteora

import (
	"github.com/blocto/solana-go-sdk/types"
	"github.com/shopspring/decimal"
)

const SWAP_OP_ID = 0xf8c69e91e17587c8

type MeteoraSwapInstructionParam struct {
	Accounts     []types.AccountMeta
	AmountIn     decimal.Decimal
	MinAmountOut decimal.Decimal
}
