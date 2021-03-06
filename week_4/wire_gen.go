// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"geekbang_homework/week_4/conf"
	"geekbang_homework/week_4/internal/server"
	"net/http"
)

// Injectors from wire.go:

func InitializeAllInstance() *http.Server {
	config := conf.InitConfig()
	httpServer := server.InitServer(config)
	return httpServer
}
