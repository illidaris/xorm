package xorm

import (
	"fmt"
	"github.com/illidaris/core"
	"go.uber.org/zap"
	xLog "xorm.io/xorm/log"
)

type XLogger struct{}

// BeforeSQL implements ContextLogger
func (l *XLogger) BeforeSQL(ctx xLog.LogContext) {}

// AfterSQL implements ContextLogger
func (l *XLogger) AfterSQL(ctx xLog.LogContext) {
	// write context meta data
	fields := FieldsFromCtx(ctx.Ctx)
	// write sql data
	fields = append(fields, SQLFromLogContext(ctx)...)
	// write sql param and result
	msg := MessageFromLogContext(ctx)
	// write exec error
	if ctx.Err != nil {
		fields = append(fields, zap.String(core.Error.String(), ctx.Err.Error()))
		logger.Error(msg, fields...)
	} else {
		logger.Info(msg, fields...)
	}
}

// Debug implements ContextLogger
func (l *XLogger) Debug(v ...interface{}) {
	logger.Debug(fmt.Sprint(v...))
}

// Debugf implements ContextLogger
func (l *XLogger) Debugf(format string, v ...interface{}) {
	logger.Debug(fmt.Sprintf(format, v...))
}

// Error implements ContextLogger
func (l *XLogger) Error(v ...interface{}) {
	logger.Error(fmt.Sprint(v...))
}

// Errorf implements ContextLogger
func (l *XLogger) Errorf(format string, v ...interface{}) {
	logger.Error(fmt.Sprintf(format, v...))
}

// Info implements ContextLogger
func (l *XLogger) Info(v ...interface{}) {
	logger.Info(fmt.Sprint(v...))
}

// Infof implements ContextLogger
func (l *XLogger) Infof(format string, v ...interface{}) {
	logger.Info(fmt.Sprintf(format, v...))
}

// Warn implements ContextLogger
func (l *XLogger) Warn(v ...interface{}) {
	logger.Warn(fmt.Sprint(v...))
}

// Warnf implements ContextLogger
func (l *XLogger) Warnf(format string, v ...interface{}) {
	logger.Warn(fmt.Sprintf(format, v...))
}

// Level implements ContextLogger, unavailable
func (l *XLogger) Level() xLog.LogLevel {
	return xLog.LOG_INFO
}

// SetLevel implements ContextLogger, unavailable
func (l *XLogger) SetLevel(lv xLog.LogLevel) {}

// ShowSQL implements ContextLogger
func (l *XLogger) ShowSQL(show ...bool) {

}

// IsShowSQL implements ContextLogger
func (l *XLogger) IsShowSQL() bool {
	return true
}
