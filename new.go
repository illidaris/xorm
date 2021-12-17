package xorm

import (
	"go.uber.org/zap"
	xLog "xorm.io/xorm/log"
)

func NewXLogger() xLog.ContextLogger {
	logger = zap.L().WithOptions(zap.AddCallerSkip(6))
	return &XLogger{}
}
