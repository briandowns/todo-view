package command

import (
	"os"
	"text/tabwriter"
)

// NewTabWriter sets up a new and configured tabwriter
func NewTabWriter() *tabwriter.Writer {
	t := new(tabwriter.Writer)
	t.Init(os.Stdout, 0, 8, 1, '\t', 0)
	return t
}
