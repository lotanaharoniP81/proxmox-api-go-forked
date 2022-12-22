package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/Telmate/proxmox-api-go/cli/command/commands"
	"github.com/Telmate/proxmox-api-go/proxmox"
	"log"
	"regexp"
)

//disk {
//size               = var.disks.size #The diskSize of the VM
//storage            = var.disks.storage # Storage Pool to deploy the VM on
//type               = var.disks.type # Disk virtualtzion type
//}

const cloneVMconst = `{
  "name": "my.vm",
  "desc": "Test proxmox-api-go clone",
  "storage": "local",
  "memory": 2048,
  "cores": 1,
  "sockets": 1,
  "fullclone": 1,
"disk": {
"0":{
"size": "50G",
"storage": "local-lvm",
"type": "scsi0"
}
}
}`

const memoryThreshold = 80
const storageThreshold = 10

const cloneVMCloudInit = `{
  "name": "cloudinit.test.com",
  "desc": "Test proxmox-api-go clone",
  "storage": "local",
  "memory": 2048,
  "cores": 2,
  "sockets": 1,
  "ipconfig": {
    "0": "gw=10.0.2.2,ip=10.0.2.17/24"
  },
  "sshkeys": "...",
  "nameserver": "8.8.8.8",
  ""
}`

//const createQemuConfig = `{
//	"name": "LOTAN-TEST",
//	"memory":2048,
//	"cores":1,
//	"sockets":1,
//	"kvm":false,
//	"iso":"local:iso/ubuntu-20.04.5-live-server-amd64.iso"}`

const createQemuConfig = `{
  "name": "my.vm",
  "desc": "Test proxmox-api-go clone",
  "storage": "local",
  "memory": 2048,
  "cores": 1,
  "sockets": 1,
  "fullclone": 1
}`

//const createQemuConfig = "{\n  \"name\": \"rtmLr4J\",\n  \"memory\": 2048,\n  \"cores\": 1,\n  \"sockets\": 1,\n  \"kvm\": false,\n  \"iso\": \"local:iso/ubuntu-22.04.1-live-server-amd64.iso\"\n}"

