package amm

import (
	"encoding/binary"

	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/raydium"
	"github.com/blocto/solana-go-sdk/types"
)

const (
	RAYDIUM_AMM_PROGRAM_ADDRESS = "675kPX9MHTjS2zt1qfr1NYHuzeLXfQM9H24wFSUt1Mp8"
	OP_ID_SWAP_BASE_IN          = 9
	OP_ID_SWAP_BASE_OUT         = 11
)

func SwapBaseInInstruction(param raydium.RaydiumSwapBaseInInstructionParam) types.Instruction {

	data := make([]byte, 17)
	data[0] = OP_ID_SWAP_BASE_IN
	binary.LittleEndian.PutUint64(data[1:9], param.AmountIn.BigInt().Uint64())
	binary.LittleEndian.PutUint64(data[9:17], param.MinAmountOut.BigInt().Uint64())

	return types.Instruction{common.PublicKeyFromString(RAYDIUM_AMM_PROGRAM_ADDRESS), param.Accounts, data}

}

func SwapBaseOutInstruction(param raydium.RaydiumSwapBaseOutInstructionParam) types.Instruction {

	data := make([]byte, 17)
	data[0] = OP_ID_SWAP_BASE_OUT
	binary.LittleEndian.PutUint64(data[1:9], param.MaxAmountIn.BigInt().Uint64())
	binary.LittleEndian.PutUint64(data[9:17], param.AmountOut.BigInt().Uint64())

	return types.Instruction{common.PublicKeyFromString(RAYDIUM_AMM_PROGRAM_ADDRESS), param.Accounts, data}

}
