package schwordler

import "github.com/sirupsen/logrus"

type Store struct {
	Log logrus.FieldLogger
}

func InitStore(log logrus.FieldLogger) *Store {

	return &Store{
		Log: log,
	}
}
