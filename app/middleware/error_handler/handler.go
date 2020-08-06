/*
@Time: 2020-08-04 10:27
@Auth: xujiang
@File: handler.go
@Software: GoLand
@Desc: error handler middleware for gin
*/
package error_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
}

func NewErrorResponse(statusCode, code int, msg string) *ErrorResponse {
	return &ErrorResponse{StatusCode: statusCode, Code: code, Msg: msg}
}

var (
	Success     = NewErrorResponse(http.StatusOK, 0, "success")
	ServerError = NewErrorResponse(http.StatusInternalServerError, 200500, "系统异常，请稍后重试!")
	NotFound    = NewErrorResponse(http.StatusNotFound, 200404, http.StatusText(http.StatusNotFound))
)

func OtherError(message string) *ErrorResponse {
	return NewErrorResponse(http.StatusForbidden, 100403, message)
}

func (e *ErrorResponse) Error() string {
	return e.Msg
}

func HandleNotFound(c *gin.Context) {
	err := NotFound
	c.JSON(err.StatusCode, err)
	return
}

func ErrHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var Err *ErrorResponse
				if e, ok := err.(*ErrorResponse); ok {
					Err = e
				} else if e, ok := err.(error); ok {
					Err = OtherError(e.Error())
				} else {
					Err = ServerError
				}
				// 记录一个错误的日志
				// TODO:  send alarm
				c.JSON(Err.StatusCode, Err)
				return
			}
		}()
		c.Next()
	}
}

func Register(engine *gin.Engine) {
	engine.NoRoute(HandleNotFound)
	engine.NoMethod(HandleNotFound)
	engine.Use(ErrHandler())
}
