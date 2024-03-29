// Copyright 2022 Innkeeper kiyonamiy <yuqingbo0122@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/kiyonamiy/daydown.

package model

import "time"

type PostM struct {
	ID        int64     `gorm:"column:id;primary_key"` //
	Username  string    `gorm:"column:username"`       //
	PostID    string    `gorm:"column:postID"`         //
	Title     string    `gorm:"column:title"`          //
	Content   string    `gorm:"column:content"`        //
	CreatedAt time.Time `gorm:"column:createdAt"`      //
	UpdatedAt time.Time `gorm:"column:updatedAt"`      //
}

// TableName sets the insert table name for this struct type
func (p *PostM) TableName() string {
	return "post"
}
