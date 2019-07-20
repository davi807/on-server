package main

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func uploadHandle(w http.ResponseWriter, r *http.Request) {
	log(r)
	if r.Method == "POST" {

		if r.ContentLength > flags.limit {
			w.WriteHeader(400)
			w.Write([]byte("Data is lagrer then maximum size limit " + strconv.FormatInt(flags.limit, 10) + " Mb"))
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, flags.limit)
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

				f, err := os.Create(flags.messagePath + "/" + "msg-" + time.Now().Format("2-1-2006 15:04:05") + ".txt")
				if err != nil {
					setError(w, 500, "Nessage not saved :(")
					return
				}

				f.Write([]byte("Message source address: " + r.RemoteAddr + "\n"))
				f.Write(data[0:n])
				io.Copy(f, part)
				f.Close()
			} else if !flags.noUpload && part.FormName() == "file" && len(part.FileName()) > 0 {

				if _, err := os.Stat(flags.uploadPath + "/" + part.FileName()); err == nil {
					setError(w, 400, "File '"+part.FileName()+"' already exists, try to change file name.")
					return
				}

				f, err := os.Create(flags.uploadPath + "/" + part.FileName())

				if err != nil {
					setError(w, 500, "Can not upload file :(")
					return
				}

				io.Copy(f, part)
				f.Close()
			}
		}
	}

	http.Redirect(w, r, "/", 301)
}

func setError(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	w.Write([]byte(msg))
}
