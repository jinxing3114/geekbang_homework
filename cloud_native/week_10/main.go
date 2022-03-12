package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// healthz
func healthz(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// hello
func hello(w http.ResponseWriter, req *http.Request) {
	start:=time.Now()
	defer ObserveTotal(start)

	rand.Seed(time.Now().UnixNano())
	delay := rand.Int31n(2000)
	time.Sleep(time.Millisecond*time.Duration(delay))

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "===================Details of the http request header:============\n")
	for k, v := range req.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
	glog.Infof("access delay:%d millisecond", delay)
}

func readiness(w http.ResponseWriter, req *http.Request){
	w.WriteHeader(http.StatusOK)
}

func preStop(w http.ResponseWriter, req *http.Request){
	w.WriteHeader(http.StatusOK)
}

func main() {
	flag.Parse()
	//flag.Set("v", "4")
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
	mux := http.NewServeMux()
	mux.HandleFunc("/hello",hello)
	mux.HandleFunc(healthzUri,healthz)
	mux.HandleFunc(readinessUri,readiness)
	mux.HandleFunc(preStopUri,preStop)
	mux.Handle("/metrics", promhttp.Handler())

	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed{
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	select {
	//监听退出信号
	case s := <-quit:
		log.Println("receive signal :", s)

	}
	ctx, can := context.WithTimeout(context.Background(), time.Second * 2)
	defer can()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Println("exit server")
}
