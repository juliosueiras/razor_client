package node

import (
	"github.com/dghubble/sling"
	"github.com/juliosueiras/razor_client/api/error"
)

type NodeService struct {
	Client *sling.Sling
}

func (r NodeService) ListNodes() (*Nodes, error) {
	nodes := new(Nodes)
	_, err := r.Client.Get("/api/collections/nodes").ReceiveSuccess(nodes)

	return nodes, err
}

func (r NodeService) NodeDetails(id string) (*Node, *errorMsg.ErrorMessage) {
	node := new(Node)
	resErr := new(errorMsg.ErrorMessage)
	r.Client.Get("/api/collections/nodes/"+id).Receive(node, resErr)

	return node, resErr
}

type NodeItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Spec string `json:"spec"`
}

type Nodes struct {
	Items []NodeItem `json:"items"`
	Spec  string     `json:"spec"`
	Total string     `json:"total"`
}

type Node struct {
	State        State    `json:"state"`
	Hostname     string   `json:"hostname"`
	Policy       Policy   `json:"policy"`
	ID           string   `json:"id"`
	Tags         []Policy `json:"tags"`
	RootPassword string   `json:"root_password"`
	Spec         string   `json:"spec"`
	Log          Log      `json:"log"`
	LastCheckin  string   `json:"last_checkin"`
	Name         string   `json:"name"`
	DHCPMAC      string   `json:"dhcp_mac"`
	HwInfo       HwInfo   `json:"hw_info"`
	Facts        Facts    `json:"facts"`
}

type Facts struct {
	MemoryfreeMB           string       `json:"memoryfree_mb"`
	MemorysizeMB           string       `json:"memorysize_mb"`
	Physicalprocessorcount int64        `json:"physicalprocessorcount"`
	Hardwaremodel          string       `json:"hardwaremodel"`
	IsVirtual              bool         `json:"is_virtual"`
	BIOSVendor             string       `json:"bios_vendor"`
	Manufacturer           string       `json:"manufacturer"`
	MTULo                  int64        `json:"mtu_lo"`
	Ipaddress              string       `json:"ipaddress"`
	Type                   string       `json:"type"`
	DHCPServers            DHCPServers  `json:"dhcp_servers"`
	Interfaces             string       `json:"interfaces"`
	Virtual                string       `json:"virtual"`
	Hardwareisa            string       `json:"hardwareisa"`
	Productname            string       `json:"productname"`
	Architecture           string       `json:"architecture"`
	Facterversion          string       `json:"facterversion"`
	Netmask                string       `json:"netmask"`
	Macaddress             string       `json:"macaddress"`
	Processors             Processors   `json:"processors"`
	Rubyplatform           string       `json:"rubyplatform"`
	Gid                    string       `json:"gid"`
	BIOSReleaseDate        string       `json:"bios_release_date"`
	Serialnumber           string       `json:"serialnumber"`
	Boardserialnumber      string       `json:"boardserialnumber"`
	Uniqueid               string       `json:"uniqueid"`
	Partitions             []Partitions `json:"partitions"`
	OS                     OS           `json:"os"`
	Boardproductname       string       `json:"boardproductname"`
	Blockdevices           string       `json:"blockdevices"`
	BIOSVersion            string       `json:"bios_version"`
	Processorcount         int64        `json:"processorcount"`
	SystemUptime           SystemUptime `json:"system_uptime"`
	Boardmanufacturer      string       `json:"boardmanufacturer"`
	UUID                   string       `json:"uuid"`
}

type DHCPServers struct {
	System string `json:"system"`
}

type OS struct {
	Name    string  `json:"name"`
	Release Release `json:"release"`
	Family  string  `json:"family"`
}

type Release struct {
	Minor string `json:"minor"`
	Major string `json:"major"`
	Full  string `json:"full"`
}

type Partitions struct {
	Filesystem string `json:"filesystem"`
	Size       string `json:"size"`
	Label      string `json:"label"`
}

type Processors struct {
	Physicalcount int64    `json:"physicalcount"`
	Models        []string `json:"models"`
	Count         int64    `json:"count"`
}

type SystemUptime struct {
	Uptime  string `json:"uptime"`
	Days    int64  `json:"days"`
	Hours   int64  `json:"hours"`
	Seconds int64  `json:"seconds"`
}

type HwInfo struct {
	MAC []string `json:"mac"`
}

type Log struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Params Params `json:"params"`
}

type Params struct {
	Limit Limit `json:"limit"`
	Start Limit `json:"start"`
}

type Limit struct {
	Type string `json:"type"`
}

type Policy struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Spec string `json:"spec"`
}

type State struct {
	Stage     string `json:"stage"`
	Installed bool   `json:"installed"`
}
