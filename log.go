package log

import (
	"github.com/morganxf/log/pkg/hooks/file"
	"github.com/morganxf/log/pkg/util"
	"github.com/sirupsen/logrus"
	"os"
)

var logger *logrus.Logger

func init() {}


func InitLogger(level logrus.Level) {
	logger = util.NewLogger(os.Stderr, level)
}

func InitFileHook(logDir string) {
	hook := file.NewHook(logDir)
	logger.AddHook(hook)
}
