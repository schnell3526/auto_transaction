package symbol

type SymbolType string

const (
	BTC     = "BTC" // ビットコイン
	ETH     = "ETC" // イーサリアム
	BCH     = "BCH" //ビットコインキャッシュ
	LTC     = "LTC" // ライトコイン
	XRP     = "XRP" // リップル
	XEM     = "XEM" // ネム
	XLM     = "XLM" // ステラルーメン
	BTC_JPY = "BTC_JPY"
	ETH_JPY = "ETH_JPY"
	BCH_JPY = "BCH_JPY"
	LTC_JPY = "LTC_JPY"
	XRP_JPY = "XRP_JPY"
)

// 銘柄名
type Symbol struct {
	Type    SymbolType
	Literal string
}

func NewSymbol(symbol SymbolType) Symbol {
	return Symbol{Type: symbol, Literal: string(symbol)}
}
