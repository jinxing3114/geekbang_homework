package server

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// ListenSignal 监听 signal 信号，如果是退出，停止服务，等待1秒后退出
func ListenSignal(srv *http.Server) {
	//监听退出信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	s := <-quit

	//设置超时context

	logrus.Infoln("get a signal, Shutdown Server: ", s.String())
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	//结束http服务，等待处理完成当前请求，拒绝新请求
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatalln("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		logrus.Infoln("timeout of 1 seconds.")
	}
	logrus.Infoln("Server exiting")
}

