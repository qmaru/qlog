package qlog

import (
	"testing"

	"github.com/qmaru/qlog/base"
	"github.com/qmaru/qlog/console"
)

func TestBase(t *testing.T) {
	log := base.NewLog()
	logger, err := log.New()
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("Hello base")
}

func TestConsole(t *testing.T) {
	logger, err := console.New("test", "")
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("testing", "Hello console")
}
