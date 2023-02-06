// Copyright 2022 Innkeeper kiyonamiy <yuqingbo0122@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/kiyonamiy/myblog.

package myblog

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/kiyonamiy/myblog/internal/pkg/log"
	"github.com/kjzz/viper"
	"github.com/spf13/cobra"
)

const (
	// recommendedHomeDir 定义放置 myblog 服务配置的默认目录.
	recommendedHomeDir = ".myblog"

	// defaultConfigName 指定了 myblog 服务的默认配置文件名.
	defaultConfigName = "myblog.yaml"
)

func initConfig() {
	if cfgFile != "" {
		// 从命令行选项指定的配置文件中读取
		viper.SetConfigFile(cfgFile)
	} else {
		// 查找用户主目录
		home, err := os.UserHomeDir()
		// 如果获取用户主目录失败，打印 `'Error: xxx` 错误，并退出程序（退出码为 1）
		cobra.CheckErr(err)

		// 将用 `$HOME/<recommendedHomeDir>` 目录加入到配置文件的搜索路径中
		viper.AddConfigPath(filepath.Join(home, recommendedHomeDir))

		// 把当前目录加入到配置文件的搜索路径中
		viper.AddConfigPath(".")

		// 设置配置文件格式为 YAML(YAML格式清晰易读，并且支持复杂的配置结构)
		viper.SetConfigType("yaml")

		// 配置文件名称（没有文件扩展名）
		viper.SetConfigName(defaultConfigName)
	}

	// 读取匹配的环境变量
	viper.AutomaticEnv()

	// 读取环境变量的前缀为 MYBLOG，如果是 myblog，将自动转变为大写。
	viper.SetEnvPrefix("MYBLOG")

	// 以下 2 行，将 viper.Get(key) key 字符串中 '.' 和 '-' 替换为 '_'
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 读取配置文件。如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
	err := viper.ReadInConfig()
	if err != nil {
		log.Errorw("Failed to read viper configuration file", "err", err)
	} else {
		// 打印 viper 当前使用的配置文件，方便 Debug。
		log.Infow("Using config file", "file", viper.ConfigFileUsed())
	}
}

// logOptions 从 viper 中读取日志配置，构建 `*log.Options` 并返回。
func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		Level:             viper.GetString("log.level"),
		Format:            viper.GetString("log.format"),
		OutputPaths:       viper.GetStringSlice("log.output-paths"),
	}
}
