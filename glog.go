package glog

import "go.uber.org/zap"

var Logger *zap.SugaredLogger

func SetupLogger(verbose bool) {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stderr"}

	if verbose {
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	} else {
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	zl, _ := config.Build()
	Logger = zl.Sugar()
}

func Debug(args ...interface{}) {
	if Logger != nil {
		Logger.Debug(args...)
	}
}

func Info(args ...interface{}) {
	if Logger != nil {
		Logger.Info(args...)
	}
}

func Warn(args ...interface{}) {
	if Logger != nil {
		Logger.Warn(args...)
	}
}

func Error(args ...interface{}) {
	if Logger != nil {
		Logger.Error(args...)
	}
}

func Fatal(args ...interface{}) {
	if Logger != nil {
		Logger.Fatal(args...)
	}
}

func Panic(args ...interface{}) {
	if Logger != nil {
		Logger.Panic(args...)
	}
}

func Debugf(template string, args ...interface{}) {
	if Logger != nil {
		Logger.Debugf(template, args...)
	}
}

func Infof(template string, args ...interface{}) {
	if Logger != nil {
		Logger.Infof(template, args...)
	}
}

func Warnf(template string, args ...interface{}) {
	if Logger != nil {
		Logger.Warnf(template, args...)
	}
}

func Errorf(template string, args ...interface{}) {
	if Logger != nil {
		Logger.Errorf(template, args...)
	}
}

func Fatalf(template string, args ...interface{}) {
	if Logger != nil {
		Logger.Fatalf(template, args...)
	}
}

func Panicf(template string, args ...interface{}) {
	if Logger != nil {
		Logger.Panicf(template, args...)
	}
}
