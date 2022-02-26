package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var Version string

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func wrapHandlerWithLogging(wrappedHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()

		//写入header信息
		for k, v := range req.Header {
			w.Header().Set(k, strings.Join(v, ","))
		}
		//写入环境变量VERSION
		w.Header().Set("VERSION", Version)

		lrw := NewLoggingResponseWriter(w)
		wrappedHandler.ServeHTTP(lrw, req)

		ip, _, _ := net.SplitHostPort(req.RemoteAddr)
		//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
		log.Printf("request IP:%s, response code:%d, request method:%s, request path:%s, body length: %d, handle time: %s\n", ip,
			lrw.statusCode, req.Method, req.URL.Path, req.ContentLength, time.Since(start))
	})
}

// healthz
func healthz(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func readiness(w http.ResponseWriter, req *http.Request){
	w.WriteHeader(http.StatusOK)
}

func preStop(w http.ResponseWriter, req *http.Request){
	w.WriteHeader(http.StatusOK)
}

func main() {
	Version = os.Getenv("VERSION")
	healthzUri := os.Getenv("healthzUri")
	if len(healthzUri) == 0 {
		healthzUri = "/healthz"
	}
	readinessUri := os.Getenv("readinessUri")
	if len(readinessUri) == 0 {
		readinessUri = "/readiness"
	}
	preStopUri := os.Getenv("preStopUri")
	if len(preStopUri) == 0 {
		preStopUri = "/preStop"
	}
	//定义healthz请求处理
	http.Handle(healthzUri, wrapHandlerWithLogging(http.HandlerFunc(healthz)))
	http.Handle(readinessUri, wrapHandlerWithLogging(http.HandlerFunc(readiness)))
	http.Handle(preStopUri, wrapHandlerWithLogging(http.HandlerFunc(preStop)))

	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	select {
	//监听退出信号
	case s := <-quit:
		log.Println("receive signal :", s)
		time.Sleep(time.Second*2)
		log.Println("exit server")
	}
}
