package schwordler

import "github.com/sirupsen/logrus"

type Store struct {
	Log *logrus.Entry
}

func InitStore() *Store {
	format := new(logrus.TextFormatter)
	format.FullTimestamp = true
	logger := logrus.StandardLogger()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(format)
	return &Store{
		Log: logrus.NewEntry(logger),
	}
}
