package tlog

import (
	"go.uber.org/zap"
)

var S *Logger

type Logger struct {
	*zap.SugaredLogger
}

func ReplaceS(l *Logger) {
	S = l
}