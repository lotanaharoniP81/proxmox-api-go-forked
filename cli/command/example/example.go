package example

import (
	"github.com/lotanaharoniP81/proxmox-api-go-forked/cli"
	"github.com/spf13/cobra"
)

var exampleCmd = &cobra.Command{
	Use:   "example",
	Short: "This function show examples of fully populated config files",
}

func init() {
	cli.RootCmd.AddCommand(exampleCmd)
}
