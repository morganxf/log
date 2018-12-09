package util

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)



func NewLogger(writer io.Writer, level logrus.Level) *logrus.Logger {
	logger := &logrus.Logger{
		Out:          writer,
		Formatter:    &logrus.TextFormatter{
			FullTimestamp: true,
		},
		Hooks:        make(logrus.LevelHooks),
		Level:        level,
		ExitFunc:     os.Exit,
		ReportCaller: true,
	}
	return logger
}
