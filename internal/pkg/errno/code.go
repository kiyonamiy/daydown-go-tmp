// Copyright 2022 Innkeeper kiyonamiy <yuqingbo0122@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/kiyonamiy/myblog.

package errno

var (
	OK = &Errno{HTTP: 200, Code: "", Message: ""}

	InternalServerError = &Errno{HTTP: 500, Code: "InternalError", Message: "Internal server error."}

	ErrPageNotFound = &Errno{HTTP: 404, Code: "ResourceNotFound.PageNotFound", Message: "Page not found."}
)
