package main

import "github.com/jinzhu/gorm"

type meme struct {
	gorm.Model
	Guild    *guild
	Filename string `gorm:"unique_index"`
	Name     string `gorm:"index:idx_guild_meme"`
}
