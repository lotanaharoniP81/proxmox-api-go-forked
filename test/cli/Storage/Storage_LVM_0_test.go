package cli_storage_test

import (
	"testing"

	_ "github.com/lotanaharoniP81/proxmox-api-go-forked/cli/command/commands"
	"github.com/lotanaharoniP81/proxmox-api-go-forked/proxmox"
	storagesubtests "github.com/lotanaharoniP81/proxmox-api-go-forked/test/cli/Storage/storage-sub-tests"
)

func Test_Storage_LVM_0_Cleanup(t *testing.T) {
	storagesubtests.Cleanup("lvm-test-0", t)
}

func Test_Storage_LVM_0_Create_Full(t *testing.T) {
	s := storagesubtests.CloneJson(storagesubtests.LVMFull)
	storagesubtests.Create(s, "lvm-test-0", t)
}

func Test_Storage_LVM_0_Get_Full(t *testing.T) {
	storagesubtests.LVMGetFull("lvm-test-0", t)
}

func Test_Storage_LVM_0_Update_Empty(t *testing.T) {
	s := storagesubtests.CloneJson(storagesubtests.LVMEmpty)
	s.BackupRetention = &proxmox.ConfigStorageBackupRetention{}
	storagesubtests.Update(s, "lvm-test-0", t)
}

func Test_Storage_LVM_0_Get_Empty(t *testing.T) {
	storagesubtests.LVMGetEmpty("lvm-test-0", t)
}

func Test_Storage_LVM_0_Delete(t *testing.T) {
	storagesubtests.Delete("lvm-test-0", t)
}
