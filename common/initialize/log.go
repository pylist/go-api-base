package initialize

import (
	"go-api-base/common"
	"go.uber.org/zap"
)

func Log() {
	common.Logger, _ = zap.NewProduction()
}
