// Copyright 2022 Innkeeper kiyonamiy <yuqingbo0122@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/kiyonamiy/daydown.

package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/kiyonamiy/daydown/internal/pkg/core"
	"github.com/kiyonamiy/daydown/internal/pkg/errno"
	"github.com/kiyonamiy/daydown/internal/pkg/log"
	v1 "github.com/kiyonamiy/daydown/pkg/api/daydown/v1"
)

const defaultMethods = "(GET)|(POST)|(PUT)|(DELETE)"

func (ctrl *UserController) Create(ctx *gin.Context) {
	// 在 Controller 层实现有限的功能（参数解析、校验、逻辑分发、请求聚合和返回）
	log.C(ctx).Infow("Create user function called")

	var r v1.CreateUserRequest
	if err := ctx.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(ctx, errno.ErrBind, nil)
		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(ctx, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)
		return
	}

	if err := ctrl.b.Users().Create(ctx, &r); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	// 添加授权信息
	if _, err := ctrl.a.AddNamedPolicy("p", r.Username, "/v1/users/"+r.Username, defaultMethods); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, nil)
}
