package handlers 

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping (ctx * gin.Context) {
	ctx.JSON(200, gin.H{
		"response" :  gin.H {
			"method" : http.MethodGet,
			"code": http.StatusOK,
			"message": "pong",
		},
	})		
}
func GetMessage (ctx *gin.Context) {

}

func SendMessage(ctx *gin.Context) {

}
