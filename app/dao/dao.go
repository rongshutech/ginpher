/*
@Time: 2020-08-05 11:29
@Auth: Yuewei
@File: db.config
@Software: GoLand
@Desc: TODO
*/
package dao

import (
	"fmt"
	"ginpher/app/config"

	orm "github.com/jinzhu/gorm"

	"ginpher/app/tools/database/gorm"

	"github.com/vgmdj/plugins/redis"
	"github.com/vgmdj/utils/logger"
)

type Dao struct {
	db    *orm.DB
	redis *redis.Client
}

var daoHelper *Dao

func NewDao() *Dao {
	if daoHelper == nil || daoHelper.Ping() != nil {

		dbSettings := &gorm.Config{
			DSN: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
				config.DbConfig.UserName,
				config.DbConfig.PassWd,
				config.DbConfig.Host,
				config.DbConfig.Port,
				config.DbConfig.Database),
			Active:      100,
			Idle:        100,
			IdleTimeout: 10,
		}
		redisSettings := &redis.ClientConf{
			Address:  config.RedisConfig.Address,
			Password: config.RedisConfig.PassWd,
			DB:       config.RedisConfig.DB,
		}
		daoHelper = &Dao{
			db:    gorm.NewMysqlOrm(dbSettings),
			redis: redis.NewClient(redisSettings),
		}
	}
	return daoHelper
}

func (d *Dao) Ping() error {
	err := d.db.DB().Ping()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return d.redis.Ping()
}
