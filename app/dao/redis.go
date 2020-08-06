package dao

import (
	"encoding/json"

	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/vgmdj/utils/logger"
)

// RedisSet set
func RedisSet(key string, value interface{}) (err error) {
	key = makeRedisKey(key)
	v, err := json.Marshal(value)
	if err != nil {
		logger.Info(key, value)
		return err
	}
	err = NewDao().redis.Set(key, v)
	if err != nil {
		return
	}
	return
}

// RedisGetBytes get []byte
func RedisGetBytes(key string) (bytes []byte, err error) {
	key = makeRedisKey(key)
	bytes, err = NewDao().redis.GetBytes(key)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	return
}

// RedisGetString get string
func RedisGetString(key string) (str string, err error) {
	key = makeRedisKey(key)
	bts, err := NewDao().redis.GetBytes(key)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	err = json.Unmarshal(bts, &str)
	if err != nil {
		logger.Info(string(bts))
		logger.Error(err.Error())
		return
	}
	return
}
func RedisGetInt(key string) (num int, err error) {
	key = makeRedisKey(key)
	bts, err := NewDao().redis.GetBytes(key)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	err = json.Unmarshal(bts, &num)
	if err != nil {
		logger.Info(string(bts))
		logger.Error(err.Error())
		return
	}
	return
}

// Gets the remaining cache time
func RedisKeyTime(key string) (l int64, err error) {
	key = makeRedisKey(key)
	ttl, err := NewDao().redis.Do("TTL", key)
	if err == nil {
		l = ttl.(int64)
	}
	return
}

// Expire
func Expire(key string, second int) error {
	key = makeRedisKey(key)

	return NewDao().redis.Expire(key, second)
}

func AddRedisHash(key string, value interface{}) (err error) {
	key = makeRedisKey(key)

	v, err := json.Marshal(value)
	if err != nil {
		logger.Info(key, value)
		return err
	}
	_, err = NewDao().redis.Do("HMSET", "taofu_tel", key, v)

	return
}

func GetRedisHash(key string) ([]byte, error) {
	key = makeRedisKey(key)
	result, err := redis.Values(NewDao().redis.Do("HMGET", "taofu_tel", key)) //读
	if err != nil || result[0] == nil {
		return nil, err
	}

	for _, v := range result {
		return v.([]byte), nil
	}

	return nil, nil
}

func RedisDelete(key string) (err error) {
	key = makeRedisKey(key)
	err = NewDao().redis.DEL(key)
	return
}

func makeRedisTableKey(table, key string) (k string) {
	k = fmt.Sprintf("%s:%s", table, key)
	return
}

func makeRedisKey(key string) (k string) {
	k = fmt.Sprintf("tf-5g:%s", key)
	return
}
