package fs

import (
	"fmt"
	"os"
	"strconv"
)

// URLRoot is static file server handler location
var URLRoot = "/files-list/"

// SetFilesRoot location of files list
func SetFilesRoot(addr string) {
	filesRoot = addr
	if addr[len(addr)-1] == '/' {
		filesRoot = addr[0 : len(addr)-2]
	}
}

var filesRoot string

func init() {
	var err error
	filesRoot, err = os.Getwd()

	if err != nil {
		fmt.Printf("!! Can not initialize files root to current workig directory")
	}

}

func formatSize(size int64) string {
	if size < 1024 {
		return strconv.FormatInt(size, 10)
	}

	var finalSize = float64(size)
	var metrics = [4]string{1: "K", "M", "G"}
	var result string

	for i := 1; i < 4; i++ {
		finalSize /= 1024
		if finalSize < 1024 || i == 3 {
			result = fmt.Sprintf("%.1f%s", finalSize, metrics[i])
			break
		}
	}

	return result
}

func makeFileTemplate(root string, f os.FileInfo) string {

	row := `<div class="file-row">`

	row += `<span class="file-type">`
	if f.IsDir() {
		row += `🗀.`
	} else {
		row += `📄`
	}
	row += `</span>`

	row += `<span class="file-name">` +
		`<a href="` + URLRoot + f.Name() + `"`
	if !f.IsDir() {
		row += ` target="_blank"`
	}
	row += `>` + f.Name() + `</a>` +
		`</span>`

	row += `<span class="file-size">` + formatSize(f.Size()) + `</span>`

	row += `<span class="file-modtime">` + f.ModTime().Format("2-1-2006 15:04") + `</span>`

	row += `</div>`
	return row
}

func makeList(root string, list []os.FileInfo) string {
	var result string

	for _, item := range list {
		result += makeFileTemplate(root, item)
	}

	return result
}
