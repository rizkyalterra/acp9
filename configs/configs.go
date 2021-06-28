package configs

import (
	"app/models/user"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type ConfigDB struct {
	DB_HOST          string
	DB_PORT          string
	DB_USERNAME      string
	DB_PASSWORD      string
	DB_DATABASE_NAME string
}

func InitDB() {
	config := ConfigDB{
		DB_HOST:          "127.0.0.1",
		DB_PORT:          "3306",
		DB_USERNAME:      "root",
		DB_PASSWORD:      "123ABC4d.",
		DB_DATABASE_NAME: "acp9",
	}
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USERNAME, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_DATABASE_NAME)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	AutoMigrate()
}

func AutoMigrate() {
	DB.AutoMigrate(&user.School{})
	DB.AutoMigrate(&user.User{})
}

func InitDBTest() {
	config := ConfigDB{
		DB_HOST:          "127.0.0.1",
		DB_PORT:          "3306",
		DB_USERNAME:      "root",
		DB_PASSWORD:      "123ABC4d.",
		DB_DATABASE_NAME: "acp9_test",
	}
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USERNAME, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_DATABASE_NAME)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	AutoMigrateTest()
}

func AutoMigrateTest() {
	DB.Migrator().DropTable(&user.User{})
	DB.AutoMigrate(&user.User{})
}
