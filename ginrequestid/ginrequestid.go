package ginrequestid

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

var RequestID string = "X-Request-Id"

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for incoming header, use it if exists
		requestID := c.Request.Header.Get(RequestID)

		// Create request id with UUID4
		if requestID == "" {
			requestID = GetRequestId()
		}

		// Expose it for use in the application
		c.Set(RequestID, requestID)

		// Set X-Request-Id header
		c.Writer.Header().Set(RequestID, requestID)
		c.Next()
	}
}

func GetRequestId() string {
	uuid4 := uuid.NewV4()
	return uuid4.String()
}
