package main

func main() {
	memeBot = newBot(conf.Discord.BotToken)
	memeBot.AddHandler(handlePutMemeMessage)
}
