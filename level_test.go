package go_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()

	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("trace")
	logger.Debug("debug")
	logger.Infoln("info")
	logger.Warn("warn")
	logger.Error("error")
}
