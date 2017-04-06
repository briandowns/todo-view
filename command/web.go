package command

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/fatih/flags"
	"github.com/mitchellh/cli"
)

// Html  data type
type Web struct{}

// NewHtml creates a new instance
func NewWeb() cli.CommandFactory {
	return func() (cli.Command, error) {
		return &Web{}, nil
	}
}

// Run a web server
func (w *Web) Run(args []string) int {
	if flags.Has("help", args) || len(args) < 1 {
		fmt.Print(w.Help())
		return 1
	}

	// process the subcommand and it's options
	switch args[0] {
	case "port":
		if len(args) == 2 {
			w.Serve(":" + args[1])
		} else {
			w.Serve("")
		}
	default:
		fmt.Print("ERROR: invalid option for web\n\n")
	}

	return 1
}

// Synopsis provides a brief description of the command
func (w *Web) Synopsis() string {
	return "Show all todos in a web browser"
}

// Help provides full help inforamation for the subcommand
func (w *Web) Help() string {
	return `Usage: todo-view web <option> <arguments> 
  Show a resource
  
Options:
  port               Display all todos in a browser on selected port (default 9876)
`
}

var todoData = template.FuncMap{
	"user": func(t *Todo) string {
		return t.User()
	},
	"file": func(t *Todo) string {
		return t.File()
	},
	"message": func(t *Todo) string {
		return t.Message()
	},
	"timestamp": func(t *Todo) time.Time {
		return t.Timestamp()
	},
	"priority": func(t *Todo) int {
		return t.Priority()
	},
}

func newTemplate() *template.Template {
	return template.Must(template.New("").Funcs(todoData).Parse(htmlTemplate))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tpl := newTemplate()
	todos, err := search()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = tpl.Execute(w, todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (w *Web) Serve(port string) {
	if port == "" {
		port = ":9876"
	}
	fmt.Print("\ntodo-view web: \n\n")
	fmt.Println("Serving on http://localhost" + port)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(port, nil))
}

var htmlTemplate = `
  <!DOCTYPE html>
  <html>
    <head>
      <style type="text/css">
        @import url(https://fonts.googleapis.com/css?family=Raleway);
        body {
          background-color: white;
          font-family: 'Raleway', sans-serif;
        }

        .info-text {
          width: 100%;
          background-color: #576876;
          color: white;
          border: 1px solid black;
          margin-bottom: 10px;
        }
        .info-text p {
          margin: 10px 0 10px 10px;
        }
        .contenedor {
          display: table;
          border: 1px solid #7f7f7f;
          width: 100%;
          border-radius: 5px;
          text-align: center;
          margin: 0 auto;
        }
        .row {
          display: table-row;
        }
        .cell {
          display: table-cell;
          border: 1px solid #7f7f7f;
          vertical-align: middle;
          padding: 10px;
        }
        .row.header {
          background-color: #576876;
          color: white;
        }
      </style>
    </head>
    <body>
      <div class="info-header">
        <div class="info-text">
          <p>TODO VIEW LIST</p>
        </div>
      </div>
      <div class="contenedor">
        <div class="row header">
          <div class="cell">User</div>
          <div class="cell">File</div>
          <div class="cell">Message</div>
          <div class="cell">Timestamp</div>
          <div class="cell">Priority</div>
        </div>
          {{range .}}
            <div class="row">
              <div class="cell">{{user .}}</div>
              <div class="cell">{{file .}}</div>
              <div class="cell">{{message .}}</div>
              <div class="cell">{{timestamp .}}</div>
              <div class="cell">{{priority .}}</div>
            </div>
          {{end}}
        </div>
      </div>
    </body>
  </html>
`
