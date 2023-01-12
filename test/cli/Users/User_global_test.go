package cli_user_test

import (
	_ "github.com/lotanaharoniP81/proxmox-api-go-forked/cli/command/commands"
	cliTest "github.com/lotanaharoniP81/proxmox-api-go-forked/test/cli"
	"testing"
)

func Test_User_List(t *testing.T) {
	Test := cliTest.Test{
		Expected: `"userid":"root@pam"`,
		ReqErr:   false,
		Contains: true,
		Args:     []string{"-i", "list", "users"},
	}
	Test.StandardTest(t)
}
