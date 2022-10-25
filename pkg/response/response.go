package response

import (
	"github.com/gin-gonic/gin"
	"go-api-base/pkg/response/code"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Code:    code.OK,
		Message: "成功",
		Data:    map[string]string{},
	})
}

func Failed(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, Response{
		Code:    code.ErrNot,
		Message: msg,
		Data:    map[string]string{},
	})
	ctx.Abort()
}

func FailedData(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Code:    code.ErrData,
		Message: "请求数据错误",
		Data:    map[string]string{},
	})
}
