package main

import (
	"net/http"
	"on-server/fs"
)

func main() {

	http.Handle(fs.URLRoot, fs.Handler{})

	// return
	// res := ""
	// dir, _ := ioutil.ReadDir(".")
	// for _, d := range dir {
	// 	res += fmt.Sprintln(d.IsDir(), d.Name())
	// }

	// http.HandleFunc("/fs/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println(r.URL)
	// 	w.Write([]byte(res))
	// })

	http.ListenAndServe("0.0.0.0:1200", nil)
}
