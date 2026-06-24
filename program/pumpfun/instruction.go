package pumpfun

import (
	"encoding/binary"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/types"
)

func SellInstruction(param PumpFunSwapInstructionParam) types.Instruction {
	data := make([]byte, 24)
	binary.LittleEndian.PutUint64(data[0:8], SELL_OP_ID)
	binary.LittleEndian.PutUint64(data[8:16], param.AmountIn.BigInt().Uint64())
	binary.LittleEndian.PutUint64(data[16:24], param.MinAmountOut.BigInt().Uint64())

	return types.Instruction{common.PublicKeyFromString(PUMPFUN_AMM_PROGRAM_ADDRESS), param.Accounts, data}
}

func BuyInstruction(param PumpFunSwapInstructionParam) types.Instruction {
	data := make([]byte, 24)
	binary.LittleEndian.PutUint64(data[0:8], BUY_OP_ID)
	binary.LittleEndian.PutUint64(data[8:16], param.AmountIn.BigInt().Uint64())
	binary.LittleEndian.PutUint64(data[16:24], param.MinAmountOut.BigInt().Uint64())

	return types.Instruction{common.PublicKeyFromString(PUMPFUN_AMM_PROGRAM_ADDRESS), param.Accounts, data}
}
