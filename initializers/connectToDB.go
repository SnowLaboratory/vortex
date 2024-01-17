package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbProtocol := os.Getenv("DBPROTOCOL")
	dbHost := os.Getenv("DBHOST")
	dbPort := os.Getenv("DBPORT")
	dbName := os.Getenv("DBNAME")

	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@%s(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbProtocol, dbHost, dbPort, dbName),
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the Database", err)
	}

}
