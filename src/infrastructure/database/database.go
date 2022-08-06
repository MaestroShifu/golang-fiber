package database

import (
	"fmt"
	"log"

	"github.com/MaestroShifu/golang-fiber/src/infrastructure/config"
	"github.com/MaestroShifu/golang-fiber/src/infrastructure/database/models"
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

func AutoMigrate(db *gorm.DB) {
	dbName := db.Migrator().CurrentDatabase()
	log.Println("start auto migration in ", dbName)

	// Install complement working generate uuid
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	db.AutoMigrate(&models.ProductModel{})
}

func Connect() (*gorm.DB, error) {
	dataBaseConfig := configDatabase()
	db, err := gorm.Open(dataBaseConfig, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("error connect to db:", err)
		return nil, err
	}
	return db, nil
}
