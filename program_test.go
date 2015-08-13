package argparse

import "testing"

var name = "example"
var description = "description"
var version = "0.0.0"
var epilog = "epilog"

func ExampleProgram() {
	prog := New()
	prog.Description(name, description)
	prog.Version(version)
	prog.Epilog(epilog)
	prog.Command("s", "start", "starts application", func(a Args) {})

	prog.ShowHelp()
	// Output:
	// example: description
	//
	// Available Commands:
	//     H, help       show help information
	//     V, version    show version information
	//     s, start      starts application
	//
	// epilog
}

func TestVersion(t *testing.T) {
	prog := New()
	prog.Version("0.1.1")
	prog.ShowVersion()
	// Output: v0.1.1
}
