//+build wireinject

package main

import (
	"geekbang_homework/week_4/conf"
	"geekbang_homework/week_4/internal/dao"
	"geekbang_homework/week_4/internal/server"
	"github.com/google/wire"
	"net/http"
)

func InitializeAllInstance() *http.Server {
	////初始化配置信息，使用配置文件
	//conf.InitConfig()
	////初始化redis

	wire.Build(conf.InitConfig, dao.InitRedis, server.InitServer)
	return &http.Server{}
}

