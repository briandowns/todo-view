package command

import (
	"fmt"
	"sort"

	"github.com/fatih/flags"
	"github.com/mitchellh/cli"
)

// Show holds in the passed in configuration
type Show struct {
	todoFormat string
}

// NewShow creates a new instance of Delete
func NewShow() cli.CommandFactory {
	return func() (cli.Command, error) {
		return &Show{
			todoFormat: "TODO(<user>) <message> <weight> <timestamp>",
		}, nil
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
	case "weights":
		s.showWeights()
	case "format":
		s.showFormat()
	default:
		fmt.Println("ERROR: invalid option for show\n")
	}

	return 1
}

// Help provides full help inforamation for the subcommand
func (s *Show) Help() string {
	return `Usage: todo-view show <option> <arguments> 
  Show a resource
Options:
  weights            Display the todo-view weights
  format             Display the todo-view todo format
  
`
}

// Synopsis provides a brief description of the command
func (s *Show) Synopsis() string {
	return "Show a todo-view resource"
}

// showWeights shows all configured weights
func (s *Show) showWeights() {
	fmt.Print("\ntodo-view weights:\n\n")
	w := NewTabWriter()

	var keys []int
	for k := range Weights {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Fprintf(w, "%d\t%s\n", k, Weights[k])
	}

	fmt.Fprintf(w, "\n")
	w.Flush()
}

func (s *Show) showFormat() {
	fmt.Print("\ntodo-view todo format:\n\n")
	w := NewTabWriter()

	fmt.Fprintf(w, s.todoFormat+"\n")

	w.Flush()
}
