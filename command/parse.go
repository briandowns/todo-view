package command

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"time"

	"github.com/fatih/flags"
	"github.com/mitchellh/cli"
)

// regex holds the pattern necessary to match the todo in the
// files parsed
var regex = regexp.MustCompile(`TODO\((?P<user>[a-z].+)\)(?P<msg>.+)(?P<timestamp>\d\d\d\d\D?\d\d\D?\d\d\D?\d\d\D?\d\d\D?(\d\d\.?(\d*))?(\d\d(:\d\d)?)?).(?P<weight>\d)`)

// Parse type for command
type Parse struct{}

// NewParse creates a new Parse reference
func NewParse() cli.CommandFactory {
	return func() (cli.Command, error) {
		return &Parse{}, nil
	}
}

// Run shows a given resource.
func (p *Parse) Run(args []string) int {
	if flags.Has("help", args) || len(args) < 1 {
		fmt.Print(p.Help())
		return 1
	}

	// process the subcommand and it's options
	switch args[0] {
	case "by-user":
		p.byUser()
	case "by-file":
		p.byFile()
	case "by-date":
		p.byDate()
	case "by-weight":
		p.byWeight()
	default:
		fmt.Println("ERROR: invalid option for parse\n")
	}

	return 1
}

// Help provides full help inforamation for the subcommand
func (p *Parse) Help() string {
	return `Usage: todo-view parse <option> <arguments> 
  Parse a source tree
Options:
  by-user            Parse todo's by user
  by-file            Parse todo's by file   
  by-date            Parse todo's by date
  by-weight          Parse todo's by weight
  
`
}

// Synopsis provides a brief description of the command
func (p *Parse) Synopsis() string {
	return "Parse source and display results"
}

// printOutput prints the given output to screen
func printOutput(ot string, t Swapper) {
	fmt.Printf("\nTodo's by %s:\n\n", ot)
	w := NewTabWriter()
	defer w.Flush()
	switch todoerType := t.(type) {
	case UserTodos:
		for _, todo := range todoerType {
			fmt.Fprintf(w, "%s\t%s\t%s\t%v\t%d\n",
				todo.User(), todo.File(), todo.Message(), todo.Timestamp(), todo.Weight())
		}
	case FileTodos:
		for _, todo := range todoerType {
			fmt.Fprintf(w, "%s\t%s\t%s\t%v\t%d\n",
				todo.User(), todo.File(), todo.Message(), todo.Timestamp(), todo.Weight())
		}
	case TimestampTodos:
		for _, todo := range todoerType {
			fmt.Fprintf(w, "%s\t%s\t%s\t%v\t%d\n",
				todo.User(), todo.File(), todo.Message(), todo.Timestamp(), todo.Weight())
		}
	case WeightTodos:
		for _, todo := range todoerType {
			fmt.Fprintf(w, "%s\t%s\t%s\t%v\t%d\n",
				todo.User(), todo.File(), todo.Message(), todo.Timestamp(), todo.Weight())
		}
	}
	fmt.Fprintf(w, "\n")
}

// byUser outputs the data by user
func (p *Parse) byUser() {
	todos, err := search()
	if err != nil {
		log.Fatalln(err)
	}
	userTodos := make(UserTodos, len(todos))
	for i := 0; i <= len(todos)-1; i++ {
		userTodos[i] = todos[i]
	}
	sort.Sort(userTodos)
	printOutput("user", userTodos)
}

// byFile outputs the data by file
func (p *Parse) byFile() {
	todos, err := search()
	if err != nil {
		log.Fatalln(err)
	}
	fileTodos := make(FileTodos, len(todos))
	for i := 0; i <= len(todos)-1; i++ {
		fileTodos[i] = todos[i]
	}
	sort.Sort(fileTodos)
	printOutput("file", fileTodos)
}

// byDate outputs the data by date
func (p *Parse) byDate() {
	todos, err := search()
	if err != nil {
		log.Fatalln(err)
	}
	timestampTodos := make(TimestampTodos, len(todos))
	for i := 0; i <= len(todos)-1; i++ {
		timestampTodos[i] = todos[i]
	}
	sort.Sort(timestampTodos)
	printOutput("file", timestampTodos)
}

// byWeight outputs the data by weight
func (p *Parse) byWeight() {
	todos, err := search()
	if err != nil {
		log.Fatalln(err)
	}
	weightTodos := make(WeightTodos, len(todos))
	for i := 0; i <= len(todos)-1; i++ {
		weightTodos[i] = todos[i]

	}
	sort.Sort(weightTodos)
	printOutput("file", weightTodos)
}

// search the directory path recursively
func search() ([]Todo, error) {
	var todos []Todo
	fileList := []string{}
	err := filepath.Walk(".", func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	for _, file := range fileList {
		fh, err := os.Open(file)
		if err != nil {
			log.Fatalln(err)
		}
		defer fh.Close()

		fs := bufio.NewScanner(fh)

		for fs.Scan() {
			s := fs.Text()

			match := regex.FindStringSubmatch(s)

			if match != nil {
				var todo Todo
				//TODO(briandowns) have this controlled by a CLI flag 2016-05-13T16:14 2
				/*fp, err := filepath.Abs(fh.Name())
				if err != nil {
					return nil, err
				}
				todo.File = fp*/
				todo.file = file
				for i, name := range regex.SubexpNames() {
					if i == 0 || name == "" {
						continue
					}

					switch name {
					case "user":
						todo.user = match[i]
					case "msg":
						todo.message = match[i]
					case "timestamp":
						ts, err := time.Parse(Format, match[i])
						if err != nil {
							return nil, err
						}
						todo.timestamp = ts
					case "weight":
						s, err := strconv.Atoi(match[i])
						if err != nil {
							return nil, err
						}
						todo.weight = s
					}
				}
				todos = append(todos, todo)
			}
		}
	}
	return todos, nil
}
