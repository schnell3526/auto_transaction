package gmo

import (
	"encoding/json"
	"time"
)

type JsonData interface {
	Dump() string
}

// 取引所ステータス
type ServiceStatus struct {
	Status int `json:"status"`
	Data   struct {
		Status string `json:"status"`
	} `json:"data"`
	Responsetime time.Time `json:"responsetime"`
}

func (ss *ServiceStatus) Dump() string {
	i, _ := json.MarshalIndent(ss, "", " ")
	return string(i)
}

// 最新レート
type LatestRate struct {
	Status int `json:"status"`
	Data   []struct {
		Ask       string    `json:"ask"`
		Bid       string    `json:"bid"`
		High      string    `json:"high"`
		Last      string    `json:"last"`
		Low       string    `json:"low"`
		Symbol    string    `json:"symbol"`
		Timestamp time.Time `json:"timestamp"`
		Volume    string    `json:"volume"`
	} `json:"data"`
	Responsetime time.Time `json:"responsetime"`
}

func (lr *LatestRate) Dump() string {
	i, _ := json.MarshalIndent(lr, "", " ")
	return string(i)
}

// 板情報
type OrderBook struct {
	Status int `json:"status"`
	Data   struct {
		Asks []struct { //売り注文の情報
			Price string `json:"price"`
			Size  string `json:"size"`
		} `json:"asks"`
		Bids []struct { // 買い注文の情報
			Price string `json:"price"`
			Size  string `json:"size"`
		} `json:"bids"`
		Symbol string `json:"symbol"`
	} `json:"data"`
	Responsetime time.Time `json:"responsetime"`
}

func (ob *OrderBook) Dump() string {
	i, _ := json.MarshalIndent(ob, "", " ")
	return string(i)
}

// 取引履歴
type TradeHistory struct {
	Status int `json:"status"`
	Data   struct {
		Pagination struct {
			CurrentPage int `json:"currentPage"`
			Count       int `json:"count"`
		} `json:"pagination"`
		List []struct {
			Price     string    `json:"price"`
			Side      string    `json:"side"`
			Size      string    `json:"size"`
			Timestamp time.Time `json:"timestamp"`
		} `json:"list"`
	} `json:"data"`
	Responsetime time.Time `json:"responsetime"`
}

func (th *TradeHistory) Dump() string {
	i, _ := json.MarshalIndent(th, "", " ")
	return string(i)
}

// KLine情報の取得
type KLineData struct {
	Status int `json:"status"`
	Data   []struct {
		OpenTime string `json:"openTime"`
		Open     string `json:"open"`
		High     string `json:"high"`
		Low      string `json:"low"`
		Close    string `json:"close"`
		Volume   string `json:"volume"`
	} `json:"data"`
	Responsetime time.Time `json:"responsetime"`
}

func (kld *KLineData) Dump() string {
	i, _ := json.MarshalIndent(kld, "", " ")
	return string(i)
}
