package main

import (
	"context"
	"encoding/json"
	"time"

	_ "github.com/Telmate/proxmox-api-go/cli/command/commands"
	"github.com/Telmate/proxmox-api-go/proxmox"
	"log"
	"regexp"
)

func main() {
	ca := []byte(caPem)
	cert := []byte(certPem)
	key := []byte(certKey)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*30)
	defer cancel()

	c, err := Connect(ctx, PmApiUrl, ca, cert, key)
	if err != nil {
		return
	}

	//// get the information about the hosts
	//nodeList, err := getNodeList(c)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//// information about all the storage
	//list := nodeList["data"].([]interface{})
	//for _, node := range list {
	//	node2 := node.(map[string]interface{})
	//	nodeName := node2["id"].(string)[5:]
	//	storage, err := c.GetStorage(nodeName)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	for _, eachStorage := range storage {
	//		fmt.Println(eachStorage.(map[string]interface{})["storage"])
	//	}
	//}

	//// check if the vm exist
	//ifExist, err := ifVMIdExists(c, 100)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(ifExist)
	//
	//// get vm status
	//status, err := getVMInfo(c, 100)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(status["status"])

	//// stop the vm
	//err = stopVM(c, 100)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//// start the vm
	//err = startVM(c, 100)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//// check if the vm is still in storage (for example in the 'local-lvm'). Returns the data if the VM was stopped
	//stor, err := getStorageContent(c, "local-lvm", 100)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(stor)

	//// delete the vm
	//exitStatus, err := deleteVM(c, 100)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(exitStatus)
}

func createQemu(c *proxmox.Client, vmID int, fConfigFile string) error {
	configTemp, err := GetConfig(fConfigFile)
	if err != nil {
		return err
	}
	config, err := proxmox.NewConfigQemuFromJson(configTemp)
	if err != nil {
		return err
	}
	vmr := proxmox.NewVmRef(vmID)
	vmr.SetNode(host)
	return config.CreateVm(vmr, c)
}

func failError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var rxUserRequiresToken = regexp.MustCompile("[a-z0-9]+@[a-z0-9]+![a-z0-9]+")

func userRequiresAPIToken(userID string) bool {
	return rxUserRequiresToken.MatchString(userID)
}

func getVMList(c *proxmox.Client) (string, error) {
	vms, err := c.GetVmList()
	if err != nil {
		return "", err
	}
	vmList, err := json.Marshal(vms)
	if err != nil {
		return "", err
	}
	return string(vmList), nil
}

func getStorage(c *proxmox.Client, storageID string) (string, error) {
	config, err := proxmox.NewConfigStorageFromApi(storageID, c)
	if err != nil {
		return "", err
	}
	cj, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return "", err
	}
	return string(cj), nil
}

func getStorageContent(c *proxmox.Client, storageID string, vmID int) (map[string]interface{}, error) {
	vmr := proxmox.NewVmRef(vmID)
	return c.GetStorageContent(vmr, storageID)
}

func getStorageStatus(c *proxmox.Client, storageID string, vmID int) (map[string]interface{}, error) {
	vmr := proxmox.NewVmRef(vmID)
	return c.GetStorageStatus(vmr, storageID)
}

func getStorageFull(c *proxmox.Client) ([]interface{}, error) {
	return c.GetStorageFull()
}

func deleteStorage(c *proxmox.Client, storageID string) error {
	return c.DeleteStorage(storageID)
}

func stopVM(c *proxmox.Client, vmID int) error {
	vmr := proxmox.NewVmRef(vmID)
	_, err := c.StopVm(vmr)
	return err
}

func startVM(c *proxmox.Client, vmID int) error {
	vmr := proxmox.NewVmRef(vmID)
	_, err := c.StartVm(vmr)
	return err
}

func destroyVM(c *proxmox.Client, vmID int) error {
	vmr := proxmox.NewVmRef(vmID)
	_, err := c.StopVm(vmr)
	if err != nil {
		return err
	}
	_, err = c.DeleteVm(vmr)
	return err
}

func ifVMIdExists(c *proxmox.Client, vmID int) (bool, error) {
	ifVMIdExists, err := c.VMIdExists(vmID)
	if err != nil {
		return false, err
	}
	return ifVMIdExists, nil
}

func resetVM(c *proxmox.Client, vmID int) error {
	vmr := proxmox.NewVmRef(vmID)
	_, err := c.ResetVm(vmr)
	return err
}

func getStorageList(c *proxmox.Client) (string, error) {
	storage, err := c.GetStorageList()
	if err != nil {
		return "", err
	}
	storageList, err := json.Marshal(storage)
	if err != nil {
		return "", err
	}
	return string(storageList), nil
}

func getNodeList(c *proxmox.Client) (map[string]interface{}, error) {
	return c.GetNodeList()
}

func getVMInfo(c *proxmox.Client, vmID int) (map[string]interface{}, error) {
	vmr := proxmox.NewVmRef(vmID)
	info, err := c.GetVmInfo(vmr)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func getVMState(c *proxmox.Client, vmID int) (map[string]interface{}, error) {
	vmr := proxmox.NewVmRef(vmID)
	info, err := c.GetVmState(vmr)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func createStorage(c *proxmox.Client, storageID string, fConfigFile string) error {
	configTemp, err := GetConfig(fConfigFile)
	if err != nil {
		return err
	}
	config, err := proxmox.NewConfigStorageFromJson(configTemp)
	if err != nil {
		return err
	}
	return config.CreateWithValidate(storageID, c)
}

func deleteVM(c *proxmox.Client, vmID int) (exitStatus string, err error) {
	vmr := proxmox.NewVmRef(vmID)
	return c.DeleteVm(vmr)
}
