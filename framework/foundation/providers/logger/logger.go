package logger

import (
	"ginco/framework/contract"
	"ginco/framework/logger"
)

type Logger struct {
}

var _ contract.Provider = (*Logger)(nil)

func (l *Logger) Build(container contract.Container, params ...interface{}) (interface{}, error) {
	appServer, err := container.Get("app")
	if err != nil {
		return nil, err
	}

	log := logger.NewManager(appServer.(contract.Application))
	return log, nil
}
