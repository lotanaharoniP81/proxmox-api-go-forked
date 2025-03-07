package qemu

import (
	"github.com/lotanaharoniP81/proxmox-api-go-forked/cli"
	"github.com/lotanaharoniP81/proxmox-api-go-forked/proxmox"
	"github.com/spf13/cobra"
)

var qemu_pauseCmd = &cobra.Command{
	Use:   "pause GUESTID",
	Short: "Pauses the specified guest",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		vmr := proxmox.NewVmRef(cli.ValidateIntIDset(args, "GuestID"))
		c := cli.NewClient()
		_, err = c.PauseVm(vmr)
		if err == nil {
			cli.PrintGuestStatus(qemuCmd.OutOrStdout(), vmr.VmId(), "paused")
		}
		return
	},
}

func init() {
	qemuCmd.AddCommand(qemu_pauseCmd)
}
