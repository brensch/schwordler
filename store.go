package main

import "github.com/sirupsen/logrus"

type store struct {
	log *logrus.Entry
}

func initStore() *store {
	format := new(logrus.TextFormatter)
	format.FullTimestamp = true
	logger := logrus.StandardLogger()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(format)
	return &store{
		log: logrus.NewEntry(logger),
	}
}
