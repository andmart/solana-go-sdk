package clmm

import (
	"encoding/binary"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/raydium"
	"github.com/blocto/solana-go-sdk/types"
)

const (
	OP_ID_V1, OP_ID_V2           uint64 = 0xde331ec4da5abe8f, 0x621ec91a0bed042b
	RAYDIUM_CLMM_PROGRAM_ADDRESS        = "CAMMCzo5YL8w4VFF8KVHrK22GGUsp5VTaW7grrKgrWqK"
)

func bool2byte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func SwapBaseVersionedInInstruction(param raydium.RaydiumCLLMSwapInstructionParam) types.Instruction {

	OP_ID := OP_ID_V1
	if param.Version == 2 {
		OP_ID = OP_ID_V2
	}

	data := make([]byte, 33)
	binary.LittleEndian.PutUint64(data[0:8], OP_ID)
	binary.LittleEndian.PutUint64(data[8:16], param.Amount.BigInt().Uint64())
	binary.LittleEndian.PutUint64(data[16:24], param.OtherAmountThreshod.BigInt().Uint64())
	binary.LittleEndian.PutUint64(data[24:32], param.SqrtPriceLimitX64.BigInt().Uint64())
	data[32] = bool2byte(param.IsBaseInput)

	return types.Instruction{common.PublicKeyFromString(RAYDIUM_CLMM_PROGRAM_ADDRESS), param.Accounts, data}
}
