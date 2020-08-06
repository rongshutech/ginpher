package api

import (
	"ginpher/app/tools/api"
	"net/http"
	"time"

	myjwt "ginpher/app/middleware/jwt_handler"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 生成令牌
func GenerateToken(c *gin.Context) {
	j := &myjwt.JWT{
		[]byte("ginpher"),
	}
	claims := myjwt.CustomClaims{
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "ginpher",                       // 签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, api.NewResponse(api.QueryFailed, ""))
		return
	}

	c.JSON(http.StatusOK, api.NewResponse(api.Success, token))
	return
}
