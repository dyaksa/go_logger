package go_logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

type SampleHook struct {
}

func (hook *SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (hook *SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("Hello", entry.Level, entry.Message)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()

	logger.AddHook(&SampleHook{})

	logger.Warn("Hello")
}
