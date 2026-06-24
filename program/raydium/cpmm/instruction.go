package cpmm

import (
	"encoding/binary"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/raydium"
	"github.com/blocto/solana-go-sdk/types"
)

const (
	OP_ID                        = 0xde331ec4da5abe8f
	RAYDIUM_CPMM_PROGRAM_ADDRESS = "CPMMoo8L3F4NbTegBCKVNunggL7H1ZpdTHKxQB5qKP1C"
)

func SwapBaseInInstruction(param raydium.RaydiumSwapBaseInInstructionParam) types.Instruction {
	data := make([]byte, 25)
	binary.LittleEndian.PutUint64(data[0:8], OP_ID)
	binary.LittleEndian.PutUint64(data[8:16], param.AmountIn.BigInt().Uint64())
	binary.LittleEndian.PutUint64(data[16:24], param.MinAmountOut.BigInt().Uint64())

	return types.Instruction{common.PublicKeyFromString(RAYDIUM_CPMM_PROGRAM_ADDRESS), param.Accounts, data}
}
