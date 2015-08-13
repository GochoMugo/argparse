package argparse

import (
	"github.com/mgutz/minimist"
)

// command is a [sub-]command following the program name
type command struct {
	short       string
	long        string
	handler     func(Args)
	label       string
	description string
}

// Args is a struct passed to all command functions. It has the
// ArgMap from minimist embedded.
type Args struct {
	minimist.ArgMap
}
