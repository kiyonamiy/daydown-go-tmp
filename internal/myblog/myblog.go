// Copyright 2022 Innkeeper kiyonamiy <yuqingbo0122@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/kiyonamiy/myblog.

package myblog

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kiyonamiy/myblog/internal/pkg/log"
	"github.com/kiyonamiy/myblog/pkg/version/verflag"
	"github.com/kjzz/viper"
	"github.com/spf13/cobra"
)

var cfgFile string

func NewMyBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "myblog",
		Short: "My little blog",
		Long: `My little blog.
		
Find more myblog information at:
https://github.com/kiyonamiy/myblog`,
		// 命令出错时，不打印帮助信息。不需要打印帮助信息，设置为 true 可以保持命令出错时一眼就能看到错误信息
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时，执行的 Run 函数
		RunE: func(cmd *cobra.Command, args []string) error {
			// 如果 `--version=true`，则打印版本并退出
			verflag.PrintAndExitIfRequested()

			log.Init(logOptions())
			defer log.Sync() // Sync 将缓存中的日志刷新到磁盘文件中
			return run()
		},
		// 这里设置命令运行时，不需要指定命令行参数（例如执行 `_output/myblog test`，会抛错；执行 `_output/myblog -x` 会报“Error: unknown shorthand flag: 'x' in -x”，说明 tag 不属于 args）
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	// 以下设置，使得 initConfig 函数在每个命令运行时都会被调用以读取配置
	cobra.OnInitialize(initConfig)

	// Cobra 支持持久性标志(PersistentFlag)，该标志可用于它所分配的命令以及该命令下的每个子命令
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the myblog configuration file. Empty string for no configuration file.")

	// Cobra 也支持本地标志，本地标志只能在其所绑定的命令上使用
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// 添加 --version 标志
	verflag.AddFlags(cmd.PersistentFlags())

	return cmd
}

// run 函数是实际的业务代码入口函数.
func run() error {
	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()

	g.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"code": 10003, "message": "Page not found."})
	})

	g.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	httpsrv := &http.Server{Addr: viper.GetString("addr"), Handler: g}

	// 运行 HTTP 服务器
	// 打印一条日志，用来提示 HTTP 服务已经起来，方便排障
	log.Infow("Start to listening the incoming requests on http address", "addr", viper.GetString("addr"))
	if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalw(err.Error())
	}

	return nil
}
