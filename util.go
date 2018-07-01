package main

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func retryOnBadGateway(fn func() error) {
	var err error
	for i := 0; i < 3; i++ {
		err = fn()
		if err != nil {
			if strings.HasPrefix(err.Error(), "HTTP 502") {
				log.Errorf("Bad Gateway; retrying... (attempt %d/3)", i+1)
				time.Sleep(1 * time.Second)
				continue
			} else {
				if err != nil {
					panic(err)
				}
			}
		} else {
			return
		}
	}
}

func getContentType(file *os.File) (string, error) {
	buf := make([]byte, 512)
	_, err := file.Read(buf)
	defer file.Close()
	if err != nil {
		return "", err
	}

	file.Seek(0, 0)

	return http.DetectContentType(buf), nil
}

func wrapFile(file *os.File) *discordgo.File {
	contype, err := getContentType(file)
	if err != nil {
		log.Errorf("failed to detect content type for file `%s`; letting http infer", file.Name())
	}

	return &discordgo.File{
		Name:        file.Name(),
		Reader:      file,
		ContentType: contype,
	}
}
