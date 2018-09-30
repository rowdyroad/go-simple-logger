package logger

import (
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

func TestFlags(t *testing.T) {
	std.SetFlags(LstdFlags | Lshortlevel)
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