func main() {
	ca := []byte(caPem)
	cert := []byte(certPem)
	key := []byte(certKey)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*30)
	defer cancel()

	//_, err := SSHConnect(fmt.Sprintf("%s:%d", "64.226.130.1", 10080),"admin")

	c, err := Connect(ctx, PmApiUrl, ca, cert, key)
	if err != nil {
		return
	}

	//// get vm list
	//list, err := getVMList(c)
	//if err != nil {
	//	fmt.Printf("error: %v", err)
	//}
	//fmt.Println(list)

	//// get the information about the hosts
	//storageList, err := getStorageStatus(c, "local-lvm", 120)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(storageList)

	//number, err := generateRandomVMID(c)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("The number is: %d", number)
	//}
	hosts := []string{"host1", "host2", "host3"}
	fmt.Println(chooseValidHost(c, hosts))

	// todo: add this! (the hosts)
	// get the information about the hosts
	//nodeList, err := getNodeList(c)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//list := nodeList["data"].([]interface{})
	//for _, node := range list {
	//	node2 := node.(map[string]interface{})
	//	percentage := node2["mem"].(float64) / node2["maxmem"].(float64) * 100 // todo: check before diveded by zero?
	//	if node2["status"] == "online" && percentage < memoryThreshold{
	//		percantage := node2["mem"].(float64) / node2["maxmem"].(float64) * 100
	//		fmt.Printf("%s has is under the threshold: %d%%, it has memory percentage of: %f\n", node2["node"], memoryThreshold, percantage)
	//	}
	//}
	////fmt.Println(nodeList)

	// todo: add this! (the storages) - need to check the pools (right now just the 'local' and the 'local-lvm')

	//// information about all the storage
	//list := nodeList["data"].([]interface{})
	//for _, node := range list {
	//	node2 := node.(map[string]interface{})
	//	if node2["status"] != "online" {
	//		continue
	//	}
	//	nodeName := node2["id"].(string)[5:]
	//	storage, err := c.GetStorage(nodeName)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	for _, eachStorage := range storage {
	//		//fmt.Println(eachStorage.(map[string]interface{}))
	//		freeStorage := eachStorage.(map[string]interface{})["avail"].(float64) / 1000000000
	//		if freeStorage > storageThreshold {
	//			fmt.Printf("%s has %f storage available in storage: %s, less than the threshold: %d\n", nodeName, freeStorage, eachStorage.(map[string]interface{})["storage"], storageThreshold)
	//		}
	//	}
	//}

	/////

	//// get the information about the hosts
	//status, err := getStorageStatus(c, "local-lvm", 100)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(status)

	//
	//// get vm status
	//status, err := getVMInfo(c, 100)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(status["status"])

	//// stop the vm
	//err = stopVM(c, 124)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//// start the vm
	//err = startVM(c, 200)
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

	////// del	ete the vm
	//exitStatus, err := deleteVM(c, 124)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(exitStatus)

	// do load-balancing
	//exitStatus, err := migrateNode(c, 101, "host2", "qemu", "host1", false)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(exitStatus)

	// stop vm flow!

	//state, err := getVMState(c, 100)
	//if err != nil {
	//	fmt.Printf("failed getting vm state")
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("state:", state)
	////
	//err = stopVM(c, 100)
	//if err != nil {
	//	fmt.Printf("error: %v", err)
	//	return
	//}
	//fmt.Println("stopped!")
	//
	//exitStatus, err := deleteVM(c, 100)
	//if err != nil {
	//	fmt.Printf("error: %v, exis status: %s", err, exitStatus)
	//	return
	//}
	//fmt.Println("Yes!")

	//// reset flow!
	//
	//vmID := 200
	//
	//vmr := proxmox.NewVmRef(vmID)
	//
	//// check vm state after resetting it
	//state, err := c.GetVmState(vmr)
	//if err != nil {
	//	fmt.Printf("failed getting vm: %d state before reseting it: %v", vmID, err)
	//}
	//if state["status"] != "running" {
	//	// run the vm
	//	fmt.Println("vm isn't running! starting it now...")
	//	err = startVM(c, 100)
	//	if err != nil {
	//		fmt.Println("vm can't start!")
	//		return
	//	}
	//	fmt.Println("started")
	//	state2, err := c.GetVmState(vmr)
	//	if err != nil {
	//		fmt.Printf("failed getting vm: %d state before reseting it: %w", vmID, err)
	//	}
	//	if state2["running"] != nil {
	//		fmt.Println("vm isn't running!")
	//		return
	//	}
	//}
	//fmt.Println("ok, started nice, vm is running!")
	//
	//
	//exitStatus, err := c.ResetVm(vmr)
	//if err != nil {
	//	fmt.Printf("failed reset vm: %d, exis status: %s: %w", vmID, exitStatus, err)
	//}
	//fmt.Println("vm was reset!")
	//
	//
	//// check vm state after resetting it
	//state, err = c.GetVmState(vmr)
	//if err != nil {
	//	fmt.Printf("failed getting vm: %d state after reseting it: %w", vmID, err)
	//}
	//if state["status"] != "running" {
	//	fmt.Printf("vm: %d isn't rinning after resetign it. vm state: %+v", vmID, state)
	//}
	//fmt.Println("vm is running!")
	//fmt.Println("Yes!")

	//// create vm
	//err = createQemu(c, 200)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("Yes!")
	//
	//err = cloneVM(c,100, cloneVMconst)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("Yes!")
	//
	//exisStatus, err := resizeVMDisk(c,101,"scsi0",10)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(exisStatus)
	//fmt.Println("Yes!")

}

