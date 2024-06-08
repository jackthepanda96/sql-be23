package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Genre struct {
	ID    int
	Name  string
	Books []Book `gorm:"foreignKey:GenreID"`
}

type Book struct {
	gorm.Model
	Title       string `gorm:"type:varchar(200)"`
	Author      string `gorm:"type:varchar(200)"`
	Publisher   string `gorm:"type:varchar(200)"`
	PublishYear uint8  `gorm:"column:taun;type:int8"`
	GenreID     int
}

func connectDB() (*gorm.DB, error) {
	var connStr = "host=aws-0-ap-southeast-1.pooler.supabase.com user=postgres.cihiokxhntapbfoqzmqu password=9Q6a5tOxJDA837m3 port=5432 dbname=postgres"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "be23.",
		},
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := connectDB()
	if err != nil {
		fmt.Println("terjadi masalah:", err.Error())
		return
	}

	db.AutoMigrate(&Genre{}, &Book{})

}
