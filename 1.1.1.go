package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

//“1.1.1　テストエコーサーバーの実行”

func handler(w http.ResponseWriter, r *http.Request) { //クライアントからアクセスがあったときに呼ばれる。
	dump, err := httputil.DumpRequest(r, true) //HTTP /1.x形式で返す。
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler) //"/"でアクセスしたらhandlerを実行する。
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
