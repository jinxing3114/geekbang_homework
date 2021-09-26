package server

import (
	"geekbang_homework/week_4/app/account"
	"geekbang_homework/week_4/conf"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitServer(config *conf.Config) *http.Server {
	router := gin.New()
	// 主页方法
	router.GET("/", account.GetInfo)

	srv := &http.Server{
		Addr:         config.Server.Addr,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return srv
}
