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

	o := make([]string, r.NumAttrs()-1)
	r.Attrs(func(a slog.Attr) bool {
		key := a.Key
		value := a.Value.String()
		o = append(o, key, value)
		return true
	})

	fmt.Printf("%s %s %s %s\n", t, l, m, strings.Join(o, " "))
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
