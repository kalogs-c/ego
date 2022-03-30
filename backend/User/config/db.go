package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {	
	DB = newDB()
}

func newDB() *gorm.DB {
	dsn := "root:root@tcp(mysql)/ego?charset=utf8mb4\u0026readTimeout=30s\u0026writeTimeout=30s&parseTime=true"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB.Table("user")
}
