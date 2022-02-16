package main

import "github.com/sirupsen/logrus"

type store struct {
	log *logrus.Entry
}

func initStore() *store {
	return &store{
		log: logrus.NewEntry(logrus.StandardLogger()),
	}
}
