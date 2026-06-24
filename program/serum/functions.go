package serum

import (
	"context"
	"fmt"

	c "github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/token"
	bin "github.com/gagliardetto/binary"
)

type AccountFlag uint64

const (
	AccountFlagInitialized = AccountFlag(1 << iota)
	AccountFlagMarket
	AccountFlagOpenOrders
	AccountFlagRequestQueue
	AccountFlagEventQueue
	AccountFlagBids
	AccountFlagAsks
	AccountFlagDisabled
)

type MarketV2 struct {
	SerumPadding           [5]byte `json:"-"`
	AccountFlags           AccountFlag
	OwnAddress             common.PublicKey
	VaultSignerNonce       uint64
	BaseMint               common.PublicKey
	QuoteMint              common.PublicKey
	BaseVault              common.PublicKey
	BaseDepositsTotal      uint64
	BaseFeesAccrued        uint64
	QuoteVault             common.PublicKey
	QuoteDepositsTotal     uint64
	QuoteFeesAccrued       uint64
	QuoteDustThreshold     uint64
	RequestQueue           common.PublicKey
	EventQueue             common.PublicKey
	Bids                   common.PublicKey
	Asks                   common.PublicKey
	BaseLotSize            uint64
	QuoteLotSize           uint64
	FeeRateBPS             uint64
	ReferrerRebatesAccrued uint64
	EndPadding             [7]byte `json:"-"`
}

type MarketMeta struct {
	Address         common.PublicKey `json:"address"`
	Name            string           `json:"name"`
	Deprecated      bool             `json:"deprecated"`
	QuoteMint       token.MintAccount
	BaseMint        token.MintAccount
	MarketProgramId common.PublicKey

	MarketV2 MarketV2
}

func FetchMarket(ctx context.Context, rpcCli *c.Client, marketAddr common.PublicKey) (*MarketMeta, error) {
	acctInfo, err := rpcCli.GetAccountInfo(ctx, marketAddr.String())
	if err != nil {
		return nil, fmt.Errorf("unable to get market account:%w", err)
	}

	meta := &MarketMeta{
		Address: marketAddr,
		MarketProgramId: acctInfo.Owner,
	}

	dataLen := len(acctInfo.Data)
	switch dataLen {
	// case 380:
	// 	// if err := meta.MarketV1.Decode(acctInfo.Value.Data); err != nil {
	// 	// 	return nil, fmt.Errorf("decoding market v1: %w", err)
	// 	// }
	// 	return nil, fmt.Errorf("Unsupported market version, w/ data length of 380")

	case 388:
		if err := meta.MarketV2.Decode(acctInfo.Data); err != nil {
			return nil, fmt.Errorf("decoding market v2: %w", err)
		}

	default:
		return nil, fmt.Errorf("unsupported market data length: %d", dataLen)
	}

	//Pumpfun exchange base and quote mints
	if meta.MarketV2.BaseMint.String() == common.Solana.String(){ //Trocado
		mint := meta.MarketV2.BaseMint
		meta.MarketV2.BaseMint = meta.MarketV2.QuoteMint
		meta.MarketV2.QuoteMint = mint
	}

	resp, err := rpcCli.GetAccountInfo(ctx, meta.MarketV2.QuoteMint.String())
	if err != nil {
		return nil, fmt.Errorf("getting quote mint: %w", err)
	}
	bin.NewBinDecoder(resp.Data).Decode(&meta.QuoteMint)

	resp, err = rpcCli.GetAccountInfo(ctx, meta.MarketV2.BaseMint.String())
	if err != nil {
		return nil, fmt.Errorf("getting quote mint: %w", err)
	}
	bin.NewBinDecoder(resp.Data).Decode(&meta.BaseMint)

	return meta, nil
}

func (m *MarketV2) Decode(in []byte) error {
	decoder := bin.NewBinDecoder(in)
	err := decoder.Decode(&m)
	if err != nil {
		return fmt.Errorf("unpack: %w", err)
	}
	return nil
}
