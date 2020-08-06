package gorm

import (
	"time"

	lt "ginpher/app/tools/time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/vgmdj/utils/logger"
)

type Config struct {
	DSN         string      `json:"dsn"`          // data source name.
	Active      int         `json:"active"`       // pool
	Idle        int         `json:"idle"`         // pool
	IdleTimeout lt.Duration `json:"idle_timeout"` // connect max life time.
}

func NewMysqlOrm(c *Config) (db *gorm.DB) {
	db, err := gorm.Open("mysql", c.DSN)
	if err != nil {
		logger.Error("db dsn(%s) error: %v", c.DSN, err)
		panic(err)
	}
	db.DB().SetMaxIdleConns(c.Idle)
	db.DB().SetMaxOpenConns(c.Active)
	db.DB().SetConnMaxLifetime(time.Duration(c.IdleTimeout) / time.Second)
	db.LogMode(true)
	return

}
