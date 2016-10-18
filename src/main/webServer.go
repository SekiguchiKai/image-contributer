package main

import (
"fmt"
"net/http"
)

func init() {
	// "/hello"パターンとsayHelloをバンドル
	http.HandleFunc("/", sayHello)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, GAE/GO!")
}

func main() {
	// 8080ポートをリッスンするようにする=>クライアントからの8080への接続を待ち受けている
	fmt.Println("server run at http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
