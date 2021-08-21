package side

type SideType string

const (
	BUY  = "BUY"
	SELL = "SELL"
)

// 売買区分
type Side struct {
	Type    SideType
	Literal string
}

func NewSide(side SideType) Side {
	return Side{Type: side, Literal: string(side)}
}
