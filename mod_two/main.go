package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	VERSION = "VERSION"
)

// 1、接收客户端 request，并将 request 中带的 header 写入 response header
func newRequest1() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		logAccess(request)
		for k, v := range request.Header {
			writer.Header().Set(k, strings.Join(v, ";"))
		}
	})
}

// 2、读取当前系统的环境变量中的 VERSION 配置，并写入 response header
func newRequest2() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		logAccess(request)
		writer.Header().Set(VERSION, os.Getenv(VERSION))
	})
}

// 3、Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
func logAccess(request *http.Request) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), request.RemoteAddr, "200")
}

// 4、当访问 localhost/healthz 时，应返回 200
func healthz() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		logAccess(request)
		writer.Write([]byte("200"))
	})
}

func GracefulExit(server *http.Server) {
	fmt.Println("GracefulExit...")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("server.Shutdown, ", err)
	}
	fmt.Println("shutdown ok ")
}

func main() {
	fmt.Println("server start!")
	mux := http.NewServeMux()

	mux.Handle("/request1", newRequest1())
	mux.Handle("/request2", newRequest2())
	mux.Handle("/healthz", healthz())
	//_ = http.ListenAndServe(":80", mux)
	server := &http.Server{Addr: ":80", Handler: mux}

	// 监听指定信号：sigterm
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGTERM:
				fmt.Println("sigterm")
				GracefulExit(server)
			default:
				fmt.Println("other signal", s)
			}
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("exit reason, ", err)
	}

	fmt.Println("main goroutine exit")
}
