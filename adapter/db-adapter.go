package adapter

import (
	"fmt"
	config "github.com/arief-hidayat/gin-gorm-api/config"
	"github.com/arief-hidayat/gin-gorm-api/internal/model"
	"github.com/caarlos0/env/v9"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectWithDB() (*gorm.DB, *gorm.DB) {
	dbConfig := config.DbConfig{}
	if err := env.Parse(&dbConfig); err != nil {
		fmt.Printf("%+v\n", err)
	}
	writerDb, err := gorm.Open(postgres.Open(dbConfig.GetWriterConnStr()), &gorm.Config{})
	readerDb, readerDbErr := gorm.Open(postgres.Open(dbConfig.GetReaderConnStr()), &gorm.Config{})
	if err != nil || readerDbErr != nil {
		panic("Failed to create connection with database")
	}
	writerDb.AutoMigrate(&model.Contact{})
	return writerDb, readerDb
}

func CloseDbConnection(db *gorm.DB) {
	dbSql, err := db.DB()
	if err != nil {
		panic("Failed to close connection with Database")
	}
	dbSql.Close()
}
