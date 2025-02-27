package cli_acmeaccount_test

import (
	"testing"

	_ "github.com/lotanaharoniP81/proxmox-api-go-forked/cli/command/commands"
	cliTest "github.com/lotanaharoniP81/proxmox-api-go-forked/test/cli"
)

func Test_AcmeAccount_0_Cleanup(t *testing.T) {
	Test := cliTest.Test{
		ReqErr:      true,
		ErrContains: "test-0",
		Args:        []string{"-i", "delete", "acmeaccount", "test-0"},
	}
	Test.StandardTest(t)
}

func Test_AcmeAccount_0_Set(t *testing.T) {
	Test := cliTest.Test{
		InputJson: `
{
	"contact": [
		"a@nonexistantdomain.com",
		"b@nonexistantdomain.com",
		"c@nonexistantdomain.com",
		"d@nonexistantdomain.com"
	],
	"directory": "https://acme-staging-v02.api.letsencrypt.org/directory",
	"tos": true
}`,
		Expected: "(test-0)",
		Contains: true,
		Args:     []string{"-i", "create", "acmeaccount", "test-0"},
	}
	Test.StandardTest(t)
}

func Test_AcmeAccount_0_Get(t *testing.T) {
	Test := cliTest.Test{
		OutputJson: `
{
	"name": "test-0",
	"contact": [
		"a@nonexistantdomain.com",
		"b@nonexistantdomain.com",
		"c@nonexistantdomain.com",
		"d@nonexistantdomain.com"
	],
	"directory": "https://acme-staging-v02.api.letsencrypt.org/directory",
	"tos": true
}`,
		Args: []string{"-i", "get", "acmeaccount", "test-0"},
	}
	Test.StandardTest(t)
}

func Test_AcmeAccount_0_Delete(t *testing.T) {
	Test := cliTest.Test{
		Expected: "",
		ReqErr:   false,
		Args:     []string{"-i", "delete", "acmeaccount", "test-0"},
	}
	Test.StandardTest(t)
}
