// Copyright 2022 Innkeeper kiyonamiy <yuqingbo0122@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/kiyonamiy/daydown.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kiyonamiy/daydown/internal/pkg/core"
	"github.com/kiyonamiy/daydown/internal/pkg/errno"
	"github.com/kiyonamiy/daydown/internal/pkg/known"
	"github.com/kiyonamiy/daydown/pkg/token"
)

func Authn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解析 JWT Token
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
