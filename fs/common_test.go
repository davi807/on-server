package fs

import (
	"testing"
)

func TestFormatSize(t *testing.T) {

	if formatSize(2048) != "2.0K" {
		t.Error("wrong number for 2048 -> 2.0K")
	}

	if formatSize(3000) != "2.9K" {
		t.Error("wrong number for 3000 -> 2.9K")
	}

	if formatSize(566528000) != "540.3M" {
		t.Error("wrong number for 566528000 -> 540.2M")
	}

	if formatSize(3758096384) != "3.5G" {
		t.Error("wrong number for 3758096384 -> 3.5G")
	}

}
