package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"echo-template/config"
)

var DB *gorm.DB

func Init() {
	cfg := config.GetDBConfig()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.Dbname,
		cfg.Port,
		cfg.Zone,
	)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&TodoModel{})
	DB.AutoMigrate(&UserModel{})
}
