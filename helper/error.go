package helper

import "github.com/sirupsen/logrus"

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func LogInfo(key string, args ...interface{}) *logrus.Entry {
	logging := logrus.New()
	logging.SetFormatter(&logrus.JSONFormatter{})
	return logging.WithField(key, args)
}
