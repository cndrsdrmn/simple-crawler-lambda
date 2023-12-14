package database

import (
	cfg "exchangerate/config"
	"exchangerate/model"
	"fmt"
	"log"

	driver "github.com/go-sql-driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Instance *gorm.DB
var TABLES []interface{} = []interface{}{
	&model.ExchangeRate{},
}

func init() {
	option := &driver.Config{
		User:      cfg.DB.User,
		Passwd:    cfg.DB.Password,
		DBName:    cfg.DB.Name,
		Addr:      fmt.Sprintf("%s:%d", cfg.DB.Host, cfg.DB.Port),
		Net:       "tcp",
		ParseTime: true,
		Params: map[string]string{
			"charset": "utf8mb4",
		},
	}

	var dbg logger.Interface
	if cfg.App.Debug {
		dbg = logger.Default.LogMode(logger.Info)
	}

	conn, err := gorm.Open(mysql.New(mysql.Config{
		DSNConfig: option,
	}), &gorm.Config{
		Logger: dbg,
	})

	if err != nil {
		txt := fmt.Sprintf("Error connected to database with setting: %v", err)
		panic(txt)
	}

	log.Println("Connected to database with setting:", option.FormatDSN())

	if err := conn.AutoMigrate(TABLES...); err != nil {
		panic(err.Error())
	}

	Instance = conn
}
