package logger

import (
	"context"

	log "github.com/sirupsen/logrus"
)

func Info(ctx context.Context, args ...interface{}) {
	log.WithContext(ctx).Info(args)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	log.WithContext(ctx).Infof(format, args)
}

func Error(ctx context.Context, args ...interface{}) {
	log.WithContext(ctx).Error(args)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	log.WithContext(ctx).Errorf(format, args)
}
