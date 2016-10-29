package command

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

// Version represents the version command
type Version struct {
	version string
}

// NewVersion creates a new version command
func NewVersion(version string) cli.CommandFactory {
	return func() (cli.Command, error) {
		return &Version{
			version: version,
		}, nil
	}
}

// Run shows the current version
func (v *Version) Run(args []string) int {
	fmt.Fprintln(os.Stderr, v.version)
	return 0
}

// Help displays the below string
func (v *Version) Help() string {
	return "Prints the todo-view version"
}

// Synopsis provides a brief description of the command
func (v *Version) Synopsis() string {
	return "Prints the todo-view version"
}
