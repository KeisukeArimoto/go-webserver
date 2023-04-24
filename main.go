package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// 受け取ったリクエストに対するレスポンスを返すためのハンドラ関数
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}

	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	}

	// ハンドラとURLの紐づけ
	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)

	// サーバの起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}
