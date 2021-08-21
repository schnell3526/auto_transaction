package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
)

// 設定用構造体
type Config struct {
	ApiKey    string `json:"apiKey"`
	SecretKey string `json:"secretKey"`

	SlackApiKey string `json:"slack-api-key"`
}

// 綺麗に出力
func (c *Config) Dump() string {
	str, _ := json.MarshalIndent(c, "", " ")
	return string(str)
}

func NewConfig() (*Config, error) {
	var config Config

	// ファイルを全て読み込み
	raw, err := ioutil.ReadFile("src/config/config.json")
	if err != nil {
		return &config, errors.Wrap(err, "at config.NewConfig()")
	}

	// jsonデータを構造体に変換
	if err := json.Unmarshal(raw, &config); err != nil {
		return &config, errors.Wrap(err, "at config.NewConfig()")
	}

	return &config, err
}
