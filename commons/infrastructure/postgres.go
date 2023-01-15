package infrastructure

import (
	"fmt"
	"log"

	"github.com/Theoblanc/url-shortener/config"
	"github.com/Theoblanc/url-shortener/shortener"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GetPostgreSQLClient postgres connection
func GetPostgreSQLClient(config config.Interface) {
	host := config.Database().Host()
	user := config.Database().User()
	password := config.Database().Password()
	dbname := config.Database().Dbname()
	port := config.Database().Port()
	sslmode := config.Database().Sslmode()

	dsn := "host=" + host + " " + "user=" + user + " " + "password=" + password + " " + "dbname=" + dbname + " " + "port=" + port + " " + "sslmode=" + sslmode

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database!🤣 \n", err)
		panic("failed to connect database")
	}

	log.Println("Connected to database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)

	err = db.AutoMigrate(&shortener.Entity{})

	if err != nil {
		fmt.Println(err)
	}
}
