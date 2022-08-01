package database

import (
	"fmt"

	"github.com/MaestroShifu/golang-fiber/src/infrastructure/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func parseToDSN(host string, user string, password string, dbname string, port int) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/Bogota", host, user, password, dbname, port)
}

func configDatabase() gorm.Dialector {
	configApp := config.GetConfigSystem()
	configPostgreSQL := postgres.Config{
		DSN: parseToDSN(configApp.DB_HOST, configApp.DB_USER, configApp.POSTGRES_PASSWORD, configApp.DB_NAME, 5432),
	}
	return postgres.New(configPostgreSQL)
}

func Connect() (*gorm.DB, error) {
	dataBaseConfig := configDatabase()
	db, err := gorm.Open(dataBaseConfig, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("Error connect to db:", err)
		return nil, err
	}
	return db, nil
}
