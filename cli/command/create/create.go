package create

import (
	"github.com/lotanaharoniP81/proxmox-api-go-forked/cli"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "With this command you can create new items in proxmox",
}

func init() {
	cli.RootCmd.AddCommand(CreateCmd)
}
