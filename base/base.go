package base

import (
	"log/slog"
	"os"
)

type Level = slog.Level
type Logger = slog.Logger

const (
	LevelDebug Level = -4
	LevelInfo  Level = 0
	LevelWarn  Level = 4
	LevelError Level = 8
)

type qLog struct {
	timeFormat string
	logfile    string
	level      Level
}

func (q *qLog) SetTimeFormat(layout string) {
	q.timeFormat = layout
}

func (q *qLog) SetOutput(logfile string) {
	q.logfile = logfile
}

func (q *qLog) SetLevel(level Level) {
	q.level = level
}

func (q *qLog) options() *slog.HandlerOptions {
	return &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == "time" {
				if q.timeFormat != "" {
					t := a.Value.Time().Format(q.timeFormat)
					a.Value = slog.StringValue(t)
				}
			}
			return a
		},
	}
}

func (q *qLog) New() (*Logger, error) {
	opts := q.options()

	if q.logfile == "" {
		opts.Level = slog.LevelDebug
		handler := slog.NewJSONHandler(os.Stdout, opts)
		return slog.New(handler), nil
	}

	accessFile, err := os.OpenFile(q.logfile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	opts.Level = q.level
	handler := slog.NewJSONHandler(accessFile, opts)
	return slog.New(handler), nil
}

func NewLog() *qLog {
	return &qLog{}
}
