// Copyright 2022 Innkeeper kiyonamiy <yuqingbo0122@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/kiyonamiy/daydown.

package daydown

import (
	"github.com/gin-gonic/gin"
	"github.com/kiyonamiy/daydown/internal/daydown/controller/v1/user"
	"github.com/kiyonamiy/daydown/internal/daydown/store"
	"github.com/kiyonamiy/daydown/internal/pkg/core"
	"github.com/kiyonamiy/daydown/internal/pkg/errno"
	"github.com/kiyonamiy/daydown/internal/pkg/log"
	"github.com/kiyonamiy/daydown/internal/pkg/middleware"
	"github.com/kiyonamiy/daydown/pkg/auth"
)

func installRouters(g *gin.Engine) error {
	g.NoRoute(func(ctx *gin.Context) {
		core.WriteResponse(ctx, errno.ErrPageNotFound, nil)
	})

	g.GET("/healthz", func(ctx *gin.Context) {
		log.C(ctx).Infow("Healthz function called")

		core.WriteResponse(ctx, nil, map[string]string{"status": "ok"})
	})

	authz, err := auth.NewAuthz(store.S.DB())
	if err != nil {
		return err
	}

	uc := user.New(store.S, authz)

	g.POST("/login", uc.Login)

	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	{
		// 创建 users 路由分组
		userv1 := v1.Group("/users")
		{
			userv1.POST("", uc.Create)
			userv1.PUT(":name/change-password", uc.ChangePassword)
			userv1.Use(middleware.Authn(), middleware.Authz(authz))
			userv1.GET(":name", uc.Get) // 获取用户详情
		}
	}

	return nil
}
