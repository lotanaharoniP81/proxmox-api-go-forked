package storagesubtests

import (
	"encoding/json"

	_ "github.com/lotanaharoniP81/proxmox-api-go-forked/cli/command/commands"
	"github.com/lotanaharoniP81/proxmox-api-go-forked/proxmox"
)

func CloneJson(jsonStruct proxmox.ConfigStorage) *proxmox.ConfigStorage {
	s := &proxmox.ConfigStorage{}
	ori, _ := json.Marshal(jsonStruct)
	json.Unmarshal(ori, s)
	return s
}

func InlineMarshal(jsonStruct *proxmox.ConfigStorage) string {
	j, _ := json.Marshal(jsonStruct)
	return string(j)
}
