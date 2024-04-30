package database

import (
	"CRUD/configs"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mySqlDB struct {
	Client *gorm.DB
}

func NewMySqlDB(conf configs.Mysql, secrets configs.MysqlScretes) (*mySqlDB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?checkConnLiveness=false&loc=Local&parseTime=true&readTimeout=%s&timeout=%s&writeTimeout=%s&maxAllowedPacket=0",
		secrets.Username,
		secrets.Password,
		secrets.Host,
		conf.Database,
		conf.TimeOut,
		conf.TimeOut,
		conf.TimeOut,
	)

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, errors.Wrap(err, "gorm open")
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.Wrap(err, "get db")
	}

	sqlDB.SetMaxIdleConns(conf.Max.IdleConns)
	sqlDB.SetMaxOpenConns(conf.Max.OpenConns)
	sqlDB.SetConnMaxLifetime(conf.Max.Lifetime)

	return &mySqlDB{db}, nil
}

func (db *mySqlDB) Close() error {
	mysql, err := db.Client.DB()
	if err != nil {
		return errors.Wrap(err, "client db")
	}
	if err := mysql.Close(); err != nil {
		return errors.Wrap(err, "mysql close")
	}
	return nil
}

func (db *mySqlDB) Ping() error {
	sql, err := db.Client.DB()
	if err != nil {
		return errors.Wrap(err, "mysql")
	}
	err = sql.Ping()
	return errors.Wrap(err, "mysql")
}
