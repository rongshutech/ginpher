package router

import (
	"ginpher/app/api"
	"ginpher/app/api/web"
	"ginpher/app/middleware/jwt_handler"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {

	r.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// 获取令牌
	r.POST("/generatetoken", api.GenerateToken)

	jwt_handler.Register(r)
	/*
		用户
	*/
	//{
	r.POST("/personal", web.Personal) // 获取用户信息
	//}
}
