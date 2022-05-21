package handlers 

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping (ctx * gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})		
}
func GetMessage (ctx *gin.Context) {

}
