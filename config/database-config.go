package config

import (
	"fmt"
)

type DbConfig struct {
	WriterEndpoint string `env:"DB_WRITER_ENDPOINT" envDefault:"127.0.0.1"`
	ReaderEndpoint string `env:"DB_READER_ENDPOINT" envDefault:"127.0.0.1"`
	User           string `env:"DB_USER" envDefault:"postgres"`
	Password       string `env:"DB_PWD,unset"`
	Dbname         string `env:"DB_NAME" envDefault:"service"`
	Port           int    `env:"DB_PORT" envDefault:"5432"`
	Sslmode        string `env:"DB_SSL_MODE" envDefault:"disable"`
	Timezone       string `env:"DB_TIMEZONE" envDefault:"UTC"`
}

func (conf *DbConfig) GetReaderConnStr() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		conf.ReaderEndpoint,
		conf.User,
		conf.Password,
		conf.Dbname,
		conf.Port,
		conf.Sslmode,
		conf.Timezone,
	)
}

func (conf *DbConfig) GetWriterConnStr() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		conf.WriterEndpoint,
		conf.User,
		conf.Password,
		conf.Dbname,
		conf.Port,
		conf.Sslmode,
		conf.Timezone,
	)
}
