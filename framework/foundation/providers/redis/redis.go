package redis

import (
	"ginco/framework/contract"
	"ginco/framework/redis"
)

type Redis struct {
}

var _ contract.Provider = (*Redis)(nil)

func (r *Redis) Build(container contract.Container, params ...interface{}) (interface{}, error) {
	appServer, err := container.Get("app")
	if err != nil {
		return nil, err
	}

	return redis.NewRedis(appServer.(contract.Application)), nil
}
