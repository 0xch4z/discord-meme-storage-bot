package main

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.New()

	file := conf.Logger.File
	level := conf.Logger.Level

	if len(level) != 0 {
		log.Level = map[string]logrus.Level{
			"debug": logrus.DebugLevel,
			"info":  logrus.InfoLevel,
			"warn":  logrus.WarnLevel,
			"error": logrus.ErrorLevel,
			"fatal": logrus.FatalLevel,
			"panic": logrus.PanicLevel,
		}[level]
	}

	if len(file) != 0 {
		logFile, err := os.Open(file)
		if err != nil {
			log.Errorf("Error: cannot open file `%s` for logging; defaulting to stdOut", file)
		} else {
			mw := io.MultiWriter(os.Stdout, logFile)
			logrus.SetOutput(mw)
			log.Error("test")
		}
	}
}
