/*
Package argparse keeps parsing command-line arguments simple!
*/
package argparse

import (
	"fmt"
	"os"
	"strings"

	"github.com/mgutz/minimist"
)

// Program is the terminal interface object
type Program struct {
	name        string
	version     string
	description string
	commands    []command
	epilog      string
	padLength   int
}

// New returns a new Program parser
func New() *Program {
	prog := Program{}
	prog.Command("H", "help", "show help information", func(a Args) { prog.ShowHelp() })
	prog.Command("V", "version", "show version information", func(a Args) { prog.ShowVersion() })
	return &prog
}

// Description adds a name and description of the program
func (prog *Program) Description(name, description string) *Program {
	prog.name = name
	prog.description = description
	return prog
}

// Version adds version information
func (prog *Program) Version(version string) *Program {
	prog.version = version
	return prog
}

// Command adds a new command
func (prog *Program) Command(short, long, description string, handler func(Args)) *Program {
	cmd := command{
		short:       short,
		long:        long,
		handler:     handler,
		description: description,
	}
	var label string
	if short != "" {
		label += short + ", "
	}
	label += long
	cmd.label = label

	if len(label) > prog.padLength {
		prog.padLength = len(label)
	}

	prog.commands = append(prog.commands, cmd)
	return prog
}

// Epilog adds an epilog
func (prog *Program) Epilog(epilog string) *Program {
	prog.epilog = epilog
	return prog
}

// Parse does the argument parsing
func (prog *Program) Parse() *Program {
	if len(os.Args) == 1 {
		prog.ShowHelp()
		return prog
	}

	var commandName = os.Args[1]
	var cmd command
	var found = false

	for _, v := range prog.commands {
		if v.short == commandName || v.long == commandName {
			cmd = v
			found = true
			break
		}
	}

	if found == false {
		fmt.Println("INVALID OPTION: " + commandName)
		fmt.Println("Try \"help\" for a list of available commands")
		return prog
	}

	args := minimist.ParseArgv(os.Args[2:])
	cmd.handler(Args(args))

	return prog
}

// ShowHelp shows help information to users
func (prog *Program) ShowHelp() *Program {
	var header string
	var commands []string
	var epilog string

	if prog.name != "" {
		header += prog.name + ": "
	}

	if prog.description != "" {
		header += prog.description + "\n"
	}

	if prog.epilog != "" {
		epilog = "\n\n" + prog.epilog
	}

	for _, cmd := range prog.commands {
		commands = append(commands, rpad(cmd.label, prog.padLength+4)+cmd.description)
	}

	output := header + "\n" + "Available Commands:\n    " + strings.Join(commands, "\n    ") + epilog + "\n"
	fmt.Print(output)

	return prog
}

// ShowVersion shows version information to users
func (prog *Program) ShowVersion() *Program {
	fmt.Println("v" + prog.version)
	return prog
}
