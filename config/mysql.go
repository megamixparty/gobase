package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Mysql driver
	"github.com/jmoiron/sqlx"
)

// NewMysql create new mysql database instance and return it
func NewMysql(cfg *EnvConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		cfg.Mysql.Username,
		cfg.Mysql.Password,
		cfg.Mysql.Host,
		cfg.Mysql.Port,
		cfg.Mysql.Name,
		cfg.Mysql.Charset))
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.Mysql.MaxIdle)
	db.SetMaxOpenConns(cfg.Mysql.MaxOpen)
	err = db.Ping()
	return db, err
}
