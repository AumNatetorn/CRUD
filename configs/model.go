package configs

import (
	"time"

	"github.com/go-playground/validator"
)

type Config struct {
	App     App     `mapstructure:"app"`
	Mysql   Mysql   `mapstructure:"mysql"`
	Secrets Secrets `mapstructure:"secrets"`
}

func (c Config) Validate() error {
	return validator.New().Struct(c)
}

type App struct {
	Name    string        `mapstructure:"name" validate:"required"`
	TimeOut time.Duration `mapstructure:"timeout" validate:"required"`
	Port    int           `mapstructure:"port" validate:"required"`
}

type Mysql struct {
	Database string        `mapstructure:"database" validate:"required"`
	TimeOut  time.Duration `mapstructure:"timeout" validate:"gt=0"`
	Max      MysqlMax      `mapstructure:"max"`
}

type MysqlMax struct {
	IdleConns int           `mapstructure:"idle_conns" validate:"gt=0"`
	OpenConns int           `mapstructure:"open_conns" validate:"gt=0"`
	Lifetime  time.Duration `mapstructure:"lifetime" validate:"gt=0"`
}

type Secrets struct {
	SecretsMysql MysqlScretes `mapstructure:"mysql"`
}

type MysqlScretes struct {
	Host     string `mapstructure:"host" validate:"required"`
	Username string `mapstructure:"username" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
}
