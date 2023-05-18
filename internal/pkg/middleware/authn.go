// // Copyright 2023 Innkeeper Belm(郁凯) <yukai98@foxmain.com>. All rights reserved.
// // Use of this source code is governed by a MIT style
// // license that can be found in the LICENSE file. The original repo for
// // this file is https://github.com/ProgramKai/miniblog

package middleware

import (
	"cn.xdmnb/study/miniblog/internal/pkg/errno"
	"cn.xdmnb/study/miniblog/internal/pkg/known"
	"cn.xdmnb/study/miniblog/internal/pkg/response"
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
			response.WriteResponse(c, errno.ErrTokenInvalid, nil)
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
