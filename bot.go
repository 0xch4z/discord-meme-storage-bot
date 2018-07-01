package main

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

var memeBot *bot

type bot struct {
	Session *discordgo.Session
}

func newBot(botTok string) *bot {
	sess, err := discordgo.New("Bot " + botTok)
	if err != nil {
		log.Fatalf("Session could not be created: %s", err.Error())
	}

	return &bot{
		Session: sess,
	}
}

func (b *bot) sendTextMessage(chanID string, msg string) {
	log.WithFields(logrus.Fields{
		"channelID": chanID,
		"message":   msg,
	}).Info("Sending text message")
	retryOnBadGateway(func() error {
		_, err := b.Session.ChannelMessageSend(chanID, msg)
		return err
	})
}

func (b *bot) sendFileMessage(chanID string, file *os.File) {
	log.WithFields(logrus.Fields{
		"channelID": chanID,
		"fileName":  file.Name,
	}).Info("Sending file message")
	wrappedFile := wrapDiscordFile(file)
	retryOnBadGateway(func() error {
		_, err := b.Session.ChannelMessageSendComplex(chanID, &discordgo.MessageSend{
			File: wrappedFile,
		})
		return err
	})
}

func (b *bot) AddHandler(handleFunc interface{}) {
	b.Session.AddHandler(handleFunc)
}

func (b *bot) Listen() {
	b.Session.Open()
}
