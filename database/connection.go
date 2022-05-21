package database

import (
	"go.uber.org/zap"
	"github.com/savaukr/restApiGinGorm/model"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"time"
	"fmt"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	prdLogger, _ := zap.NewProduction()
	defer prdLogger.Sync()
	logger := prdLogger.Sugar()

	dns := "host=localhost user=sl password=1111 dbname=messages port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		logger.Fatalw("failed to connetc to DB", "err", err)
	}
	db.AutoMigrate(&model.Users{}, &model.Messages{})
	logger.Infow("connected to DB ...")
	return db
}

func GetDB() *gorm.DB {
	if dbase == nil {
		dbase = Init()
		sleep := time.Duration(1)
		for dbase == nil {
			sleep := sleep * 2
			fmt.Printf("database is unavailable. Wait for %d sec.\n", sleep)
			time.Sleep(sleep * time.Second)
			dbase = Init()
		}
 	}
	return dbase
}