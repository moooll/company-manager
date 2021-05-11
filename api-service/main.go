package main

import (
	"fmt"

	"github.com/moooll/company-manager/api-service/cmd"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("can't initialize zap logger: %v", err)
		return
	}

	defer logger.Sync()
	cmd.Serve()
}
