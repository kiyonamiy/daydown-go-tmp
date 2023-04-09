// Copyright 2022 Innkeeper kiyonamiy <yuqingbo0122@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/kiyonamiy/daydown.

package user

import (
	"github.com/kiyonamiy/daydown/internal/daydown/biz"
	"github.com/kiyonamiy/daydown/internal/daydown/store"
	"github.com/kiyonamiy/daydown/pkg/auth"
)

type UserController struct {
	a *auth.Authz
	b biz.IBiz
}

func New(ds store.IStore, a *auth.Authz) *UserController {
	return &UserController{a: a, b: biz.NewBiz(ds)}
}
