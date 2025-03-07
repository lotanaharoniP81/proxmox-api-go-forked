package cli_storage_test

import (
	"testing"

	_ "github.com/lotanaharoniP81/proxmox-api-go-forked/cli/command/commands"
	"github.com/lotanaharoniP81/proxmox-api-go-forked/proxmox"
	cliTest "github.com/lotanaharoniP81/proxmox-api-go-forked/test/cli"
	storagesubtests "github.com/lotanaharoniP81/proxmox-api-go-forked/test/cli/Storage/storage-sub-tests"
)

func Test_Storage_SMB_0_Cleanup(t *testing.T) {
	storagesubtests.Cleanup("smb-test-0", t)
}

func Test_Storage_SMB_0_Create_Full(t *testing.T) {
	s := storagesubtests.CloneJson(storagesubtests.SMBFull)
	s.SMB.Password = proxmox.PointerString("Enter123!")
	storagesubtests.Create(s, "smb-test-0", t)
}

func Test_Storage_SMB_0_Get_Full(t *testing.T) {
	storagesubtests.SMBGetFull("smb-test-0", t)
}

func Test_Storage_SMB_0_Update_Empty(t *testing.T) {
	cliTest.SetEnvironmentVariables()
	s := storagesubtests.CloneJson(storagesubtests.SMBEmpty)
	s.BackupRetention = &proxmox.ConfigStorageBackupRetention{}
	storagesubtests.Update(s, "smb-test-0", t)
}

func Test_Storage_SMB_0_Get_Empty(t *testing.T) {
	storagesubtests.SMBGetEmpty("smb-test-0", t)
}

func Test_Storage_SMB_0_Delete(t *testing.T) {
	storagesubtests.Delete("smb-test-0", t)
}
