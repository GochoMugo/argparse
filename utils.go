package argparse

import (
	"strings"
)

func rpad(s string, l int) string {
	return s + strings.Repeat(" ", l-len(s))
}
