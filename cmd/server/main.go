package main

import (
	"fmt"
	"github.com/savaukr/restApiGinGorm/pkg/config"
	"go.uber.org/zap"
	// "gorm.io/gorm"
	// "gorm.io/driver/postgres"
	"github.com/savaukr/restApiGinGorm/database"


	"github.com/gin-gonic/gin"

)

func main() {
	prdLogger, _ := zap.NewProduction()
	defer prdLogger.Sync()
	logger := prdLogger.Sugar()

	logger.Infow("Hello!")

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatalw("failed to parse config", "err", err)
	}
	fmt.Printf("cfg = %+v\n", cfg)


	// dns := "host=localhost user=sl password=1111 dbname=messages port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// // db, err := gorm.Open(postgres.Open(cfg.DBAddr), &gorm.Config{})
	// db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	// if err != nil {
	// 	logger.Fatalw("failed to connetc to DB", "err", err)
	// }

	db:= database.Init()
	fmt.Printf("db = %+v\n", db)


	//routers
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	
}