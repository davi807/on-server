package main

import (
	"io"
	"net/http"
	"on-server/fs"
	"os"
	"strconv"
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
		switch r.Method {
		case "GET":
			w.Header().Add("Content-Type", "text/html; charset=utf8")
			w.Write([]byte(top + makeBody(!flags.noMessage, !flags.noUpload, !flags.noFiles) + bottom))
		case "POST":
			//todo upload
			var limit int64 = flags.limit * (2 << 20)

			if r.ContentLength > limit {

				w.WriteHeader(400)
				w.Write([]byte("Maximum data size is lagre then limit " + strconv.FormatInt(flags.limit, 10) + " Mb"))
			}

			r.Body = http.MaxBytesReader(w, r.Body, limit)
			defer r.Body.Close()

			parts, _ := r.MultipartReader()
			for {
				part, err := parts.NextPart()
				if err != nil {
					break
				}

				//todo
				if len(part.FileName()) > 0 {
					f, _ := os.Create(flags.uploadPath + "/" + part.FileName())
					io.Copy(f, part)
					f.Close()
				}
			}
		}
		/*


				for {

					part, err := parts.NextPart()

					if err != nil {
						break
					}

					if part.FormName() != "" {
						fmt.Println("___++----", part.FormName())
					}

					b := make([]byte, 1024)
					n, e := part.Read(b)

					if n == 0 {
						println("nothing", e)
						continue
					}
					fmt.Println("__", string(b))
				}

				return
			}
		*/
	})

	if !flags.noFiles {
		fs.SetFilesRoot(flags.path)
		http.Handle(fs.URLRoot, fs.Handler{})
	}

	http.ListenAndServe(flags.ip+":"+flags.port, nil)
}