func createQemu(c *proxmox.Client, vmID int) error {
	//configTemp, err := GetConfig([]byte(createQemuConfig))
	//if err != nil {
	//	return err
	//}
	config, err := proxmox.NewConfigQemuFromJson([]byte(createQemuConfig))
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

func resetVM(c *proxmox.Client, vmID int) (string, error) {
	vmr := proxmox.NewVmRef(vmID)
	exitStatus, err := c.ResetVm(vmr)
	return exitStatus, err
}

//func getStorageList(c *proxmox.Client) (string, error) {
//	storage, err := c.GetStorageList()
//	if err != nil {
//		return "", err
//	}
//	storageList, err := json.Marshal(storage)
//	if err != nil {
//		return "", err
//	}
//	return string(storageList), nil
//}

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

func cloneVM(c *proxmox.Client, vmID int, vmParams string) error {
	sourceVmr := proxmox.NewVmRef(vmID)
	sourceVmr.SetNode("host1")
	conf, err := GetConfig(vmParams)
	if err != nil {
		return err
	}
	config, err := proxmox.NewConfigQemuFromJson(conf)
	if err != nil {
		return err
	}
	vmr := proxmox.NewVmRef(102)
	return config.CloneVm(sourceVmr, vmr, c)
	//return c.CloneQemuVm(vmr, config)
}

func migrateNode(c *proxmox.Client, vmID int, node string, vmType string, newTargetNode string, online bool) (exitStatus interface{}, err error) {
	vmr := proxmox.NewVmRef(vmID)
	vmr.SetNode(node)
	vmr.SetVmType(vmType)
	return c.MigrateNode(vmr, newTargetNode, online)
}

func getStorageList(c *proxmox.Client) (metricServers map[string]interface{}, err error) {
	return c.GetStorageList()
}

func getResourceList() {

}

func resizeVMDisk(c *proxmox.Client, vmID int, disk string, moreSizeGB int) (exitStatus interface{}, err error) {
	vmr := proxmox.NewVmRef(vmID)
	vmr.SetNode("host1")
	vmr.SetVmType("qemu")
	return c.ResizeQemuDisk(vmr, disk, moreSizeGB)
}

func getAllValidHosts(c *proxmox.Client) {
	nodeList, err := getNodeList(c)
	if err != nil {
		fmt.Println(err)
	}
	list := nodeList["data"].([]interface{})
	for _, node := range list {
		node2 := node.(map[string]interface{})
		percentage := node2["mem"].(float64) / node2["maxmem"].(float64) * 100 // todo: check before diveded by zero?
		if node2["status"] == "online" && percentage < memoryThreshold {
			percantage := node2["mem"].(float64) / node2["maxmem"].(float64) * 100
			fmt.Printf("%s has is under the threshold: %d%%, it has memory percentage of: %f\n", node2["node"], memoryThreshold, percantage)
		}
	}
	//fmt.Println(nodeList)
}

func getAllValidStorages(c *proxmox.Client) {
	nodeList, err := getNodeList(c)
	if err != nil {
		fmt.Println(err)
	}
	// information about all the storage
	list := nodeList["data"].([]interface{})
	for _, node := range list {
		node2 := node.(map[string]interface{})
		if node2["status"] != "online" {
			continue
		}
		nodeName := node2["id"].(string)[5:]
		storage, err := c.GetStorage(nodeName)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, eachStorage := range storage {
			//fmt.Println(eachStorage.(map[string]interface{}))
			freeStorage := eachStorage.(map[string]interface{})["avail"].(float64) / 1000000000
			if freeStorage > storageThreshold {
				fmt.Printf("%s has %f storage available in storage: %s, less than the threshold: %d\n", nodeName, freeStorage, eachStorage.(map[string]interface{})["storage"], storageThreshold)
			}
		}
	}

}

func generateRandomVMID(c *proxmox.Client) (int, error) {
	for i := 0; i < 5; i++ {
		// Seed the random number generator with the current time
		// in nanoseconds. This ensures that the random numbers
		// generated will be different each time the program is run.
		rand.Seed(time.Now().UnixNano())

		// Generate a random integer between 200 and 10,000 ([,))
		min := 200
		max := 10000
		randomNumber := rand.Intn(max-min+1) + min
		ifExists, err := c.VMIdExists(randomNumber)
		if err != nil {
			return -1, err
		}
		if !ifExists {
			return randomNumber, nil
		}
	}
	return 0, fmt.Errorf("no valid VMID number was found")
}

func chooseValidHost(c *proxmox.Client, hosts []string) string {
	// Seed the random number generator with the current time
	// in nanoseconds. This ensures that the random numbers
	// generated will be different each time the program is run.
	rand.Seed(time.Now().UnixNano())

	i := rand.Intn(len(hosts))
	return hosts[i]
}

// cluster/status
