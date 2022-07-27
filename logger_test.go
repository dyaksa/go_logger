package go_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := logrus.New()

	logger.Println("hello")
}
