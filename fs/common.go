package fs

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// URLRoot is static file server handler location
var URLRoot = "/files/"

// SetFilesRoot location of files list
func SetFilesRoot(addr string) {
	filesRoot = addr
	if addr[len(addr)-1] == '/' {
		filesRoot = addr[0 : len(addr)-2]
	}
}

var filesRoot string

func formatSize(size int64) string {
	if size < 1024 {
		return strconv.FormatInt(size, 10) + "B"
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

	row := `<tr class="file-row">`

	row += `<td class="file-type">`
	if f.IsDir() {
		row += `🗀`
	} else {
		row += `.📄`
	}
	row += `</td>`

	row += `<td class="file-name">` +
		`<a href="` + root + "/" + f.Name() + `"`
	if !f.IsDir() {
		row += ` target="_blank"`
	}
	row += `>` + f.Name() + `</a>` +
		`</td>`

	row += `<td class="file-size">` + formatSize(f.Size()) + `</td>`

	row += `<td class="file-modtime">` + f.ModTime().Format("2-1-2006 15:04") + `</td>`

	row += `</tr>`
	return row
}

func upLink(u string) string {
	if u[len(u)-1] == '/' {
		u = u[0 : len(u)-2]
	}
	return u[0:strings.LastIndex(u, "/")]
}

func makeList(root string, list []os.FileInfo) string {
	var path = strings.Replace(root, URLRoot, "/", 1)

	var result = "<h1>Index of " + path + "</h1>"

	result += `<tr>
		<th></th>
		<th>Name</th>
		<th>Size</th>
		<th>Last Modified</th>
	</tr>`

	if path == "/" {
		if root == "/" {
			root = ""
		}
	} else {
		result += `<tr><td></td>` +
			`<td><a href="` + upLink(root) + `">../</a></td>` +
			`<td></td><td></td></tr>`
	}

	for _, item := range list {
		result += makeFileTemplate(root, item)
	}

	return result
}
