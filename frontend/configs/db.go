package configs

import (
	"fmt"
	"os"
	"win/lab8/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection()  {
	dsn := os.Getenv("DB_DSN")
   
	db, err := gorm.Open(postgres.Open(dsn),&gorm.Config{})

	if err != nil{
		panic("fail to connect database")
	}

	fmt.Println("DB Connected !!")

	//migrate
	db.AutoMigrate(models.User{})

	DB = db

}
