package main

import (
	"github.com/brensch/schwordler"
	"github.com/sirupsen/logrus"
)

func main() {
	s := schwordler.InitStore(logrus.New())

	_ = s

}
