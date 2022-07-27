package go_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestFormatter(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello")
}
