package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"strings"
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

func main() {
	Version = os.Getenv("VERSION")
	//定义healthz请求处理
	http.Handle("/healthz", wrapHandlerWithLogging(http.HandlerFunc(healthz)))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
