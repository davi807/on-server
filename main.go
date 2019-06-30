package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"on-server/fs"
)

func main() {
	fs.FormatSize(1024)
	return
	res := ""
	dir, _ := ioutil.ReadDir(".")
	for _, d := range dir {
		res += fmt.Sprintln(d.IsDir(), d.Name())
	}

	http.HandleFunc("/fs/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		w.Write([]byte(res))
	})

	http.ListenAndServe("0.0.0.0:1200", nil)
}
