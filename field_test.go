package go_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestFields(t *testing.T) {
	logger := logrus.New()

	logger.WithFields(logrus.Fields{
		"name": "dyaksa",
		"age":  "nour",
	}).Info("Hello")

	logger.Info("Hello")
}
