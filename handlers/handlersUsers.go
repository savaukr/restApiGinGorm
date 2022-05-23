package handlers 

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/savaukr/restApiGinGorm/models"
)



func GetUser (ctx *gin.Context) {
	var user models.Users
	id := ctx.Param("id")

	if err := db.Find(&user, id).Error; err != nil {
		ctx.AbortWithStatus(404)
	} else {
    ctx.JSON(200, user)
	}

}
func GetUsers (ctx *gin.Context) {
	var users []models.Users
	if err := db.Find(&users).Error; err != nil {
		ctx.AbortWithStatus(404)
	} else {
    ctx.JSON(200, users)
	}
}

func SaveUser(ctx *gin.Context)  {
	var user models.Users
	if err := ctx.BindJSON(&user); err != nil {
		return
	}
	db.Create(&user)
	ctx.JSON(http.StatusOK, gin.H{
		"method" : http.MethodPost,
		"code": http.StatusOK,
		"messageId": user.Id,
	})
	
}

func DeleteUser(ctx *gin.Context) {
	var user models.Users
	id := ctx.Param("id")
	if err := db.Delete(&user, id).Error; err != nil {
		ctx.AbortWithStatus(404)
	} else {
    ctx.JSON(200, &user)
	}	
}

func UpdateUser(ctx *gin.Context) {
	var user models.Users
	id := ctx.Param("id")
	if err := ctx.BindJSON(&user); err != nil {
		return
	}
	if err := db.Model(&user).Where("Id = ?", id).Update("UserName", user.UserName ).Error; err != nil {
		ctx.AbortWithStatus(404)
	} else {
    ctx.JSON(200, &user)
	}	
}
