package middleware

import (
	"cn.xdmnb/study/miniblog/internal/pkg/core"
	"cn.xdmnb/study/miniblog/internal/pkg/errno"
	"cn.xdmnb/study/miniblog/internal/pkg/known"
	"cn.xdmnb/study/miniblog/internal/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"regexp"
)

func Authn() gin.HandlerFunc {
	return func(c *gin.Context) {
		support := authnSupport(c)
		if support {
			c.Next()
			return
		}
		username, err := token.ParseRequest(c)
		if err != nil {
			core.WriteResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Set(known.XUsernameKey, username)
		c.Next()
	}
}

func authnSupport(c *gin.Context) bool {
	ignoreAuthnUri := viper.GetStringSlice("ignore-authn-uri")
	for _, uri := range ignoreAuthnUri {
		matched, _ := regexp.MatchString(uri, c.Request.URL.Path)
		if matched {
			return true
		}
	}
	return false
}
