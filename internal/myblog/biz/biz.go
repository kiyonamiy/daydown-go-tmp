// Copyright 2022 Innkeeper kiyonamiy <yuqingbo0122@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/kiyonamiy/myblog.

package biz

import (
	"github.com/kiyonamiy/myblog/internal/myblog/biz/user"
	"github.com/kiyonamiy/myblog/internal/myblog/store"
)

type IBiz interface {
	Users() user.UserBiz
}

type biz struct {
	ds store.IStore
}

// Users 返回一个实现了 UserBiz 接口的实例。
func (b *biz) Users() user.UserBiz {
	return user.New(b.ds)
}

// NewBiz 创建一个 IBiz 类型的实例。
func NewBiz(ds store.IStore) *biz {
	return &biz{ds: ds}
}

var _ IBiz = (*biz)(nil)
