// colorized logging
package clog

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

const (
	LevelTrace   = slog.Level(-8)
	LevelDebug   = slog.LevelDebug
	LevelInfo    = slog.LevelInfo
	LevelWarning = slog.LevelWarn
	LevelError   = slog.LevelError
	LevelFatal   = slog.Level(12)
)

func NewConsoleHandler(options *slog.HandlerOptions) *consoleHandler {
	handler := &consoleHandler{
		out:     os.Stdout,
		options: options,
		groups:  []group{},
	}

	return handler
}

func (handler *consoleHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= handler.options.Level.Level()
}

func (handler *consoleHandler) Handle(ctx context.Context, record slog.Record) error {
	levelText, levelName := record.Level.String(), record.Level.String()

	switch record.Level {
	case LevelTrace:
		levelName = "TRACE"
		levelText = fmt.Sprintf("\033[30m\033[44m %s \033[0m", levelName)
	case LevelDebug:
		levelText = fmt.Sprintf("\033[30m\033[46m %s \033[0m", levelText)
	case LevelInfo:
		levelText = fmt.Sprintf("\033[30m\033[42m %s \033[0m", levelText)
	case LevelWarning:
		levelName = "WARNING"
		levelText = fmt.Sprintf("\033[30m\033[43m %s \033[0m", levelName)
	case LevelError:
		levelText = fmt.Sprintf("\033[30m\033[41m %s \033[0m", levelText)
	case LevelFatal:
		levelName = "FATAL"
		levelText = fmt.Sprintf("\033[30m\033[101m %s \033[0m", levelName)
	}

	var attrs []slog.Attr
	groups := handler.groups

	if record.NumAttrs() == 0 {
		for len(groups) > 0 && groups[len(groups)-1].name != "" {
			groups = groups[:len(groups)-1]
		}
	} else {
		record.Attrs(func(attr slog.Attr) bool {
			attrs = append(attrs, attr)
			return true
		})
	}

	for i := len(groups) - 1; i >= 0; i-- {
		if groups[i].name != "" {
			newGroup := slog.Attr{
				Key:   groups[i].name,
				Value: slog.GroupValue(attrs...),
			}

			attrs = []slog.Attr{newGroup}
		} else {
			attrs = append(attrs, groups[i].attrs...)
		}
	}

	message := record.Message

	if len(attrs) > 0 {
		message = fmt.Sprintf("%s %s", message, fmt.Sprint(attrs))
	}

	toWrite := fmt.Sprintf(
		"%d/%d/%d %s:%s:%s %s:%s%s\n",
		record.Time.Year(), record.Time.Month(), record.Time.Day(),
		formatTime(record.Time.Hour()),
		formatTime(record.Time.Minute()),
		formatTime(record.Time.Second()),
		levelText, strings.Repeat(" ", 8-len(levelName)), message,
	)

	_, err := handler.out.Write([]byte(toWrite))
	return err
}

func (handler *consoleHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) == 0 {
		return handler
	}

	return handler.withGroupOrAttrs(group{attrs: attrs})
}

func (handler *consoleHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return handler
	}

	return handler.withGroupOrAttrs(group{name: name})
}

func (handler *consoleHandler) withGroupOrAttrs(g group) slog.Handler {
	NewConsoleHandler := *handler
	NewConsoleHandler.groups = make([]group, len(handler.groups)+1)
	copy(NewConsoleHandler.groups, handler.groups)
	NewConsoleHandler.groups[len(NewConsoleHandler.groups)-1] = g

	return &NewConsoleHandler
}
