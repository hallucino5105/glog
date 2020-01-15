package glog

import "testing"

func TestGlog1(t *testing.T) {
	SetupLogger(true)

	Debug("test debug")
	Info("test info")
	Debug("test warn")
}
