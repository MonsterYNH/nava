package engine

import (
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/MonsterYNH/nava/setting"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	requestID       = "X-Request-Id"
	ResponseData    = "response_data"
	ResponseErrCode = "response_err_code"
)

func RecoverMiddleware(config setting.Config, logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()
		defer func() {
			duration := time.Since(now)

			logData := map[string]interface{}{
				"request_id":  c.GetHeader(requestID),
				"duration":    duration,
				"method":      c.Request.Method,
				"uri":         c.Request.RequestURI,
				"status_code": c.Writer.Status(),
			}

			if err := recover(); err != nil {
				// get stack info
				stackInfos := debug.Stack()
				// init logger info TODO get panic file name, line and func name
				logData["stack_info"] = string(stackInfos)

				log.Println(logData)
				logger.WithFields(logData).Error(err)
				c.JSON(http.StatusOK, NewApiResponse(nil, ErrUnknown))
				return
			}
			logger.WithFields(logData).Info("")

			response, _ := c.Get(ResponseData)
			switch resp := response.(type) {
			case *APIResponse:
				c.JSON(c.Writer.Status(), resp)
			default:
				c.JSON(c.Writer.Status(), NewApiResponse(nil, ErrUnknown))
			}
		}()

		c.Next()
	}
}
