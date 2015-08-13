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

// Args is a map passed to command functions. This is simply an alias to
// ArgMap from github.com/mgutz/minimist.
type Args minimist.ArgMap
