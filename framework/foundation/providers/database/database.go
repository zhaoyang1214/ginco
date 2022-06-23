package database

import (
	"ginco/framework/contract"
	"ginco/framework/database"
)

type Database struct {
}

func (d *Database) Build(container contract.Container, params ...interface{}) (interface{}, error) {
	appServer, err := container.Get("app")
	if err != nil {
		return nil, err
	}

	return database.NewDatabase(appServer.(contract.Application)), nil
}
