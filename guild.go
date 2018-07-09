package main

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type guild struct {
	gorm.Model
	ID    string `gorm:"primary_key index:idx_guild_meme"`
	Memes *[]meme
}

func findOrCreateGuild(id string) (*guild, error) {
	var gld *guild
	if err := db.FirstOrCreate(&gld, &guild{ID: id}).Error; err != nil {
		log.WithFields(logrus.Fields{
			"errorMessage": err.Error(),
			"guildID":      id,
		}).Error("could not find or create guild")
		return nil, err
	}
	return gld, nil
}

func (g *guild) findOrCreateMeme(m *meme) (*meme, error) {
	m.Guild = g
	if err := db.FirstOrCreate(m).Error; err != nil {
		log.WithFields(logrus.Fields{
			"errorMessage": err.Error(),
			"filename":     m.Filename,
		}).Error("could not create meme for guild")
		return nil, err
	}
	return m, nil
}
