package main

import (
	"fmt"
	"github.com/savaukr/restApiGinGorm/pkg/config"
	"go.uber.org/zap"

	"github.com/savaukr/restApiGinGorm/database"
	"github.com/savaukr/restApiGinGorm/handlers"
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

	db:= database.Init()
	fmt.Printf("db = %+v\n", db)

	//routers
	r := gin.Default()

	r.GET("/message/:id", handlers.GetMessage)
	r.GET("/messages", handlers.GetMessages)
	r.POST("/message", handlers.SaveMessage)
	r.DELETE("./message/:id", handlers.DeleteMessage)
	r.PUT("./message/:id", handlers.UpdateMessage)

	r.GET("/users/:id", handlers.GetUser)
	r.GET("/users", handlers.GetUsers)
	r.POST("/users", handlers.SaveUser)
	r.DELETE("./users/:id", handlers.DeleteUser)
	r.PUT("./users/:id", handlers.UpdateUser)
	
	r.Run(cfg.HttpAddr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	
}