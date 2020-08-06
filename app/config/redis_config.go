package config

type redisConfig struct {
	Address string `json:"Address" default:"localhost:6379"`
	PassWd  string `json:"Address" default:"root"`
	DB      int    `json:"DB"`
}

var RedisConfig *redisConfig

func init() {
	registerConfig(RedisConfig)
}

func (d *redisConfig) getConfigKeyName() (string, interface{}) {
	if RedisConfig == nil {
		RedisConfig = new(redisConfig)
	}
	return "Redis", DbConfig
}
