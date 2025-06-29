package clmm

import (
	"encoding/binary"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/raydium"
	"github.com/blocto/solana-go-sdk/types"
)

const (
	OP_ID                        = 0xde331ec4da5abe8f
	RAYDIUM_CLMM_PROGRAM_ADDRESS = "CPMMQoayoCZGUq4wQRxrPBNrrExU6PLg1eEAXC83KDzv"
)

func SwapBaseInInstruction(param raydium.RaydiumSwapBaseInInstructionParam) types.Instruction {
	data := make([]byte, 24)
	binary.LittleEndian.PutUint32(data[0:8], OP_ID)
	binary.LittleEndian.PutUint64(data[8:16], param.AmountIn.BigInt().Uint64())
	binary.LittleEndian.PutUint64(data[16:24], param.MinAmountOut.BigInt().Uint64())

	return types.Instruction{common.PublicKeyFromString(RAYDIUM_CLMM_PROGRAM_ADDRESS), param.Accounts, data}
}
