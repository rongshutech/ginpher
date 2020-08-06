package dao

import (
	"encoding/json"
	"ginpher/app/model"
	"strconv"

	"github.com/vgmdj/utils/logger"

	"ginpher/app/global"
)

// 获取用户基本信息(by mobile)
func UserGetInfoByMobile(mobile string) (user model.User, err error) {

	//mysql
	user, err = UserDBGetInfoByMobile(mobile)
	if err != nil {
		return
	}

	return
}

// 获取用户基本信息(by mobile)
func UserDBGetInfoByMobile(mobile string) (user model.User, err error) {
	err = NewDao().db.Table(global.TableUsers).Where("mobile = ? and status != ? ", mobile, global.Cancellation).Find(&user).Error

	return
}

// 获取用户基本信息(by user_id)
func UserGetInfoByUserId(userId int64) (user model.User, err error) {
	//redis
	if bytes, err := RedisGetBytes(makeRedisTableKey(global.TableUsers, strconv.Itoa(int(userId)))); err == nil {
		if err := json.Unmarshal(bytes, &user); err == nil {
			logger.Info("--redis--user--")
			return user, nil
		}
	} else {
		logger.Error("Query user with id from redis occured error,error is .", err.Error())
	}

	//mysql
	user, err = userDBGetInfoByUserId(userId)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	////save redis
	//if err = RedisSet(makeRedisTableKey(global.TableUsers, strconv.Itoa(int(user.Id))), user); err == nil {
	//	Expire(makeRedisTableKey("users", strconv.Itoa(int(user.Id))), global.WebApiSettings.RedisDurationDefault)
	//} else {
	//	logger.Error("save user with id to redis occured error,error is .", err.Error())
	//}

	return
}

// 获取用户基本信息(by user_id)
func userDBGetInfoByUserId(userId int64) (user model.User, err error) {
	err = NewDao().db.Table(global.TableUsers).Where("id = ? and status != ? ", userId, global.Cancellation).Find(&user).Error

	return
}
