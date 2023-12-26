package app

import (
	"gorm-example/helper"
	"strconv"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var config = viper.New()

func OpenConnection() *gorm.DB {
	config.SetConfigName("db_dev")
	config.SetConfigType("yaml")
	config.AddConfigPath("/Users/asams/personal/gorm-example/.")

	err := config.ReadInConfig()
	helper.PanicIfError(err)

	root := config.GetString("database.user_name")
	host := config.GetString("database.host")
	port := config.GetInt("database.port")
	dbName := config.GetString("database.database_name")
	portString := strconv.Itoa(port)

	var dsn string = root + ":@tcp(" + host + ":" + portString + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)
	helper.LogInfo("open DB:", err)

	sqlDB, err := db.DB()
	helper.PanicIfError(err)
	helper.LogInfo(" DB :", err)

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	return db
}
