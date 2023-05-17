package middleware

import (
	"cn.xdmnb/study/miniblog/internal/pkg/known"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestID 是一个 Gin 中间件，用来在每一个 HTTP 请求的 context, response 中注入 `X-Request-ID` 键值对.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get(known.XRequestIDKey)
		if requestID == "" {
			requestID = uuid.New().String()
		}

		c.Set(known.XRequestIDKey, requestID)
		c.Writer.Header().Set(known.XRequestIDKey, requestID)
		c.Next()
	}
}
