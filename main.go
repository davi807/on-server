package main

import (
	"fmt"
	"net/http"
	"on-server/fs"
)

func main() {

	fmt.Println(listURLIPs(ipList(true), "1200"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; charset=utf8")
		w.Write([]byte(top + makeBody(true, true, true) + bottom))
	})

	http.Handle(fs.URLRoot, fs.Handler{})

	http.ListenAndServe(":1200", nil)
}
