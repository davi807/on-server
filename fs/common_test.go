package fs

import (
	"testing"
)

func TestFormatSize(t *testing.T) {

	if FormatSize(2048) != "2.0K" {
		t.Error("wrong number for 2048 -> 2.0K")
	}

	if FormatSize(3000) != "2.9K" {
		t.Error("wrong number for 3000 -> 2.9K")
	}

	if FormatSize(566528000) != "540.3M" {
		t.Error("wrong number for 566528000 -> 540.2M")
	}

	if FormatSize(3758096384) != "3.5G" {
		t.Error("wrong number for 3758096384 -> 3.5G")
	}

}
