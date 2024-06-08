package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Setting struct {
	Host     string
	User     string
	Password string
	Port     string
	DBNAME   string
}

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

func connectDB(s Setting) (*gorm.DB, error) {
	var connStr = fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s", s.Host, s.User, s.Password, s.Port, s.DBNAME)
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

func importSetting() Setting {
	var result Setting
	err := godotenv.Load(".env")
	if err != nil {
		return Setting{}
	}
	result.Host = os.Getenv("poshost")
	result.User = os.Getenv("posuser")
	result.Password = os.Getenv("pospw")
	result.Port = os.Getenv("posport")
	result.DBNAME = os.Getenv("dbname")
	return result
}

func main() {
	setting := importSetting()
	db, err := connectDB(setting)
	if err != nil {
		fmt.Println("terjadi masalah:", err.Error())
		return
	}

	db.AutoMigrate(&Genre{}, &Book{})

}
