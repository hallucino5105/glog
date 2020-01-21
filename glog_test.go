package glog

import "testing"

func TestGlog1(t *testing.T) {
	SetupLogger(WithOutput([]string{"stdout"}), WithVerbose(true))
	SetupLogger(WithOutput([]string{"stderr"}), WithVerbose(true))

	Debug("test debug")
	Info("test info")
	Debug("test warn")
}
