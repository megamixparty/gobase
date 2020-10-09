package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Mysql driver
)

// MysqlConfig Struct contains mysql configuration
type MysqlConfig struct {
	Username string `env:"MYSQL_USERNAME,required"`
	Password string `env:"MYSQL_PASSWORD,required"`
	Host     string `env:"MYSQL_HOST,default=localhost"`
	Port     string `env:"MYSQL_PORT,default=3306"`
	Name     string `env:"MYSQL_NAME,required"`
	Charset  string `env:"MYSQL_CHARSET,default=utf8mb4"`
	MaxIdle  int    `env:"MYSQL_MAX_IDLE,default=20"`
	MaxOpen  int    `env:"MYSQL_MAX_OPEN,default=100"`
}

// NewMysql create new mysql database instance and return it
func NewMysql(cfg *MysqlConfig) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.Charset))
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.MaxIdle)
	db.SetMaxOpenConns(cfg.MaxOpen)
	err = db.Ping()
	return db, err
}
