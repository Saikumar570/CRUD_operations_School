package utils

import (
	"School_api/models"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=5152 dbname=users port=5432 sslmode=disable"
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrating
	// if err := db.AutoMigrate(&models.School{}).Error; err != nil {
	// 	log.Fatal("Failed to migrate database:", err)
	// }
	if err := db.AutoMigrate(&models.School{}, &models.Student{}).Error; err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("Database connected and migrated")
	return db
}
