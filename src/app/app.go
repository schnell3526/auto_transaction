package app

import (
	"github.com/schnell3526/auto_auto_transaction/src/config"
	"github.com/schnell3526/auto_auto_transaction/src/gmo"
	"github.com/schnell3526/auto_auto_transaction/src/gmo/symbol"
	"github.com/schnell3526/auto_auto_transaction/src/slack"
)

func StartApp(config config.Config) error {

	// slackクライアントを作成
	slackClient := slack.NewSlackClient(config.SlackApiKey)

	// ビットコインの最新レートを取得
	res := gmo.GetLatestRate(symbol.BTC, true)

	// slackにメッセージを送信
	if err := slackClient.SendMessage(res.Dump()); err != nil {
		return err
	}
	return nil
}
