package vo

import (
	"fmt"
	"time"
)

type AssetsRespond struct {
	Total     int        `json:"total"`
	Txs       []AssetsVo `json:"txs"`
	AssetType string     `json:"asset_type"`
}

type AssetsVo struct {
	TokenId         string    `json:"token_id"`
	Type            string    `json:"type"`
	Owner           string    `json:"owner"`
	Gateway         string    `json:"gateway"`
	Symbol          string    `json:"symbol"`
	InitialSupply   int64     `json:"initial_supply,string"`
	MaxSupply       int64     `json:"max_supply,string"`
	Mintable        bool      `json:"mintable,string"`
	Decimal         int32     `json:"decimal,string"`
	CanonicalSymbol string    `json:"canonical_symbol"`
	SymbolMin       string    `json:"symbol_min"`
	Name            string    `json:"name"`
	MintTo          string    `json:"mint_to"`
	Amount          int64     `json:"amount"`
	SrcOwner        string    `json:"src_owner"`
	DstOwner        string    `json:"dst_owner"`
	Height          int64     `json:"height"`
	TxHash          string    `json:"tx_hash"`
	TxFee           Fee       `json:"tx_fee"`
	TxStatus        string    `json:"tx_status"`
	Timestamp       time.Time `json:"timestamp"`
}

type Coins []Coin

type Fee struct {
	Amount Coins `json:"amount"`
	Gas    int64 `json:"gas"`
}

type AssetTokens struct {
	Symbol  string `json:"symbol"`
	Decimal int    `json:"decimal"`
}

func (b AssetTokens) String() string {

	return fmt.Sprintf(`
		Symbol          :%v
		Decimal         :%v
		`, b.Symbol, b.Decimal)
}