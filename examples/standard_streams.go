package examples

import "github.com/ubhisat/cmd"

// CreateNewCommandWithStandardStream create new standard stream example
func CreateNewCommandWithStandardStream() {
	c := cmd.NewCommand("echo hello; sleep 1; echo another;", cmd.WithStandardStreams)
	c.Execute()
}
