package fs

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Handle show file list
func Handle(w http.ResponseWriter, r *http.Request, enableIndex bool) {
	fileName := strings.Replace(r.RequestURI, URLRoot, "", 1)
	filePath, err := url.QueryUnescape(filesRoot + "/" + fileName)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`Server error`))
		return
	}
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(`"` + fileName + `"` + ` not found`))
		return
	}

	if fileInfo.IsDir() {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")

		files, err := ioutil.ReadDir(filePath)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(`Server error`))
			return
		}

		var page string

		page = top +
			makeList(URLRoot+fileName, files) +
			bottom

		w.Write([]byte(page))
	} else {
		if enableIndex {
			pathChunks := strings.Split(r.RequestURI, "/")
			if pathChunks[len(pathChunks)-1] == "index.html" {
				indexContent, _ := ioutil.ReadFile(filePath)
				w.Write(indexContent)
				return
			}
		}
		http.ServeFile(w, r, filePath)
	}

}
