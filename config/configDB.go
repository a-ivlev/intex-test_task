package config

import "os"

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func LoadConfigDB() *Config {
	return &Config{
		DB: &DBConfig{
			Host:     os.Getenv("DB_HOST"), //localhost
			Port:     os.Getenv("DB_PORT"), //5432
			Username: os.Getenv("DB_USER"), //"boris"
			Password: os.Getenv("DB_PASSWORD"), //"qwerty"
			Database: os.Getenv("DB_NAME"), //books
		},
	}
}
