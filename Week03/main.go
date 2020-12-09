package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
)

const HttpPort = 8087

func startServer() error {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Go!")
	})

	log.Printf("server start: %d\n", HttpPort)
	return http.ListenAndServe(fmt.Sprintf(":%d", HttpPort), nil)
}

// 监听系统signal
func Signal(ctx context.Context) error {
	// 创建信号chan
	quit := make(chan os.Signal, 1)

	// 监听signal
	signal.Notify(quit)
	defer signal.Stop(quit)

	for {
		select {
		case <- ctx.Done():
			return fmt.Errorf("http server exited")
		case s := <-quit:
			return fmt.Errorf("get signal: %v\n", s)
		}
	}
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	// 启动http server
	g.Go(func() error {
		return startServer()
	})

	g.Go(func() error {
		return Signal(ctx)
	})

	//g.Go(func() error {
	//	return errors.New("test")
	//})

	g.Go(func() error {
		<-ctx.Done()
		return ctx.Err()
	})

	// 等待错误
	if err := g.Wait(); err != nil {
		fmt.Printf("service exited %s\n", err.Error())
	}
	//err := g.Wait()
	//fmt.Println(err)

	//fmt.Println(ctx.Err())
}
