package logger

import (
	"os"
	"testing"
)

func TestDefault(t *testing.T) {
	Info("Info", 1)
	Infof("Infof %d", 1)
	Debug("Debug", 2)
	Debugf("Debugf %d", 1)
	Trace("Trace", 3)
	Tracef("Tracef %d", 1)
	Warn("Warn", 4)
	Warnf("Warnf %d", 1)
	Error("Error", 5)
	Errorf("Errorf %d", 1)
}

func TestShortFileFlags(t *testing.T) {
	std.SetFlags(Lshortfile | LstdFlags | Lshortlevel)
	Info("Info", 1)
	Infof("Infof %d", 1)

	log := New(os.Stderr, "", Lshortfile|LstdFlags|Lshortlevel, LevelInfo)
	log.Info("Info", 1)
	log.Infof("Info %d", 1)
}

func TestLevel(t *testing.T) {
	SetLevel(LevelWarning)
	Info("Info", 1)
	Infof("Infof %d", 1)
	Debug("Debug", 2)
	Debugf("Debugf %d", 1)
	Trace("Trace", 3)
	Tracef("Tracef %d", 1)
	Warn("Warn", 4)
	Warnf("Warnf %d", 1)
	Error("Error", 5)
	Errorf("Errorf %d", 1)
}

func TestPrefix(t *testing.T) {
	Debug("test")
	x := NewPrefixed("[%s]", "hello")

	x.Debug("Foo")
	x2 := x.NewPrefixed("[%s]", "xxx")

	x2.Debug("Hello")
}
