package fs

import (
	"fmt"
	"strconv"
)

// FormatSize formats file size value in bytes to string format in Kb, Mb, Gb
func FormatSize(size int) string {
	if size < 1024 {
		return strconv.FormatInt(int64(size), 10)
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
