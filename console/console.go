package console

import (
	"log/slog"

	"github.com/qmaru/qlog/base"
)

type ConsoleLogger struct {
	user   string
	logger *slog.Logger
}

func New(user, output string) (*ConsoleLogger, error) {
	log := base.NewLog()
	log.SetTimeFormat("2006-01-02 15:04:05")
	if output != "" {
		log.SetOutput(output)
	}

	logger, err := log.New()
	if err != nil {
		return nil, err
	}
	return &ConsoleLogger{
		user:   user,
		logger: logger,
	}, nil
}

func (c *ConsoleLogger) Info(taskName, message string) {
	c.logger.Info(c.user, "user", taskName, "content", message)
}

func (c *ConsoleLogger) Warn(taskName, message string) {
	c.logger.Warn(c.user, "user", taskName, "content", message)
}

func (c *ConsoleLogger) Error(taskName, message string) {
	c.logger.Error(c.user, "user", taskName, "content", message)
}
