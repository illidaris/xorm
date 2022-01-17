package xorm

import (
	"context"
	"fmt"

	"github.com/illidaris/core"
	"go.uber.org/zap"
	xLog "xorm.io/xorm/log"
)

var (
	logger *zap.Logger // log core without context
)

// FieldsFromCtx write context meta data
func FieldsFromCtx(ctx context.Context) []zap.Field {
	return []zap.Field{
		buildZapField(ctx, core.TraceID),
		buildZapField(ctx, core.SessionID),
		buildZapField(ctx, core.Action),
		buildZapField(ctx, core.Step),
	}
}

// SQLFromLogContext write sql data
func SQLFromLogContext(ctx xLog.LogContext) []zap.Field {
	// write sql cost
	return []zap.Field{
		zap.String(core.Category.String(), "SQL"),
		zap.Int64(core.Duration.String(), ctx.ExecuteTime.Milliseconds()),
	}
}

// MessageFromLogContext build message
func MessageFromLogContext(ctx xLog.LogContext) string {
	var queryString string
	if len(ctx.SQL) > 2048 {
		queryString = ctx.SQL[0:2048]
	} else {
		queryString = ctx.SQL
	}
	args := make([]interface{}, 0)
	if len(ctx.Args) > 100 {
		args = ctx.Args[0:100]
	}
	var rows int64
	var rowErr error
	if ctx.Result != nil {
		rows, rowErr = ctx.Result.RowsAffected()
	}
	return fmt.Sprintf("[sql]%s,[args]%v,[rows]%d,[err]%s", queryString, args, rows, rowErr)
}

// buildZapField use core meta build field
func buildZapField(ctx context.Context, key core.MetaData) zap.Field {
	return zap.String(key.String(), key.GetString(ctx))
}
