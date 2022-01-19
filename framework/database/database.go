package database

import (
	"github.com/zhaoyang1214/ginco/framework/contract"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"time"
)

type Database struct {
	*gorm.DB
	app         contract.Application
	connections map[string]*gorm.DB
}

var _ contract.Database = (*Database)(nil)

func NewDatabase(app contract.Application) *Database {
	db := &Database{
		app:         app,
		connections: make(map[string]*gorm.DB),
	}
	defaultName := app.GetIgnore("config").(contract.Config).GetString("database.default")
	db.DB = db.Connection(defaultName)
	return db
}

func (db *Database) Connection(names ...string) *gorm.DB {
	name := "mysql"
	if len(names) > 0 {
		name = names[0]
	}
	if c, ok := db.connections[name]; ok {
		return c
	}
	db.connections[name] = db.Resolve(name)
	return db.connections[name]
}

func (db *Database) Resolve(name string) *gorm.DB {
	conf := db.app.GetIgnore("config").(contract.Config).Sub("database.connections." + name)
	if conf == nil {
		panic("Database config [" + name + "] is not defined")
	}
	var connection *gorm.DB
	driver := conf.GetString("driver")
	switch driver {
	case "mysql":
		connection = db.resolveMysql(conf)
	case "sqlite":
		connection = db.resolveSqlite(conf)
	case "sqlserver":
		connection = db.resolveSqlserver(conf)
	case "postgres":
		connection = db.resolvePostgres(conf)
	default:
		panic("Database driver [" + driver + "] is not supported")

	}

	return connection
}

func (db *Database) resolveMysql(conf contract.Config) *gorm.DB {
	var dsn string
	var sources, replicas []gorm.Dialector
	if conf.Has("dsn") {
		dsn = conf.GetString("dsn")
	} else if conf.Has("write") && conf.Has("read") {
		writeConf := conf.Get("write").([]interface{})
		for i, value := range writeConf {
			v := value.(map[interface{}]interface{})
			if i == 0 {
				dsn = v["dsn"].(string)
			}
			sources = append(sources, mysql.Open(v["dsn"].(string)))
		}

		readConf := conf.Get("read").([]interface{})
		for _, value := range readConf {
			v := value.(map[interface{}]interface{})
			replicas = append(replicas, mysql.Open(v["dsn"].(string)))
		}
	} else {
		panic("Database dsn is not found")
	}

	mysqlConfig := mysql.Config{
		DSN: dsn,
	}
	// TODO 后续版本支持其他参数

	conn, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if len(sources) > 0 || len(replicas) > 0 {
		resolver := dbresolver.Register(dbresolver.Config{
			Sources:  sources,
			Replicas: replicas,
			Policy:   dbresolver.RandomPolicy{},
		})

		if conf.Has("conn_max_idle_time") {
			resolver = resolver.SetConnMaxIdleTime(conf.GetDuration("conn_max_idle_time") * time.Hour)
		}

		if conf.Has("conn_max_lifetime") {
			resolver = resolver.SetConnMaxLifetime(conf.GetDuration("conn_max_lifetime") * 24 * time.Hour)
		}

		if conf.Has("max_idle_conns") {
			resolver = resolver.SetMaxIdleConns(conf.GetInt("conn_max_lifetime"))
		}

		if conf.Has("max_open_conns") {
			resolver = resolver.SetMaxOpenConns(conf.GetInt("max_open_conns"))
		}

		err := conn.Use(resolver)
		if err != nil {
			panic(err)
		}
	} else {
		sqlDB, err := conn.DB()
		if err != nil {
			panic(err)
		}
		if conf.Has("conn_max_idle_time") {
			sqlDB.SetConnMaxIdleTime(conf.GetDuration("conn_max_idle_time") * time.Hour)
		}

		if conf.Has("conn_max_lifetime") {
			sqlDB.SetConnMaxLifetime(conf.GetDuration("conn_max_lifetime") * 24 * time.Hour)
		}

		if conf.Has("max_idle_conns") {
			sqlDB.SetMaxIdleConns(conf.GetInt("conn_max_lifetime"))
		}

		if conf.Has("max_open_conns") {
			sqlDB.SetMaxOpenConns(conf.GetInt("max_open_conns"))
		}
	}

	return conn
}

func (db *Database) resolveSqlite(conf contract.Config) *gorm.DB {
	return nil
}

func (db *Database) resolveSqlserver(conf contract.Config) *gorm.DB {
	return nil
}

func (db *Database) resolvePostgres(conf contract.Config) *gorm.DB {
	return nil
}
