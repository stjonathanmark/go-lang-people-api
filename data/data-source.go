package data

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"stjonathanmark.com/people/models"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func NewDataSource(connectionString string) *DataSource {
	var db *gorm.DB
	var connErr error
	var attempts, maxAttempts, secondsBtwAttempts int64
	maxAttempts, _ = strconv.ParseInt(os.Getenv("MAX_DB_CONN_ATTEMPTS"), 0, 64)
	secondsBtwAttempts, _ = strconv.ParseInt(os.Getenv("SECONDS_BTW_DB_CONN_ATTEMPTS"), 0, 64)

	for attempts < maxAttempts {
		if connErr != nil {
			fmt.Printf("Connection attempt (%v of %v) failed. Next connection attempt in %v seconds.\n", attempts, maxAttempts, secondsBtwAttempts)
			time.Sleep(time.Duration(secondsBtwAttempts) * time.Second)
		}
		db, connErr = gorm.Open(sqlserver.Open(connectionString))
		if connErr == nil {
			break
		}
		attempts++
	}

	if connErr != nil {
		fmt.Printf("Connection attempt (%v of %v) failed. No future connection attempts.\n", attempts, maxAttempts)
		log.Panic("Connecting Error ", connErr)
	}

	migrationErr := db.AutoMigrate(&models.Person{})

	if migrationErr != nil {
		log.Panic("Migration Error ", migrationErr)
	}

	return &DataSource{NewPersonSource(db)}
}

type DataSource struct {
	models.PersonSource
}
