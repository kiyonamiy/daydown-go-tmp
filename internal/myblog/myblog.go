// Copyright 2022 Innkeeper kiyonamiy <yuqingbo0122@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/marmotedu/miniblog.

package myblog

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewMyBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "myblog",
		Short: "My little blog",
		Long: `My little blog.
		
Find more miniblog information at:
https://github.com/kiyonamiy/myblog`,
		// 命令出错时，不打印帮助信息。不需要打印帮助信息，设置为 true 可以保持命令出错时一眼就能看到错误信息
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时，执行的 Run 函数
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
		// 这里设置命令运行时，不需要指定命令行参数（例如执行 `_output/myblog test`，会抛错）
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	return cmd
}

// run 函数是实际的业务代码入口函数.
func run() error {
	fmt.Println("Hello MyBlog!")
	return nil
}
