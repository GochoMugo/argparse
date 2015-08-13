package main

import (
	"fmt"

	"github.com/GochoMugo/argparse"
)

func main() {
	prog := argparse.New()
	prog.Description("app", "sample app").Version("0.0.0")
	prog.Epilog("this is the end!")
	prog.Command("s", "start", "start application", func(a argparse.Args) {
		fmt.Println("application started")
	}).Command("x", "stop", "stop application", func(a argparse.Args) {
		fmt.Println("application stopped")
	}).Command("f", "flags", "process flags", func(a argparse.Args) {
		for key, value := range a.ArgMap {
			fmt.Println(key, value)
		}
	})
	prog.Parse()
}
