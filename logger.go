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
	Ldate = 1 << iota
	Ltime
	Lmicroseconds
	Llongfile
	Lshortfile
	LUTC
	Llevel
	Lshortlevel
	Lcolor
	LstdFlags = Ldate | Ltime
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

var logger *Logger

// Logger simple logger wrapper
type Logger struct {
	*log.Logger
	level Level
	out   io.Writer
}

// New creates a new Logger.
func New(out io.Writer, prefix string, flag int, args ...interface{}) *Logger {
	level := LevelInfo
	if len(args) > 0 {
		if l, ok := args[0].(Level); ok {
			level = l
		}
	}
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

var std = New(os.Stderr, "", LstdFlags|Llevel|Lcolor)

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
