package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// 绑定
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthHandler)

	// 监听
	err := http.ListenAndServe(":80", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	// write response header
	for k, v := range r.Header {
		w.Header().Add(k, v[0])
	}

	w.Header().Add("version", os.Getenv("VERSION"))

	// visit log
	log := fmt.Sprintf("%s ：客户端IP：%s，HTTP返回码：%d", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, http.StatusOK)
	fmt.Println(log)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200")
}
