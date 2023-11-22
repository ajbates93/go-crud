package initialisers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	// dsn := os.Getenv("DB_STRING")
	dsn := "host=localhost user=alexbates password=postgres dbname=go-crud port=5432"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database.")
	}
}
