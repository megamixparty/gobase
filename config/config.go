package config

// EnvConfig create new mysql database instance and return it
type EnvConfig struct {
	Port int `env:"PORT,default=8787"`

	// MysqlConfig Struct contains mysql configuration
	Mysql struct {
		Username string `env:"MYSQL_USERNAME,required"`
		Password string `env:"MYSQL_PASSWORD,required"`
		Host     string `env:"MYSQL_HOST,default=localhost"`
		Port     string `env:"MYSQL_PORT,default=3306"`
		Name     string `env:"MYSQL_NAME,required"`
		Charset  string `env:"MYSQL_CHARSET,default=utf8mb4"`
		MaxIdle  int    `env:"MYSQL_MAX_IDLE,default=20"`
		MaxOpen  int    `env:"MYSQL_MAX_OPEN,default=100"`
	}
}
