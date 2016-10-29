package command

import (
	"fmt"
	"sort"

	"github.com/fatih/flags"
	"github.com/mitchellh/cli"
)

// Show represents the show command
type Show struct{}

// NewShow creates a new instance of Delete
func NewShow() cli.CommandFactory {
	return func() (cli.Command, error) {
		return &Show{}, nil
	}
}

// Run shows a given resource.
func (s *Show) Run(args []string) int {
	if flags.Has("help", args) || len(args) < 1 {
		fmt.Print(s.Help())
		return 1
	}

	// process the subcommand and it's options
	switch args[0] {
	case "format":
		s.showFormat()
	case "priorities":
		s.showPriorities()
	default:
		fmt.Print("ERROR: invalid option for show\n\n")
	}

	return 1
}

// Help provides full help inforamation for the subcommand
func (s *Show) Help() string {
	return `Usage: todo-view show <option> <arguments> 
  Show a resource
  
Options:
  format             Display todo-view format
  priorities         Display todo-view priorities
  
`
}

// Synopsis provides a brief description of the command
func (s *Show) Synopsis() string {
	return "Show a todo-view resource"
}

// showFormat outputs the format to be used for the TODO's
func (s *Show) showFormat() {
	fmt.Print("\ntodo-view todo format:\n\n")
	w := NewTabWriter()

	fmt.Fprintf(w, todoFormat+"\n\n")
	fmt.Fprintf(w, "Example:\n\n")
	fmt.Fprintln(w, "TODO(briandowns) this is an example todo format 2016-05-13T18:54 4")

	w.Flush()
}

// showPriorities shows all configured priorities
func (s *Show) showPriorities() {
	fmt.Print("\ntodo-view priorities:\n\n")
	w := NewTabWriter()

	var keys []int
	for k := range Priorities {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Fprintf(w, "%d\t%s\n", k, Priorities[k])
	}

	fmt.Fprintf(w, "\n")
	w.Flush()
}
