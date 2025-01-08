package core

import "github.com/sirupsen/logrus"

const (
	DevEnv  = "dev"
	StgEnv  = "stg"
	ProdEnv = "prod"
)

type Option func(*serviceContext)

type Component interface {
	ID() string
	Activate() error
	Stop() error
}

var (
	defaultLogger = newAppLogger(&Config{
		BasePrefix:   "core",
		DefaultLevel: "trace",
	})
)

func newAppLogger(config *Config) *appLogger {
	if config == nil {
		config = &Config{}
	}

	if config.DefaultLevel == "" {
		config.DefaultLevel = "info"
	}

	logger := logrus.New()
	logger.Formatter = logrus.Formatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return &appLogger{
		logger:   logger,
		cfg:      *config,
		logLevel: config.DefaultLevel,
	}
}

type ServiceContext interface {
	Load() error
	MustGet(id string) interface{}
	Get(id string) interface{}
	Logger(prefix string) Logger
	EnvName() string
	GetName() string
}

type serviceContext struct {
	name       string
	env        string
	components []Component
	store      map[string]Component
	logger     Logger
}

func (s serviceContext) Logger(prefix string) Logger {
	//TODO implement me
	panic("implement me")
}

func (s serviceContext) EnvName() string {
	//TODO implement me
	panic("implement me")
}

func (s serviceContext) GetName() string {
	//TODO implement me
	panic("implement me")
}

func (s serviceContext) Load() error {
	//TODO implement me
	panic("implement me")
}

func (s serviceContext) MustGet(id string) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s serviceContext) Get(id string) interface{} {
	//TODO implement me
	panic("implement me")
}

func NewServiceContext(opts ...Option) ServiceContext {
	sv := &serviceContext{
		store: make(map[string]Component),
	}

	sv.components = []Component{defaultLogger}
	return sv
}
