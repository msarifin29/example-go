package config

import (
	"fmt"
	"gorm-example/helper"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	c := NewViper()
	err := c.ReadInConfig()
	if err != nil {
		panic("failed to read config")
	}
	root := c.GetString("user_name")
	host := c.GetString("host")
	port := c.GetInt("port")
	dbName := c.GetString("database_name")
	dsn := fmt.Sprintf(`%s:@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local`, root, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		helper.LogFatal("failed to connect database", err)
	}
	connection, err := db.DB()
	if err != nil {
		helper.LogFatal("failed to connect database", err)
	}
	connection.SetMaxIdleConns(5)
	connection.SetMaxOpenConns(20)
	connection.SetConnMaxLifetime(60 * time.Minute)
	connection.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
