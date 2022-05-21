package main

import (
	"fmt"
	"github.com/savaukr/restApiGinGorm/pkg/config"
	"uber-go.org/zap"
)

func main() {
	prdLogger, _ := zap.NewProduction()
	defer prdLogger.Sync()
	logger := prdLogger.Sugar()

	cfg, err != config.NewConfig()
	if err nil {
		logger.Fatalw("failed to parse config", "err", err)
	}
	fmt.Printf("cfg = %+v\n", cfg)
}