package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Level type
type Level uint

// log levels
const (
	LevelCritical Level = iota
	LevelError
	LevelWarning
	LevelTrace
	LevelDebug
	LevelInfo
)

// log flags
const (
	Ldate         = log.Ldate
	Ltime         = log.Ltime
	Lmicroseconds = log.Lmicroseconds
	Llongfile     = log.Llongfile
	Lshortfile    = log.Lshortfile
	LUTC          = log.LUTC
	LstdFlags     = log.LstdFlags
	Llevel        = LUTC << iota
	Lshortlevel
	Lcolor
)

var colors = map[Level]string{
	LevelCritical: "\033[0;31m",
	LevelError:    "\033[0;31m",
	LevelWarning:  "\033[0;33m",
	LevelTrace:    "\033[0;35m",
	LevelDebug:    "\033[0;36m",
	LevelInfo:     "\033[0;37m",
}

var titles = map[Level]string{
	LevelCritical: "CRIT ",
	LevelError:    "ERR  ",
	LevelWarning:  "WARN ",
	LevelTrace:    "TRACE",
	LevelDebug:    "DEBUG",
	LevelInfo:     "INFO",
}

const resetColor = "\033[0m"

// Logger simple logger wrapper
type Logger struct {
	*log.Logger
	level Level
	out   io.Writer
}

// New creates a new Logger.
func New(out io.Writer, prefix string, flag int, level Level) *Logger {
	return &Logger{
		log.New(out, prefix, flag),
		level,
		out,
	}
}

//SetLevel setting logger level
func (l *Logger) SetLevel(level Level) {
	l.level = level
}

//Info with args
func (l *Logger) Info(args ...interface{}) {
	l.log(LevelInfo, args...)
}

//Infof with format and args
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logf(LevelInfo, format, args...)
}

//Debug with args
func (l *Logger) Debug(args ...interface{}) {
	l.log(LevelDebug, args...)
}

//Debugf with format and args
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logf(LevelDebug, format, args...)
}

//Trace with args
func (l *Logger) Trace(args ...interface{}) {
	l.log(LevelTrace, args...)
}

//Tracef with format and args
func (l *Logger) Tracef(format string, args ...interface{}) {
	l.logf(LevelTrace, format, args...)
}

//Warn with args
func (l *Logger) Warn(args ...interface{}) {
	l.log(LevelWarning, args...)
}

//Warnf with format and args
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logf(LevelWarning, format, args...)
}

//Error with format and args
func (l *Logger) Error(args ...interface{}) {
	l.log(LevelError, args...)
}

//Errorf with format and args
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logf(LevelError, format, args...)
}

//Crit with args
func (l *Logger) Crit(args ...interface{}) {
	l.log(LevelError, args...)
	panic(fmt.Sprint(args...))
}

//Critf with format and args
func (l *Logger) Critf(format string, args ...interface{}) {
	l.logf(LevelError, format, args...)
	panic(fmt.Sprint(args...))
}

//SetFlags (see log package)
func (l *Logger) SetFlags(flag int) { l.Logger.SetFlags(flag) }

//SetOutput (see log package)
func (l *Logger) SetOutput(w io.Writer) { l.Logger.SetOutput(w) }

//SetPrefix (see log package)
func (l *Logger) SetPrefix(prefix string) { l.Logger.SetPrefix(prefix) }

//Flags (see log package)
func (l *Logger) Flags() int { return l.Logger.Flags() }

//Output (see log package)
func (l *Logger) Output(calldepth int, s string) error { return l.Logger.Output(calldepth, s) }

//Prefix (see log package)
func (l *Logger) Prefix() string { return l.Logger.Prefix() }

func (l *Logger) log(level Level, args ...interface{}) {
	l.logf(level, strings.TrimRight(strings.Repeat("%v ", len(args)), " "), args...)
}

func (l *Logger) logf(level Level, format string, args ...interface{}) {
	if l.level < level {
		return
	}
	if l.Flags()&Lcolor != 0 {
		fmt.Fprintf(l.out, colors[level])
	}
	var levelString string
	if l.Flags()&Llevel != 0 {
		levelString = titles[level] + " "
	}
	if l.Flags()&Lshortlevel != 0 {
		levelString = titles[level][0:1] + " "
	}

	l.Printf("%s%s", levelString, fmt.Sprintf(format, args...))
	if l.Flags()&Lcolor != 0 {
		fmt.Fprintf(l.out, resetColor)
	}
}

var std = New(os.Stderr, "", LstdFlags|Llevel|Lcolor, LevelInfo)

//SetLevel setting logger level
func SetLevel(level Level) {
	std.SetLevel(level)
}

//Info with args
func Info(args ...interface{}) {
	std.Info(args...)
}

//Infof with format and args
func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}

//Debug with args
func Debug(args ...interface{}) {
	std.Debug(args...)
}

//Debugf with format and args
func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}

//Trace with args
func Trace(args ...interface{}) {
	std.Trace(args...)
}

//Tracef with format and args
func Tracef(format string, args ...interface{}) {
	std.Tracef(format, args...)
}

//Warn with args
func Warn(args ...interface{}) {
	std.Warn(args...)
}

//Warnf with format and args
func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}

//Error with args
func Error(args ...interface{}) {
	std.Error(args...)
}

//Errorf with format and args
func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}

//Crit with args
func Crit(args ...interface{}) {
	std.Crit(args...)
}

//Critf with format and args
func Critf(format string, args ...interface{}) {
	std.Critf(format, args...)
}

//SetFlags (see log package)
func SetFlags(flag int) { std.Logger.SetFlags(flag) }

//SetOutput (see log package)
func SetOutput(w io.Writer) { std.Logger.SetOutput(w) }

//SetPrefix (see log package)
func SetPrefix(prefix string) { std.Logger.SetPrefix(prefix) }

//Flags (see log package)
func Flags() int { return std.Logger.Flags() }

//Output (see log package)
func Output(calldepth int, s string) error { return std.Logger.Output(calldepth, s) }

//Prefix (see log package)
func Prefix() string { return std.Logger.Prefix() }
