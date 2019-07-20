package main

import (
	"net/http"
	"on-server/fs"
	"time"
)

func main() {
	err := initFlags()
	if err != nil {
		println(err.Error())
		return
	}

	if flags.ip == "" {
		go println(startText(ipList(flags.showIP6), flags.port))
	} else {
		println(startText([]string{flags.ip}, flags.port))
	}

	if !(flags.noMessage && flags.noUpload) {

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			log(r)
			w.Header().Add("Content-Type", "text/html; charset=utf8")
			w.Write([]byte(top + makeBody(!flags.noMessage, !flags.noUpload, !flags.noFiles) + bottom))
		})

		http.HandleFunc("/upload", uploadHandle)
	} else {
		fs.URLRoot = "/"
	}

	if !flags.noFiles {
		fs.SetFilesRoot(flags.path)
		http.HandleFunc(fs.URLRoot, func(w http.ResponseWriter, r *http.Request) {
			log(r)
			fs.Handle(w, r)
		})
	}

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	http.ListenAndServe(flags.ip+":"+flags.port, nil)
}

/*** ***/

func log(r *http.Request) {
	sep := "   "
	method := r.Method
	if l := len(method); l < 4 {
		method += " "
	}
	println(method + sep + time.Now().Format("2-1-2006 21:04:05") + sep + r.URL.RequestURI())
}
