package argparse

import (
	"strings"
	"testing"
)

func TestRpad(t *testing.T) {
	sample := "sample"
	padded := rpad(sample, len(sample)+10)
	if len(padded) != len(sample)+10 {
		t.Error("did not pad to correct length")
	}
	if padded[len(sample):] != strings.Repeat(" ", 10) {
		t.Error("did not add all spaces")
	}
}
