package main

import (
	"fmt"
	"log"
	"os"
	"todoApi/models"
	"todoApi/routers"
)

func main() {
	//local settings
	//dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
	//	os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	models := []interface{}{
		&models.User{},
		&models.Todo{},
	}
	db.AutoMigrate(models...)
	router := routers.Routes(db)
	router.Run("localhost:8080")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("database connection completed!")

}
