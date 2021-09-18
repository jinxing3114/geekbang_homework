package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, time.Now().String())
	})
	server := &http.Server{
		Addr:         "127.0.0.1:8080",
		Handler:      mux,
		WriteTimeout: time.Second * 3,
	}

	eg, ctx := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})
	eg.Go(func() error {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
		select {
		case <-quit:
			return server.Shutdown(ctx)
		case <-ctx.Done():
			return nil
		}
	})
	err := eg.Wait()
	log.Println(err)
}


