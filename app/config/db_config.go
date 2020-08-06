/*
@Time: 2020-08-03 10:29
@Auth: xujiang
@File: db.config
@Software: GoLand
@Desc: TODO
*/
package config

type dbConfig struct {
	Host     string `json:"Host" default:"localhost"`
	Port     int    `json:"Port" default:"3306"`
	UserName string `json:"UserName" default:"root"`
	PassWd   string `json:"PassWd" default:"root"`
	Database string `json:"Database" default:"rongshu_test"`
}

var DbConfig *dbConfig

func init() {
	registerConfig(DbConfig)
}

func (d *dbConfig) getConfigKeyName() (string, interface{}) {
	if DbConfig == nil {
		DbConfig = new(dbConfig)
	}
	return "Db", DbConfig
}
