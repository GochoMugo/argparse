package argparse

import (
	"bytes"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func runCmd(t *testing.T, args ...string) string {
	var out bytes.Buffer
	cmd := exec.Command("go", append([]string{"run", "samples/app.go"}, args...)...)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		t.Error("errored: ", err)
	}
	return out.String()
}

func TestCommandFunctions(t *testing.T) {
	out := runCmd(t)
	assert.Contains(t, out, "show help information")

	out = runCmd(t, "H")
	assert.Contains(t, out, "show help information")

	out = runCmd(t, "V")
	assert.Contains(t, out, "v0.0.0")

	out = runCmd(t, "s")
	assert.Contains(t, out, "application started")

	out = runCmd(t, "start")
	assert.Contains(t, out, "application started")

	out = runCmd(t, "x")
	assert.Contains(t, out, "application stopped")

	out = runCmd(t, "f", "--bool", "--key=value", "--key2", "value2")
	assert.Contains(t, out, "bool true")
	assert.Contains(t, out, "key value")
	assert.Contains(t, out, "key2 value2")
}
