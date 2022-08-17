package logger

import (
	"context"
	"fmt"
	"io"

	"github.com/fatih/color"
)

var (
	Red         = color.New(color.FgRed)
	Green       = color.New(color.FgGreen)
	Yellow      = color.New(color.FgYellow)
	Blue        = color.New(color.FgBlue)
	Magenta     = color.New(color.FgMagenta)
	Cyan        = color.New(color.FgCyan)
	White       = color.New(color.FgWhite)
	BlueBold    = color.New(color.Bold, color.FgBlue)
	MagentaBold = color.New(color.Bold, color.FgMagenta)
	RedBold     = color.New(color.Bold, color.FgRed)
	YellowBold  = color.New(color.Bold, color.FgYellow)
)

// Log Level
type Level int

const (
	Silent Level = iota
	Error
	Warn
	Info
)

// Writer log writer interface
type CustomLogger interface {
	Printf(string, ...interface{})
	Writer() io.Writer
}

// Logger interface
type Logger interface {
	Info(context.Context, string, ...interface{})
	Warn(context.Context, string, ...interface{})
	Error(context.Context, string, ...interface{})
	SetLevel(Level)
}

func NewLogger(customLogger CustomLogger, level Level, colorful bool) Logger {

	lg := logger{
		customLogger: customLogger,
		Colorful:     colorful,
		//info:         make([]*content, 0),
		//warn:         make([]*content, 0),
		//err:          make([]*content, 0),
		Level: level,
	}
	return &lg
}

type logger struct {
	customLogger    CustomLogger
	Colorful        bool
	Level           Level
	info, warn, err []*content
}

type content struct {
	text    string
	newLine bool
	color   *color.Color
}

//resetInfo resets info logger
func (l *logger) resetInfo() {
	l.info = make([]*content, 0)
	if l.Colorful {
		l.info = append(l.info, &content{"[info] ", false, Green})
	} else {
		l.info = append(l.info, &content{"[info] ", false, White})
	}
}

//resetWarn resets warn logger
func (l *logger) resetWarn() {
	l.warn = make([]*content, 0)
	if l.Colorful {
		l.warn = append(l.warn, &content{"[warn] ", false, Magenta})
	} else {
		l.warn = append(l.warn, &content{"[warn] ", false, White})
	}
}

//resetErr resets err logger
func (l *logger) resetErr() {
	l.err = make([]*content, 0)
	if l.Colorful {
		l.err = append(l.err, &content{"[error] ", false, Red})
	} else {
		l.err = append(l.err, &content{"[error] ", false, White})
	}
}

// Info print info
func (l *logger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.Level >= Info {
		l.resetInfo()
		l.info = append(l.info, &content{fmt.Sprintf(msg, data...), true, White})

		for _, item := range l.info {
			if item.color != nil {
				if item.newLine {
					item.color.Fprintln(l.customLogger.Writer(), item.text)
				} else {
					item.color.Fprint(l.customLogger.Writer(), item.text)
				}
			}
		}
	}
}

// Warn print warn messages
func (l *logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.Level >= Warn {
		l.resetWarn()
		l.warn = append(l.warn, &content{fmt.Sprintf(msg, data...), true, White})

		for _, item := range l.warn {
			if item.color != nil {
				if item.newLine {
					item.color.Fprintln(l.customLogger.Writer(), item.text)
				} else {
					item.color.Fprint(l.customLogger.Writer(), item.text)
				}
			}
		}
	}
}

// Error print error messages
func (l *logger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.Level >= Error {
		l.resetErr()
		l.err = append(l.err, &content{fmt.Sprintf(msg, data...), true, White})

		for _, item := range l.err {
			if item.color != nil {
				if item.newLine {
					item.color.Fprintln(l.customLogger.Writer(), item.text)
				} else {
					item.color.Fprint(l.customLogger.Writer(), item.text)
				}
			}
		}
	}
}

func (l *logger) SetLevel(lv Level) {
	l.Level = lv
}
