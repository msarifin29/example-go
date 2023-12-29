package config

import (
	"fmt"

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
		panic("failed to connect database")
	}

	return db
}
