package glog

import (
	"sync"

	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

type LoggerOption struct {
	Output  []string
	Verbose bool
}

func SetupLogger(ofs ...func(*LoggerOption)) {
	options := setOptions(ofs)

	config := zap.NewDevelopmentConfig()
	config.OutputPaths = options.Output

	if options.Verbose {
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	} else {
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	zl, _ := config.Build()
	Logger = zl.Sugar()
}

func WithOutput(output []string) func(*LoggerOption) {
	return func(options *LoggerOption) {
		options.Output = output
	}
}

func WithVerbose(verbose bool) func(*LoggerOption) {
	return func(options *LoggerOption) {
		options.Verbose = verbose
	}
}

func setOptions(ofs []func(*LoggerOption)) *LoggerOption {
	options := LoggerOption{
		Output:  []string{"stderr"},
		Verbose: false,
	}

	for _, of := range ofs {
		of(&options)
	}

	return &options
}

func Debug(args ...interface{}) {
	if Logger != nil {
		Logger.Debug(args...)
	}
}

func LDebug(mu *sync.Mutex, args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	if Logger != nil {
		Logger.Debug(args...)
	}
}

func Info(args ...interface{}) {
	if Logger != nil {
		Logger.Info(args...)
	}
}

func LInfo(mu *sync.Mutex, args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	if Logger != nil {
		Logger.Info(args...)
	}
}

func Warn(args ...interface{}) {
	if Logger != nil {
		Logger.Warn(args...)
	}
}

func LWarn(mu *sync.Mutex, args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	if Logger != nil {
		Logger.Warn(args...)
	}
}

func Error(args ...interface{}) {
	if Logger != nil {
		Logger.Error(args...)
	}
}

func LError(mu *sync.Mutex, args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	if Logger != nil {
		Logger.Error(args...)
	}
}

func Fatal(args ...interface{}) {
	if Logger != nil {
		Logger.Fatal(args...)
	}
}

func LFatal(mu *sync.Mutex, args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	if Logger != nil {
		Logger.Fatal(args...)
	}
}

func Panic(args ...interface{}) {
	if Logger != nil {
		Logger.Panic(args...)
	}
}

func LPanic(mu *sync.Mutex, args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	if Logger != nil {
		Logger.Panic(args...)
	}
}

func Debugf(template string, args ...interface{}) {
	if Logger != nil {
		Logger.Debugf(template, args...)
	}
}

func LDebugf(mu *sync.Mutex, template string, args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	if Logger != nil {
		Logger.Debugf(template, args...)
	}
}

func Infof(template string, args ...interface{}) {
	if Logger != nil {
		Logger.Infof(template, args...)
	}
}

func LInfof(mu *sync.Mutex, template string, args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	if Logger != nil {
		Logger.Infof(template, args...)
	}
}

func Warnf(template string, args ...interface{}) {
	if Logger != nil {
		Logger.Warnf(template, args...)
	}
}

func LWarnf(mu *sync.Mutex, template string, args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	if Logger != nil {
		Logger.Warnf(template, args...)
	}
}

func Errorf(template string, args ...interface{}) {
	if Logger != nil {
		Logger.Errorf(template, args...)
	}
}

func LErrorf(mu *sync.Mutex, template string, args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	if Logger != nil {
		Logger.Errorf(template, args...)
	}
}

func Fatalf(template string, args ...interface{}) {
	if Logger != nil {
		Logger.Fatalf(template, args...)
	}
}

func LFatalf(mu *sync.Mutex, template string, args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	if Logger != nil {
		Logger.Fatalf(template, args...)
	}
}

func Panicf(template string, args ...interface{}) {
	if Logger != nil {
		Logger.Panicf(template, args...)
	}
}

func LPanicf(mu *sync.Mutex, template string, args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	if Logger != nil {
		Logger.Panicf(template, args...)
	}
}
