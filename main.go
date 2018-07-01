package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func main() {
	memeBot = newBot(conf.Discord.BotToken)
	memeBot.AddHandler(handlePutMemeMessage)
	memeBot.AddHandler(handleGetMemeMessage)
	memeBot.Listen()

	log.WithFields(logrus.Fields{"token": conf.Discord.BotToken}).Info("Bot has booted")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	memeBot.Session.Close()
}
