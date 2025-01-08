package core

import (
	"github.com/sirupsen/logrus"
)

type Fields logrus.Fields

type appLogger struct {
	logger   *logrus.Logger
	cfg      Config
	logLevel string
}

func (a appLogger) ID() string {
	//TODO implement me
	panic("implement me")
}

func (a appLogger) Activate() error {
	//TODO implement me
	panic("implement me")
}

func (a appLogger) Stop() error {
	//TODO implement me
	panic("implement me")
}

type Config struct {
	DefaultLevel string
	BasePrefix   string
}

type Logger interface {
	Print(args ...interface{})
	Debug(...interface{})
	Debugln(...interface{})
	Debugf(string, ...interface{})

	Info(...interface{})
	Infoln(...interface{})
	Infof(string, ...interface{})

	Warn(...interface{})
	Warnln(...interface{})
	Warnf(string, ...interface{})

	Error(...interface{})
	Errorln(...interface{})
	Errorf(string, ...interface{})

	Fatal(...interface{})
	Fatalln(...interface{})
	Fatalf(string, ...interface{})

	Panic(...interface{})
	Panicln(...interface{})
	Panicf(string, ...interface{})

	With(key string, value interface{}) Logger
	Withs(Fields) Logger
	// add source field to log
	WithSrc() Logger
	GetLevel() string
}

type logger struct {
	*logrus.Entry
}

func (l *logger) With(key string, value interface{}) Logger {
	//TODO implement me
	panic("implement me")
}

func (l *logger) Withs(f Fields) Logger {
	//TODO implement me
	panic("implement me")
}

func (l *logger) WithSrc() Logger {
	//TODO implement me
	panic("implement me")
}

func (l *logger) GetLevel() string {
	//TODO implement me
	panic("implement me")
}
