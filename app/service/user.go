package service

import (
	"fmt"
	"ginpher/app/dao"
	"ginpher/app/model"
)

// 获取用户信息
func GetUserInfo(mobile string) (user model.User, err error) {
	fmt.Println()
	user, err = dao.UserGetInfoByMobile(mobile)

	return
}
