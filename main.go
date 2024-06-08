package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Genre struct {
	ID   int
	Name string
}

func connectDB() (*gorm.DB, error) {
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"

	var connStr = "host=aws-0-ap-southeast-1.pooler.supabase.com user=postgres.cihiokxhntapbfoqzmqu password=9Q6a5tOxJDA837m3 port=5432 dbname=postgres"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "public.",
		},
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetAllGenre(db *gorm.DB) ([]Genre, error) {
	var result []Genre
	err := db.Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func InsertGenre(db *gorm.DB, name string) (bool, error) {
	var insetData = Genre{Name: name}
	qry := db.Create(&insetData)

	if err := qry.Error; err != nil {
		return false, err
	}

	if qry.RowsAffected < 1 {
		return false, gorm.ErrInvalidValueOfLength
	}

	return true, nil
}

func main() {
	db, err := connectDB()
	if err != nil {
		fmt.Println("terjadi masalah:", err.Error())
		return
	}

	success, _ := InsertGenre(db, "drama india")
	fmt.Println(success)
	//  mulai ngoding
	hasil, _ := GetAllGenre(db)
	fmt.Println(hasil)

}
