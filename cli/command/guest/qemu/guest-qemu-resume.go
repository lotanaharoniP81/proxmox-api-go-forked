package qemu

import (
	"github.com/lotanaharoniP81/proxmox-api-go-forked/cli"
	"github.com/lotanaharoniP81/proxmox-api-go-forked/proxmox"
	"github.com/spf13/cobra"
)

var qemu_resumeCmd = &cobra.Command{
	Use:   "resume GUESTID",
	Short: "Resumes the specified guest",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		vmr := proxmox.NewVmRef(cli.ValidateIntIDset(args, "GuestID"))
		c := cli.NewClient()
		_, err = c.ResumeVm(vmr)
		if err == nil {
			cli.PrintGuestStatus(qemuCmd.OutOrStdout(), vmr.VmId(), "resumed")
		}
		return
	},
}

func init() {
	qemuCmd.AddCommand(qemu_resumeCmd)
}
