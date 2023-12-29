package helper

import "github.com/sirupsen/logrus"

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func NewLogging() *logrus.Logger {
	logging := logrus.New()
	return logging
}

func LogInfo(key string, args ...interface{}) *logrus.Logger {
	logging := NewLogging()
	logging.SetFormatter(&logrus.JSONFormatter{})
	logging.WithField(key, args).Info(key)
	return logging
}

func LogFatal(key string, args ...interface{}) *logrus.Logger {
	logging := NewLogging()
	logging.SetFormatter(&logrus.JSONFormatter{})
	logging.Fatalf("%s : %v", key, args)
	return logging
}
