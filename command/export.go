package command

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/flags"
	"github.com/mitchellh/cli"
)

// Export
type Export struct{}

// NewExport creates a new instance of Delete
func NewExport() cli.CommandFactory {
	return func() (cli.Command, error) {
		return &Export{}, nil
	}
}

// Run shows a given resource.
func (e *Export) Run(args []string) int {
	if flags.Has("help", args) || len(args) < 1 {
		fmt.Print(e.Help())
		return 1
	}

	// process the subcommand and it's options
	switch args[0] {
	case "csv":
		e.csv()
	case "json":
		e.json()
	default:
		fmt.Println("ERROR: invalid option for show\n")
	}

	return 1
}

// Help provides full help inforamation for the subcommand
func (e *Export) Help() string {
	return `Usage: todo-view export <option> <arguments> 
  Show a resource
Options:
  csv                Display the todo-view data in csv format
  json               Display the todo-view data in json format
  
`
}

// Synopsis provides a brief description of the command
func (e *Export) Synopsis() string {
	return "Export todo-view data in a given format"
}

// csv exports the todo-view data in cvs format
func (e *Export) csv() {
	fmt.Print("\ntodo-view export: cvs\n\n")
	todos, err := search()
	if err != nil {
		log.Fatalln(err)
	}
	for _, todo := range todos {
		fmt.Fprintf(os.Stdout, "%s,%s,%s,%v,%d\n", todo.User(), todo.File(), todo.Message(), todo.Timestamp(), todo.Weight())
	}
}

// json exports the todo-view data in json format
func (e *Export) json() {
	fmt.Print("\ntodo-view export: json\n\n")
	w := NewTabWriter()
	defer w.Flush()

	fmt.Fprintf(w, "\n")
}
