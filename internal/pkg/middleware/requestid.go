// Copyright 2022 Innkeeper kiyonamiy <yuqingbo0122@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/kiyonamiy/daydown.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kiyonamiy/daydown/internal/pkg/known"
)

func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 检查请求头中是否有 `X-Request-ID`，如果有则复用，没有则新建
		request := ctx.Request.Header.Get(known.XRequestIDKey)
		if request == "" {
			request = uuid.New().String()
		}

		// 将 RequestID 保存在 gin.Context 中，方便后边程序使用
		ctx.Set(known.XRequestIDKey, request)

		ctx.Writer.Header().Set(known.XRequestIDKey, request)
		ctx.Next()
	}
}
