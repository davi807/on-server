package main

import (
	"net/http"
	"on-server/fs"
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
		}
		/*		switch r.Method == "POST" {

					parts, _ := r.MultipartReader()

					for {

						part, err := parts.NextPart()

						if err != nil {
							print("errrrrrrrrr")
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
