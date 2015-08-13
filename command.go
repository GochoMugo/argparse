package argparse

import (
	"github.com/mgutz/minimist"
)

// Command is a [sub-]command following the program name
type Command struct {
	short       string
	long        string
	handler     func(Args)
	label       string
	description string
}

// Args is a map passed to command functions
type Args minimist.ArgMap
