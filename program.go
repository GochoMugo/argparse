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
	commands    []Command
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
	command := Command{
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
	command.label = label

	if len(label) > prog.padLength {
		prog.padLength = len(label)
	}

	prog.commands = append(prog.commands, command)
	return prog
}

// Epilog adds an epilog
func (prog *Program) Epilog(epilog string) *Program {
	prog.epilog = epilog
	return prog
}

// Parse does the argument parsing
func (prog *Program) Parse() *Program {
	commandName := os.Args[0]
	var command Command
	var found bool

	for _, v := range prog.commands {
		if v.short == commandName || v.long == commandName {
			command = v
			break
		}
	}

	if found == false {
		println(`INVALID OPTION: {{ commandName }}\nTry \"help\" for a list of available commands`)
		return prog
	}

	args := minimist.ParseArgv(os.Args[1:])
	command.handler(Args(args))

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
		epilog = prog.epilog
	}

	for _, cmd := range prog.commands {
		commands = append(commands, rpad(cmd.label, prog.padLength+4)+cmd.description)
	}

	output := header + "\n" + "Available Commands:\n    " + strings.Join(commands, "\n    ") + "\n\n" + epilog + "\n"
	fmt.Print(output)

	return prog
}

// ShowVersion shows version information to users
func (prog *Program) ShowVersion() *Program {
	fmt.Println("v" + prog.version)
	return prog
}
