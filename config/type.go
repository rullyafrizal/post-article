package config

type Config struct {
	AppPort string `mapstructure:"APP_PORT"`

	MysqlHost     string `mapstructure:"MYSQL_HOST"`
	MysqlPort     string `mapstructure:"MYSQL_PORT"`
	MysqlDatabase string `mapstructure:"MYSQL_DATABASE"`
	MysqlUsername string `mapstructure:"MYSQL_USERNAME"`
	MysqlPassword string `mapstructure:"MYSQL_PASSWORD"`
}
