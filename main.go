package main

func main() {
	memeBot = newBot(conf.Discord.BotToken)
	memeBot.AddHandler(handlePutMemeMessage)
	memeBot.AddHandler(handleGetMemeMessage)
	memeBot.Listen()
	<-make(chan struct{})
}
