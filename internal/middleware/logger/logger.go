package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func LoggerMiddlerware() gin.HandlerFunc {
	logPath := "log/log"
	scr, _ := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE, 0755)

	logger := logrus.New()
	logger.Out = scr
	logger.SetLevel(logrus.DebugLevel)

	logWriter, _ := rotatelogs.New(
		logPath+"%Y%m%d.log",
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(Hook)

	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		costTime := time.Since(startTime).Milliseconds()
		cost := fmt.Sprintf("%d ms", costTime)

		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}

		statusCode := ctx.Writer.Status()
		clientIp := ctx.ClientIP()
		userAgent := ctx.Request.UserAgent()
		dataSize := ctx.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := ctx.Request.Method
		path := ctx.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"status":    statusCode,
			"SpendTime": cost,
			"Ip":        clientIp,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})

		if len(ctx.Errors) > 0 {
			entry.Error(ctx.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
