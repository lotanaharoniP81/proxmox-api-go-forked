package delete

import (
	"fmt"

	"github.com/lotanaharoniP81/proxmox-api-go-forked/cli"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "With this command you can delete existing items from proxmox",
}

func init() {
	cli.RootCmd.AddCommand(deleteCmd)
}

func deleteID(args []string, IDtype string) (err error) {
	var exitStatus string
	id := cli.RequiredIDset(args, 0, IDtype+"ID")
	c := cli.NewClient()
	switch IDtype {
	case "AcmeAccount":
		exitStatus, err = c.DeleteAcmeAccount(id)
	case "MetricServer":
		err = c.DeleteMetricServer(id)
	case "Pool":
		err = c.DeletePool(id)
	case "Storage":
		err = c.DeleteStorage(id)
	case "User":
		err = c.DeleteUser(id)
	}
	if err != nil {
		if exitStatus != "" {
			err = fmt.Errorf("error deleting %s (%s): %v, error status: %s ", IDtype, id, err, exitStatus)
		}
		return
	}
	cli.PrintItemDeleted(deleteCmd.OutOrStdout(), id, IDtype)
	return
}
