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
		User:    "С��",
		Comment: "�Ȿ��ܺÿ���",
	}

	jsonData, err := json.Marshal(comment)
	if err != nil {
		log.Printf("JSON����ʧ��: %v", err)
		return
	}

	resp, err := http.Post("http://localhost:8080/comment", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("POST����ʧ��: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("��ȡ��Ӧʧ��: %v", err)
		return
	}

	fmt.Printf("POST /comment ��Ӧ:\n%s\n", string(body))
}

func commenthandler(w http.ResponseWriter, r *http.Request) {
	var p person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "����������ʧ��: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"Message": "�������ύ",
		"User":    p.User,
		"Comment": p.Comment,
	})
}

func main() {
	http.HandleFunc("/comment", commenthandler)
	fmt.Println("�������������У����ʵ�ַ: http://localhost:8080/comment")

	go func() {
		time.Sleep(500 * time.Millisecond)
		simulatePostCommentRequest()
	}()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("����������ʧ��:", err)
	}
}
