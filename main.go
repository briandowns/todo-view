package main

import (
	"fmt"
	"os"

	"github.com/briandowns/todo-view/command"
	"github.com/mitchellh/cli"
)

const (
	todoViewVersion = "0.2"
	todoViewName    = "todo-view"
)

func main() {
	if retval, err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(retval)
	}
}

func run() (int, error) {
	c := &cli.CLI{
		Name:     todoViewName,
		Version:  todoViewVersion,
		Args:     os.Args[1:],
		HelpFunc: cli.BasicHelpFunc(todoViewName),
		Commands: map[string]cli.CommandFactory{
			"parse":   command.NewParse(),
			"show":    command.NewShow(),
			"export":  command.NewExport(),
			"web":     command.NewWeb(),
			"version": command.NewVersion(todoViewVersion),
		},
	}

	retval, err := c.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1, err
	}

	return retval, nil
}
