package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	down := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		return webServer(down)
	})

	g.Go(func() error {
		return signalServer(down)
	})

	if err := g.Wait(); err != nil {
		fmt.Println("g.Wait err:", err)
	}
}

func webServer(down chan struct{}) error {
	server := http.Server{
		Addr:    ":8081",
		Handler: nil,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	http.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		close(down)
		fmt.Fprintf(w, "shutdown")
	})

	done := make(chan error)
	go func() {
		<-down
		fmt.Println("shutdown")
		ctx, cancle := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancle()

		done <- server.Shutdown(ctx)
	}()

	err := server.ListenAndServe()
	if err != nil {
		return <-done
	}

	return nil
}

func signalServer(down chan struct{}) error {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-down:
		close(quit)
	case <-quit:
		close(down)
	}

	signal.Reset(syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("signal return")
	return nil
}
