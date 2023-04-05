// Copyright 2022 Innkeeper kiyonamiy <yuqingbo0122@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/kiyonamiy/daydown.

package user

import (
	"context"
	"regexp"

	"github.com/jinzhu/copier"
	"github.com/kiyonamiy/daydown/internal/daydown/store"
	v1 "github.com/kiyonamiy/daydown/internal/pkg/api/daydown/v1"
	"github.com/kiyonamiy/daydown/internal/pkg/errno"
	"github.com/kiyonamiy/daydown/internal/pkg/model"
)

type UserBiz interface {
	Create(ctx context.Context, r *v1.CreateUserRequest) error
}

type userBiz struct {
	store store.IStore
}

var _ UserBiz = (*userBiz)(nil)

// New 创建一个实现了 UserBiz 接口的实例。
func New(store store.IStore) *userBiz {
	return &userBiz{store: store}
}

func (biz *userBiz) Create(ctx context.Context, r *v1.CreateUserRequest) error {
	var userM model.UserM
	_ = copier.Copy(&userM, r)

	if err := biz.store.Users().Create(ctx, &userM); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'username'", err.Error()); match {
			return errno.ErrUserAlreadyExist
		}

		return err
	}

	return nil
}
