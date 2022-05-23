package handlers 

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// "github.com/savaukr/restApiGinGorm/database"
)

// var db = database.GetDB()

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
	// result:=db.
	ctx.JSON(200, gin.H{
		"response" :  gin.H {
			"method" : http.MethodGet,
			"code": http.StatusOK,
			"message": "pong",
		},
	})	

}
func GetMessages (ctx *gin.Context) {

}

func SaveMessage(ctx *gin.Context) {
	// id := ctx.Query("name")
	

}
