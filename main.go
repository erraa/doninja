package main

import (
	"os"

	bot "github.com/erraa/doninja/bot"
	"github.com/erraa/doninja/config"
	"github.com/erraa/doninja/utils"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var BotID string

func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	textFormatter := new(prefixed.TextFormatter)
	textFormatter.FullTimestamp = true

	log := utils.LogWithPrefix("main")

	conf := config.ReadConfig()
	log.Info("Prefix", conf.Discord.BotPrefix)

	bot.Start()

	<-make(chan struct{})
	return
}
