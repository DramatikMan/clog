package clog

import (
	"io"
	"log/slog"
)

type group struct {
	name  string
	attrs []slog.Attr
}

type consoleHandler struct {
	out     io.Writer
	options *slog.HandlerOptions
	groups  []group
}
