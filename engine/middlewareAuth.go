package engine

import (
	"net/http"

	"github.com/MonsterYNH/nava/setting"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func AuthMiddleware(config setting.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestId string
		// If request id exists, use it to pass it to the next server. If it does not exist, generate one
		if len(c.GetHeader(requestID)) == 0 {
			requestId = uuid.NewV4().String()
		} else {
			requestId = c.GetHeader(requestID)
		}

		// set RequestId
		c.Set(requestID, requestId)
		c.Header(requestID, requestId)

		if !config.Server.EnableAuthCheck {
			c.Next()
			return
		}

		if !checkAuth() {
			c.JSON(http.StatusOK, NewApiResponse(nil, ErrNoAuth))
			c.Abort()
			return
		}

		c.Next()
	}
}

func checkAuth() bool {
	return false
}
