package log

import (
	"github.com/sirupsen/logrus"
)

type loggerType struct{}

var Logger *loggerType

func (l *loggerType) Log(args ...interface{}) {
	logrus.Error(args...)
}

type Fields map[string]interface{}

func Info(msg string, fields ...Fields) {
	logrus.WithFields(toFields(fields)).Info(msg)
}

func Error(err error, fields ...Fields) {
	logrus.WithFields(toFields(fields)).
		Error(err)
}

func toFields(args []Fields) logrus.Fields {
	fields := logrus.Fields{}

	for _, arg := range args {
		for key, val := range arg {
			fields[key] = val
		}
	}

	return fields
}
