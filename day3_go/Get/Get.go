package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	fmt.Fprintf(w, "您正在查询图书：《%s》", title)
}

func simulateGetBookRequest() {
	url := "http://localhost:8080/book?title=三体"

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("GET请求失败: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("读取响应失败: %v", err)
		return
	}

	fmt.Printf("GET /book 响应:\n%s\n\n", string(body))
}

func main() {
	http.HandleFunc("/book", handler)
	fmt.Println("服务器正在运行，访问地址：http://localhost:8080/book?title=三体")

	go func() {
		time.Sleep(500 * time.Millisecond)
		simulateGetBookRequest()
	}()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务器启动失败:", err)
	}
}
