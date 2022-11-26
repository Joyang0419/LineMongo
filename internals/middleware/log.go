package middleware

import (
	"LineMongo"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"time"
)

func JSONLogMiddleware(isSaveFile bool, setLevel log.Level) gin.HandlerFunc {
	logger := Logger(isSaveFile, setLevel)
	return func(c *gin.Context) {
		startTime := time.Now()
		// Process Request
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)

		entry := logger.WithFields(log.Fields{
			"client_ip":  c.ClientIP(),
			"duration":   latencyTime,
			"method":     c.Request.Method,
			"path":       c.Request.RequestURI,
			"status":     c.Writer.Status(),
			"referrer":   c.Request.Referer(),
			"request_id": c.Writer.Header().Get("Request-Id"),
		})

		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("")
		}
	}
}

func Logger(isSaveFile bool, setLevel log.Level) *log.Logger {
	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{})
	logger.SetLevel(setLevel)

	if !isSaveFile {
		return logger
	}

	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}

	src := &lumberjack.Logger{
		Filename:   path.Join(logFilePath, now.Format(LineMongo.DateStringFormat)+".log"),
		MaxSize:    1,  // megabytes after which new file is created
		MaxBackups: 3,  // number of backups
		MaxAge:     10, //days
	}
	writer := io.MultiWriter(src, os.Stdout)
	logger.SetOutput(writer)
	return logger
}
