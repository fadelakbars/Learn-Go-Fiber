package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
var DB *gorm.DB
func InitDatabase()  {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/go_fiber_satu?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Konesi ke database gagal")
	}
	fmt.Println("Koneksi ke database berhasil")

}