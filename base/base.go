package base

import (
	"log/slog"
	"os"

	h "github.com/qmaru/qlog/base/handler"
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
	json       bool
	text       bool
	simple     bool
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

func (q *qLog) SetTextHandler() {
	q.text = true
	q.json = false
	q.simple = false
}

func (q *qLog) SetJSONHandler() {
	q.text = false
	q.json = true
	q.simple = false
}

func (q *qLog) SetSimpleHandler() {
	q.text = false
	q.json = false
	q.simple = true
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

	var handler slog.Handler
	var output *os.File

	if q.logfile == "" {
		opts.Level = slog.LevelDebug
		output = os.Stdout
	} else {
		opts.Level = q.level
		accessFile, err := os.OpenFile(q.logfile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			return nil, err
		}
		output = accessFile
	}

	if q.simple {
		handler = h.NewSimpleHandler(opts, q.timeFormat)
	} else if q.json {
		handler = slog.NewJSONHandler(output, opts)
	} else if q.text {
		handler = slog.NewTextHandler(output, opts)
	} else {
		handler = slog.NewJSONHandler(output, opts)
	}

	return slog.New(handler), nil
}

func NewLog() *qLog {
	return &qLog{}
}
