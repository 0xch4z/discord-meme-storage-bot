package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

// var handleGetMemeMessage = func(sess *discordgo.Session, evt *discordgo.MessageCreate) {
// 	msg := evt.Message
// 	switch strings.ToLower(strings.trimSpace(msg.Content)) {
// 	case "!meme":

// 	}
// }

var handlePutMemeMessage = func(sess *discordgo.Session, evt *discordgo.MessageCreate) {
	msg := evt.Message
	switch strings.ToLower(strings.TrimSpace(msg.Content)) {
	case "!meme":
		chanID := msg.ChannelID
		contentParts := strings.Fields(msg.Content)

		if !(len(contentParts) <= 2) {
			log.WithFields(logrus.Fields{
				"messageId": msg.ID,
				"content":   msg.Content,
			}).Error("Could not find meme name to save")

			res := "God damnit! I couldn't figure out what name you wanted me to " +
				"save this meme under! When sending an image in, format the message " +
				"like so `!meme <meme_name>`. (:"
			memeBot.sendTextMessage(chanID, res)
			return
		}

		name := contentParts[1]

		if msg.Embeds != nil && len(msg.Embeds) != 0 && msg.Embeds[0].Image != nil {
			uri := msg.Embeds[0].Image.URL
			err := storage.Put(name, uri)
			if err != nil {
				log.WithFields(logrus.Fields{
					"messageID": msg.ID,
					"uri":       uri,
					"name":      name,
				}).Error("Could not save image")
			}

			memeBot.sendTextMessage(chanID, "God damnit! I couldn't save that for some reason. ):")
			return
		} else {
			log.WithFields(logrus.Fields{
				"messageId": msg.ID,
			}).Error("No imaged attached")

			memeBot.sendTextMessage(chanID, "God damnit! I couldn't find a meme to save on that message. ):")
			return
		}
	}
}
