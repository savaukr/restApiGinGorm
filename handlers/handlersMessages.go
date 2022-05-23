package handlers 

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/savaukr/restApiGinGorm/database"
	"github.com/savaukr/restApiGinGorm/models"
	
)

var db = database.GetDB()


func GetMessage (ctx *gin.Context) {
	var message models.Messages
	id := ctx.Param("id")

	if err := db.Find(&message, id).Error; err != nil {
		ctx.AbortWithStatus(404)
	} else {
    ctx.JSON(200, message)
	}

}
func GetMessages (ctx *gin.Context) {
	var messages []models.Messages
	if err := db.Find(&messages).Error; err != nil {
		ctx.AbortWithStatus(404)
	} else {
    ctx.JSON(200, messages)
	}
}

func SaveMessage(ctx *gin.Context)  {
	var message models.Messages
	if err := ctx.BindJSON(&message); err != nil {
		return
	}
	db.Create(&message)
	ctx.JSON(http.StatusOK, gin.H{
		"method" : http.MethodPost,
		"code": http.StatusOK,
		"messageId": message.Id,
	})
	
}

func DeleteMessage(ctx *gin.Context) {
	var message models.Messages
	id := ctx.Param("id")
	if err := db.Delete(&message, id).Error; err != nil {
		ctx.AbortWithStatus(404)
	} else {
    ctx.JSON(200, &message)
	}	
}

func UpdateMessage(ctx *gin.Context) {
	var message models.Messages
	id := ctx.Param("id")
	if err := ctx.BindJSON(&message); err != nil {
		return
	}
	if err := db.Model(&message).Where("Id = ?", id).Update("message", message.Message ).Error; err != nil {
		ctx.AbortWithStatus(404)
	} else {
    ctx.JSON(200, &message)
	}	
}



