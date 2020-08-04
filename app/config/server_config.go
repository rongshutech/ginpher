/*
@Time: 2020-08-03 10:33
@Auth: xujiang
@File: server_config.go
@Software: GoLand
@Desc: TODO
*/
package config

//Server config
type serverConfig struct{
	AppMode string `json:"AppMode"`
	AppHost string `json:"AppHost" default:"localhost"`
	AppPort int64 `json:"AppHost" default:"9999"`
	AppName string
	AppReadTimeout int64 `json:"AppReadTimeout" default:"30"`
	AppWriteTimeout int64 `json:"AppWriteTimeout" default:"30"`
}

var ServerConfig *serverConfig

//Init server config.
func init() {
	registerConfig(ServerConfig)
}

func (d *serverConfig)getConfigKeyName()(string,interface{}){
	if ServerConfig == nil{
		ServerConfig = new(serverConfig)
	}
	return "Server",ServerConfig
}


