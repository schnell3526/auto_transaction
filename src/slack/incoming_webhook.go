package slack

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

type requestBody struct {
	Text      string `json:"text"`
	UserName  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
}

type SlackClient struct {
	endpoint string
}

func NewSlackClient(endpoint string) *SlackClient {
	return &SlackClient{endpoint: endpoint}
}

// テキストデータをslackに通知
func (sc *SlackClient) SendMessage(text string) error {
	// リクエストbodyを作成
	requestBody := &requestBody{
		Text:      text,
		UserName:  "仮想通貨価格通知bot",
		IconEmoji: ":ghost:",
	}

	// 構造体をjson形式のバイト列に変換
	raw_json, err := json.Marshal(requestBody)
	if err != nil {
		return errors.Wrap(err, "at SendMessage() : ")
	}

	// リクエストを作成
	req, err := http.NewRequest("POST", sc.endpoint, bytes.NewBuffer(raw_json))
	if err != nil {
		return errors.Wrap(err, "at SendMessage() : ")
	}
	req.Header.Set("Content-Type", "application/json")

	// sends an HTTP request and returns an HTTP response
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "at SendMessage() : ")
	}
	defer resp.Body.Close()

	// ボディをエンコード
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "at SendMessage() : ")
	}

	// 結果を出力
	log.Print(string(byteArray))

	return nil
}
