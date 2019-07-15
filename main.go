package main

import (
	"io"
	"net/http"
	"on-server/fs"
	"os"
	"strconv"
	"time"
)

func main() {

	var err error

	err = initFlags()

	if err != nil {
		println(err.Error())
		return
	}

	if flags.ip == "" {
		go println(startText(ipList(flags.showIP6), flags.port))
	} else {
		println(startText([]string{flags.ip}, flags.port))
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log(r)
		switch r.Method {
		case "GET":
			w.Header().Add("Content-Type", "text/html; charset=utf8")
			w.Write([]byte(top + makeBody(!flags.noMessage, !flags.noUpload, !flags.noFiles) + bottom))
		case "POST":
			var limit = flags.limit * (2 << 20)

			if r.ContentLength > limit {
				w.WriteHeader(400)
				w.Write([]byte("Data is lagrer then maximum size limit " + strconv.FormatInt(flags.limit, 10) + " Mb"))
				return
			}

			r.Body = http.MaxBytesReader(w, r.Body, limit)
			defer r.Body.Close()

			parts, _ := r.MultipartReader()
			for {
				part, err := parts.NextPart()
				if err != nil {
					break
				}

				if !flags.noMessage && part.FormName() == "message" {
					data := make([]byte, 1024)
					n, _ := part.Read(data)
					if n == 0 {
						continue
					}
					f, _ := os.Create(flags.messagePath + "/" + "msg-" + time.Now().Format("2-1-2006 21:04:05") + ".txt")
					f.Write([]byte("Message source ip: " + r.RemoteAddr + "\n"))
					f.Write(data[0:n])
					io.Copy(f, part)
					f.Close()
				} else if !flags.noUpload && part.FormName() == "file" && len(part.FileName()) > 0 {
					f, _ := os.Create(flags.uploadPath + "/" + part.FileName())
					io.Copy(f, part)
					f.Close()
				}
			}
			http.Redirect(w, r, "/", 301)
		}
	})

	if !flags.noFiles {
		fs.SetFilesRoot(flags.path)
		http.Handle(fs.URLRoot, fs.Handler{})
	}

	http.ListenAndServe(flags.ip+":"+flags.port, nil)
}

func log(r *http.Request) {
	sep := "   "
	method := r.Method
	if l := len(method); l < 4 {
		method += " "
	}
	println(method + sep + time.Now().Format("2-1-2006 21:04:05") + sep + r.URL.RequestURI())
}
