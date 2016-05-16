package command

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/fatih/flags"
	"github.com/mitchellh/cli"
)

// Export
type Export struct{}

// Exporter is an interface with an export method
type Exporter interface {
	Export()
}

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
	case "jira":
		e.jira()
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
		fmt.Fprintf(os.Stdout, "%s,%s,%s,%v,%d\n",
			todo.User(), todo.File(), todo.Message(), todo.Timestamp(), todo.Weight())
	}
}

// json exports the todo-view data in json format
func (e *Export) json() {
	fmt.Print("\ntodo-view export: json\n\n")
	w := NewTabWriter()
	defer w.Flush()

	todos, err := search()
	if err != nil {
		log.Fatalln(err)
	}

	m := make(map[string][]Todo)
	m["todos"] = todos

	b, err := json.Marshal(m["todos"][0])
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintln(w, string(b))
}

// jira exports the todo-view data in jira format
func (e *Export) jira() {
	fmt.Print("\ntodo-view export: jira\n\n")
	todos, err := search()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintln(os.Stdout, "Summary,Assignee,Reporter,Priority")
	for _, todo := range todos {
		fmt.Fprintf(os.Stdout, "%s,%s,%s,%d\n",
			todo.Message(), todo.User(), todo.User(), todo.Weight())
	}

	/*var t []Exporter
	for _, todo := range todos {
		nj := NewJira(todo.Message(), todo.User(), todo.Weight())
		t = append(t, nj)
	}
	nj.Export()*/
}

// Jira holds the data necessary to export data in Jira format
type Jira struct {
	Header      string
	Summary     string
	Assignee    string
	Reporter    string
	IssueType   int
	Status      string
	Description string
	Priority    int
}

// NewJira creates a new reference to a Jira type
func NewJira(msg, user string, weight int) *Jira {
	return &Jira{
		Header:   "Summary,Assignee,Reporter,Priority",
		Summary:  msg,
		Assignee: user,
		Reporter: user,
		Priority: weight,
	}
}

// Export prints out the Jira formatted data
func (j *Jira) Export() {
	fmt.Fprintln(os.Stdout, j.Header)
}
