package file

import (
	"github.com/morganxf/log/pkg/util"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

var testLogDir = path.Join(os.Getenv("GOPATH"), "/src/github.com/morganxf/log/pkg/hooks/file/testData")

func TestNewHook(t *testing.T) {
	h := NewHook(testLogDir)
	defer h.Close()
	assert.NotNil(t, h)
	assert.Equal(t, 6, len(h.LoggerMap))
	assert.NotEqual(t, "", h.LogDir)
}

func TestHook_InitLoggerMap(t *testing.T) {
	h := &Hook{LogDir: testLogDir}
	h.InitLoggerMap()
	defer h.Close()
	assert.Equal(t, 6, len(h.LoggerMap))
}

func TestHook_Fire(t *testing.T) {
	logger := util.NewLogger(os.Stderr, logrus.InfoLevel)
	hook := NewHook(testLogDir)
	defer hook.Close()
	logger.AddHook(hook)
	//logger.Warn("fire test")
}