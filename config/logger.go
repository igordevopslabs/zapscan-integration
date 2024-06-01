package config

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger

	LOG_OUTPUT = "LOG_OUTPUT"
	LOG_LEVEL  = "LOG_LEVEL"
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{getOutputLogs()},            //personalização da saida de logs
		Level:       zap.NewAtomicLevelAt(getLevelLogs()), //personalização do level dos logs ex: debug, info, etc
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	logger, _ = logConfig.Build()
}

func LogInfo(message string, tags ...zap.Field) {
	logger.Info(message, tags...)
	logger.Sync()
}

func LogError(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	logger.Info(message, tags...)
	logger.Sync()
}

//função que pega o output através da variavel de ambiente.
//função garante que o valor é uma string e minuscula através dos metodos de `strings`

func getOutputLogs() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT)))
	if output == "" {
		return "stdout"
	}

	return output
}

func getLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL))) {
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}
