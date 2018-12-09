package file

import (
	"fmt"
	"github.com/morganxf/log/pkg/util"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

var LevelFile map[logrus.Level]string = map[logrus.Level]string{
	logrus.PanicLevel: "panic.log",
	logrus.FatalLevel: "fatal.log",
	logrus.ErrorLevel: "error.log",
	logrus.WarnLevel: "warn.log",
	logrus.InfoLevel: "info.log",
	logrus.DebugLevel: "debug.log",
}

type Hook struct {
	LoggerMap map[logrus.Level]*logrus.Logger
	LogDir string
}

func NewHook(logDir string) *Hook {
	h := &Hook{LogDir: logDir}
	h.InitLoggerMap()
	return h
}

func (hook *Hook) InitLoggerMap() {
	loggerMap := make(map[logrus.Level]*logrus.Logger)
	for level, file := range LevelFile {
		//f, err := openFile(path.Join(hook.LogDir, file))
		//if err != nil {
		//	continue
		//}
		//logger := util.NewLogger(f, level)
		lumberjackLogger := &lumberjack.Logger{
			Filename: path.Join(hook.LogDir, file),
			MaxSize: 2,	// 100M
			MaxBackups: 3,
			MaxAge:     7,
		}
		logger := util.NewLogger(lumberjackLogger, level)
		loggerMap[level] = logger
	}
	hook.LoggerMap = loggerMap
}

func openFile(filename string) (*os.File, error) {
	f, err := os.OpenFile(filename, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0644)
	return f, err
}

func (hook *Hook) Close() {
	for _, logger := range hook.LoggerMap {
		if logger == nil {
			continue
		}
		logger.Out.(*os.File).Close()
	}
}

func (hook *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *Hook) Fire(entry *logrus.Entry) error {
	msg := entry.Message
	switch entry.Level {
	case logrus.PanicLevel:
		hook.Panic(msg)
	case logrus.FatalLevel:
		hook.Fatal(msg)
	case logrus.ErrorLevel:
		hook.Error(msg)
	case logrus.WarnLevel:
		hook.Warn(msg)
	case logrus.InfoLevel:
		hook.Info(msg)
	case logrus.DebugLevel:
		hook.Debug(msg)
	default:
		fmt.Fprintf(os.Stderr, "Unrecognized the log level: %v", entry.Level)
	}
	return nil
}

func (hook *Hook) Panic(msg string) {
	hook.LoggerMap[logrus.PanicLevel].Panic(msg)
	hook.LoggerMap[logrus.FatalLevel].Panic(msg)
	hook.LoggerMap[logrus.ErrorLevel].Panic(msg)
	hook.LoggerMap[logrus.WarnLevel].Panic(msg)
	hook.LoggerMap[logrus.InfoLevel].Panic(msg)
	hook.LoggerMap[logrus.DebugLevel].Panic(msg)
}

func (hook *Hook) Fatal(msg string) {
	hook.LoggerMap[logrus.FatalLevel].Fatal(msg)
	hook.LoggerMap[logrus.ErrorLevel].Fatal(msg)
	hook.LoggerMap[logrus.WarnLevel].Fatal(msg)
	hook.LoggerMap[logrus.InfoLevel].Fatal(msg)
	hook.LoggerMap[logrus.DebugLevel].Fatal(msg)
}

func (hook *Hook) Error(msg string) {
	hook.LoggerMap[logrus.ErrorLevel].Error(msg)
	hook.LoggerMap[logrus.WarnLevel].Error(msg)
	hook.LoggerMap[logrus.InfoLevel].Error(msg)
	hook.LoggerMap[logrus.DebugLevel].Error(msg)
}

func (hook *Hook) Warn(msg string) {
	hook.LoggerMap[logrus.WarnLevel].Warn(msg)
	hook.LoggerMap[logrus.InfoLevel].Warn(msg)
	hook.LoggerMap[logrus.DebugLevel].Warn(msg)
}

func (hook *Hook) Info(msg string) {
	hook.LoggerMap[logrus.InfoLevel].Info(msg)
	hook.LoggerMap[logrus.DebugLevel].Info(msg)
}

func (hook *Hook) Debug(msg string) {
	hook.LoggerMap[logrus.DebugLevel].Debug(msg)
}