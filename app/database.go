package main

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type BeerQueue struct {
	gorm.Model
	Name   string
	Count  int
	Status string
}

func IntDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Automatically create the "accounts" table based on the `Account`
	// model.
	db.AutoMigrate(&BeerQueue{})
	return db, nil
}

func AddToQueue(db *gorm.DB, name string, count int) error {
	if err := db.Create(&BeerQueue{Name: name, Count: count, Status: "registered"}).Error; err != nil {
		return err
	}
	return nil
}

func GetQueue(db *gorm.DB, name string) []BeerQueue {
	var queue []BeerQueue
	db.Find(&queue, "name = ?", name)
	log.Println(queue)
	return queue
}
