package main

import (
	"fmt"
	"log"

	"github.com/schnell3526/auto_auto_transaction/src/app"
	"github.com/schnell3526/auto_auto_transaction/src/config"
)

type RequestBody struct {
	Text      string `json:"text"`
	UserName  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
}

func main() {

	// 設定ファイルの読み込み
	log.Println("***start***")
	log.Println("***read configuration file***")
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Configuration failed: %v", err)
	}
	fmt.Println(config.Dump())

	log.Println("***start aplication***")
	if err := app.StartApp(*config); err != nil {
		log.Fatalf("Apprication Error: %v", err)
	}
	log.Println("***stop aplication***")

}
