package main

import (
	"flag"
	"log"

	cmdOpcodes "github.com/triztian/gowss.rt/cmd/opcodes"
)

var flags struct {
	Help bool
}

func init() {
	flag.BoolVar(&flags.Help, "help", false, "Print this help")

	flag.Parse()
}

func main() {
	args := flag.Args()

	if len(args) <= 0 {
		log.Fatal("must supply a command")
	}
	command := args[0]
	switch command {
	case cmdOpcodes.Name:
		if err := cmdOpcodes.Run(args...); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("unknown command: ", command)
	}
}

const helpText = `
A WebAssembly runtime written in Go.

Usage:
	gowss.rt [flags] opcodes [files...]

Options:
	--help, -h		Print this help
`
