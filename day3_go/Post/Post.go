package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type person struct {
	User    string `json:"user"`
	Comment string `json:"comment"`
}

func simulatePostCommentRequest() {
	comment := person{
		User:    "小明",
		Comment: "这本书很好看！",
	}

	jsonData, err := json.Marshal(comment)
	if err != nil {
		log.Printf("JSON编码失败: %v", err)
		return
	}

	resp, err := http.Post("http://localhost:8080/comment", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("POST请求失败: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("读取响应失败: %v", err)
		return
	}

	fmt.Printf("POST /comment 响应:\n%s\n", string(body))
}

func commenthandler(w http.ResponseWriter, r *http.Request) {
	var p person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "解析请求体失败: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"Message": "评论已提交",
		"User":    p.User,
		"Comment": p.Comment,
	})
}

func main() {
	http.HandleFunc("/comment", commenthandler)
	fmt.Println("服务器正在运行，访问地址: http://localhost:8080/comment")

	go func() {
		time.Sleep(500 * time.Millisecond)
		simulatePostCommentRequest()
	}()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务器启动失败:", err)
	}
}
