package web

import (
	"encoding/json"
	"ginpher/app/dao"
	"ginpher/app/model"
	"ginpher/app/tools/api"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/vgmdj/utils/logger"
)

// 获取验证码
type CheckVerifyCodeReq struct {
	Mobile string //手机号
}

// 登录
// 获取用户信息
func Personal(c *gin.Context) {
	paramByte, err := formatContext(c)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusOK, api.NewResponse(api.BadRequest, ""))
		return
	}
	var param CheckVerifyCodeReq
	if err := json.Unmarshal(paramByte, &param); err != nil {
		c.JSON(http.StatusOK, api.NewResponse(api.BadRequest, ""))
		return
	}
	if param.Mobile == "" {
		c.JSON(http.StatusOK, api.NewResponse(api.LackParams, ""))
		return
	}
	// 是否注册
	user, code := checkRegisterRealName(param.Mobile, 0, false)
	if code != api.Success {
		c.JSON(http.StatusOK, api.NewResponse(code, ""))
		return
	}

	c.JSON(http.StatusOK, api.NewResponse(api.Success, user))
	return
}

func checkRegisterRealName(mobile string, userId int64, IsCheckRealName bool) (user model.User, code api.Code) {
	code = api.Forbidden
	if mobile == "" && userId == 0 {
		return
	}
	var err error
	// 是否注册
	if mobile != "" {
		user, err = dao.UserGetInfoByMobile(mobile)
		if err != nil && err.Error() == "record not found" {
			code = api.NoRegister
			return
		} else if err != nil {
			return
		}
	} else if userId != 0 {
		user, err = dao.UserGetInfoByUserId(userId)
		if err != nil && err.Error() == "record not found" {
			code = api.NoRegister
			return
		} else if err != nil {
			return
		}
	}
	if IsCheckRealName {
		// 是否实名
		if user.IdCard == "" || user.IdName == "" {
			code = api.NoRealName
			return
		}
	}

	code = api.Success

	return
}

//不加密时 获取客户端数据
func formatContext(c *gin.Context) (pridecrypt []byte, err error) {
	pridecrypt, err = ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error(err)
	}

	return
}
