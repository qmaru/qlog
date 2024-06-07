package handler

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
)

type SimpleHandler struct {
	timeformat string
	opts       slog.HandlerOptions
}

func (h *SimpleHandler) Enabled(_ context.Context, l slog.Level) bool {
	minLevel := slog.LevelInfo
	if h.opts.Level != nil {
		minLevel = h.opts.Level.Level()
	}
	return l >= minLevel
}

func (h *SimpleHandler) Handle(ctx context.Context, r slog.Record) error {
	t := r.Time.Format(h.timeformat)
	l := r.Level.String()
	m := r.Message

	o := make(map[string]any, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		o[a.Key] = a.Value.Any()
		return true
	})

	q := make([]string, 0)
	for k, v := range o {
		q = append(q, fmt.Sprintf("%s=%v", k, v))
	}

	fmt.Printf("%s - %s - %s %s\n", t, l, m, strings.Join(q, ","))
	return nil
}

func (h *SimpleHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *SimpleHandler) WithGroup(name string) slog.Handler {
	return h
}

func NewSimpleHandler(opts *slog.HandlerOptions, timeformat string) *SimpleHandler {
	return &SimpleHandler{
		timeformat: timeformat,
	}
}
