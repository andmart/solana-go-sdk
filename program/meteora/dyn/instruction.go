package dyn

import (
	"encoding/binary"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/meteora"
	"github.com/blocto/solana-go-sdk/types"
)

const (
	METEORA_DYN_PROGRAM_ADDRESS = "Eo7WjKq67rjJQSZxS6z3YkapzY3eMj6Xy8X5EQVn5UaB"
)

func SwapBaseInInstruction(param meteora.MeteoraSwapInstructionParam) types.Instruction {
	data := make([]byte, 24)
	binary.LittleEndian.PutUint64(data[0:8], meteora.SWAP_OP_ID)
	binary.LittleEndian.PutUint64(data[8:16], param.AmountIn.BigInt().Uint64())
	binary.LittleEndian.PutUint64(data[16:24], param.MinAmountOut.BigInt().Uint64())

	return types.Instruction{common.PublicKeyFromString(METEORA_DYN_PROGRAM_ADDRESS), param.Accounts, data}
}
