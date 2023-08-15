package ginlog

import (
	"time"

	"github.com/qmaru/qlog/base"

	"github.com/gin-gonic/gin"
)

func Logger(logfile string) (gin.HandlerFunc, error) {
	log := base.NewLog()
	log.SetOutput(logfile)

	logger, err := log.New()
	if err != nil {
		return nil, err
	}
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqURI := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 日志格式
		logger.Info("Gin",
			"method", reqMethod,
			"status", statusCode,
			"ip", clientIP,
			"uri", reqURI,
			"latency", latencyTime.String(),
		)
	}, nil
}
