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
type Parse struct {
}

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

//TODO(briandowns) combine functionality to eliminate all the repitition 2016-05-13T16:14 1

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

	fmt.Print("\nTodo's by user:\n\n")
	w := NewTabWriter()
	defer w.Flush()
	for _, todo := range userTodos {
		fmt.Fprintf(w, "%s\t%s\t%s\t%v\t%d\n", todo.User, todo.File, todo.Message, todo.Timestamp, todo.Weight)
	}

	fmt.Fprintf(w, "\n")
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

	fmt.Print("\nTodo's by file:\n\n")
	w := NewTabWriter()
	defer w.Flush()
	for _, todo := range fileTodos {
		fmt.Fprintf(w, "%s\t%s\t%s\t%v\t%d\n", todo.User, todo.File, todo.Message, todo.Timestamp, todo.Weight)
	}

	fmt.Fprintf(w, "\n")
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

	fmt.Print("\nTodo's by date:\n\n")
	w := NewTabWriter()
	defer w.Flush()
	for _, todo := range timestampTodos {
		fmt.Fprintf(w, "%s\t%s\t%s\t%v\t%d\n", todo.User, todo.File, todo.Message, todo.Timestamp, todo.Weight)
	}

	fmt.Fprintf(w, "\n")
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

	fmt.Print("\nTodo's by weight:\n\n")
	w := NewTabWriter()
	defer w.Flush()
	for _, todo := range weightTodos {
		fmt.Fprintf(w, "%s\t%s\t%s\t%v\t%d\n", todo.User, todo.File, todo.Message, todo.Timestamp, todo.Weight)
	}

	fmt.Fprintf(w, "\n")
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
				todo.File = file
				for i, name := range regex.SubexpNames() {
					if i == 0 || name == "" {
						continue
					}

					switch name {
					case "user":
						todo.User = match[i]
					case "msg":
						todo.Message = match[i]
					case "timestamp":
						ts, err := time.Parse(Format, match[i])
						if err != nil {
							return nil, err
						}
						todo.Timestamp = ts
					case "weight":
						s, err := strconv.Atoi(match[i])
						if err != nil {
							return nil, err
						}
						todo.Weight = s
					}
				}
				todos = append(todos, todo)
			}
		}
	}

	return todos, nil
}
