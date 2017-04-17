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
	"strings"
	"time"

	"github.com/fatih/flags"
	"github.com/mitchellh/cli"
)

// regex holds the pattern necessary to match the todo in the
// files parsed
var regex = regexp.MustCompile(`TODO\((?P<user>[a-z].+)\)(?P<msg>.+)(?P<timestamp>\d\d\d\d\D?\d\d\D?\d\d\D?\d\d\D?\d\d\D?(\d\d\.?(\d*))?(\d\d(:\d\d)?)?).(?P<priority>\d)`)

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
		if len(args) == 2 {
			switch args[1] {
			case "-d":
				p.byUser(false)
				return 1
			}
		}
		p.byUser(true)
	case "by-file":
		if len(args) == 2 {
			switch args[1] {
			case "-d":
				p.byFile(false)
				return 1
			}
		}
		p.byFile(true)
	case "by-date":
		if len(args) == 2 {
			switch args[1] {
			case "-d":
				p.byDate(false)
				return 1
			}
		}
		p.byDate(true)
	case "by-priority":
		if len(args) == 2 {
			switch args[1] {
			case "-d":
				p.byPriority(false)
				return 1
			}
		}
		p.byPriority(true)
	default:
		fmt.Print("ERROR: invalid option for parse\n\n")
	}
	return 1
}

// Help provides full help inforamation for the subcommand
func (p *Parse) Help() string {
	return `Usage: todo-view parse <option> <arguments> 
  Parse a source tree
  
Options:
  by-user       [-d decending]        Parse todo's by user
  by-file       [-d decending]        Parse todo's by file   
  by-date       [-d decending]        Parse todo's by date
  by-priority   [-d decending]        Parse todo's by priority
  
`
}

// Synopsis provides a brief description of the command
func (p *Parse) Synopsis() string {
	return "Parse source and display results"
}

// printOutput prints the given output to screen
func printOutput(ot string, t sort.Interface) {
	fmt.Printf("\nTodo's by %s:\n\n", ot)
	w := NewTabWriter()
	defer w.Flush()
	switch todoerType := t.(type) {
	case UserTodos:
		for _, todo := range todoerType {
			fmt.Fprintf(w, "%s\t%s\t%s\t%v\t%d\n",
				todo.User(), todo.File(), todo.Message(), todo.Timestamp(), todo.Priority())
		}
	case FileTodos:
		for _, todo := range todoerType {
			fmt.Fprintf(w, "%s\t%s\t%s\t%v\t%d\n",
				todo.User(), todo.File(), todo.Message(), todo.Timestamp(), todo.Priority())
		}
	case TimestampTodos:
		for _, todo := range todoerType {
			fmt.Fprintf(w, "%s\t%s\t%s\t%v\t%d\n",
				todo.User(), todo.File(), todo.Message(), todo.Timestamp(), todo.Priority())
		}
	case PriorityTodos:
		for _, todo := range todoerType {
			fmt.Fprintf(w, "%s\t%s\t%s\t%v\t%d\n",
				todo.User(), todo.File(), todo.Message(), todo.Timestamp(), todo.Priority())
		}
	}
	fmt.Fprintf(w, "\n")
}

// byUser outputs the data by user
func (p *Parse) byUser(decending bool) {
	todos, err := search()
	if err != nil {
		log.Fatalln(err)
	}
	userTodos := make(UserTodos, len(todos))
	for i := 0; i <= len(todos)-1; i++ {
		userTodos[i] = todos[i]
	}
	switch decending {
	case true:
		sort.Sort(userTodos)
		printOutput("user", userTodos)
		return
	case false:
		sort.Sort(sort.Reverse(userTodos))
		printOutput("user", userTodos)
		return
	}
}

// byFile outputs the data by file
func (p *Parse) byFile(decending bool) {
	todos, err := search()
	if err != nil {
		log.Fatalln(err)
	}
	fileTodos := make(FileTodos, len(todos))
	for i := 0; i <= len(todos)-1; i++ {
		fileTodos[i] = todos[i]
	}
	switch decending {
	case true:
		sort.Sort(fileTodos)
		printOutput("file", fileTodos)
		return
	case false:
		sort.Sort(sort.Reverse(fileTodos))
		printOutput("file", fileTodos)
		return
	}
}

// byDate outputs the data by date
func (p *Parse) byDate(decending bool) {
	todos, err := search()
	if err != nil {
		log.Fatalln(err)
	}
	timestampTodos := make(TimestampTodos, len(todos))
	for i := 0; i <= len(todos)-1; i++ {
		timestampTodos[i] = todos[i]
	}
	switch decending {
	case true:
		sort.Sort(timestampTodos)
		printOutput("priority", timestampTodos)
		return
	case false:
		sort.Sort(sort.Reverse(timestampTodos))
		printOutput("date", timestampTodos)
		return
	}
}

// byPriority outputs the data by priority
func (p *Parse) byPriority(decending bool) {
	todos, err := search()
	if err != nil {
		log.Fatalln(err)
	}
	priorityTodos := make(PriorityTodos, len(todos))
	for i := 0; i <= len(todos)-1; i++ {
		priorityTodos[i] = todos[i]
	}
	switch decending {
	case true:
		sort.Sort(priorityTodos)
		printOutput("priority", priorityTodos)
		return
	case false:
		sort.Sort(sort.Reverse(priorityTodos))
		printOutput("priority", priorityTodos)
		return
	}
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

		fs := bufio.NewScanner(fh)

		for fs.Scan() {
			s := fs.Text()

			match := regex.FindStringSubmatch(s)

			if match != nil {
				var todo Todo
				todo.file = file

				for i, name := range regex.SubexpNames() {
					if i == 0 || name == "" {
						continue
					}

					switch name {
					case "user":
						todo.user = match[i]
					case "msg":
						todo.message = strings.TrimSpace(match[i])
					case "timestamp":
						ts, err := time.Parse(Format, match[i])
						if err != nil {
							return nil, err
						}
						todo.timestamp = ts
					case "priority":
						s, err := strconv.Atoi(match[i])
						if err != nil {
							return nil, err
						}
						todo.priority = s
					}
				}
				todos = append(todos, todo)
			}
		}
		fh.Close()
	}
	return todos, nil
}
