package pumpfun

import (
	"github.com/blocto/solana-go-sdk/types"
	"github.com/shopspring/decimal"
)

const (
	BUY_OP_ID                   = 16927863322537952870
	SELL_OP_ID                  = 12502976635542562355
	PUMPFUN_AMM_PROGRAM_ADDRESS = "pAMMBay6oceH9fJKBRHGP5D4bD4sWpmSwMn52FMfXEA"
)

type PumpFunSwapInstructionParam struct {
	Accounts     []types.AccountMeta
	AmountIn     decimal.Decimal
	MinAmountOut decimal.Decimal
}
