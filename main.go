package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type HWINFO []struct {
	ID            string `json:"id"`
	Class         string `json:"class"`
	Claimed       bool   `json:"claimed"`
	Handle        string `json:"handle"`
	Description   string `json:"description"`
	Product       string `json:"product"`
	Vendor        string `json:"vendor"`
	Version       string `json:"version"`
	Serial        string `json:"serial"`
	Width         int    `json:"width"`
	Configuration struct {
		Boot    string `json:"boot"`
		Chassis string `json:"chassis"`
		Family  string `json:"family"`
		Sku     string `json:"sku"`
		UUID    string `json:"uuid"`
	} `json:"configuration"`
	Capabilities struct {
		Smbios341  string `json:"smbios-3.4.1"`
		Dmi341     string `json:"dmi-3.4.1"`
		Smp        string `json:"smp"`
		Vsyscall32 string `json:"vsyscall32"`
	} `json:"capabilities"`
	Children []struct {
		ID          string `json:"id"`
		Class       string `json:"class"`
		Claimed     bool   `json:"claimed,omitempty"`
		Handle      string `json:"handle,omitempty"`
		Description string `json:"description"`
		Product     string `json:"product"`
		Vendor      string `json:"vendor"`
		Physid      string `json:"physid"`
		Version     string `json:"version"`
		Serial      string `json:"serial"`
		Slot        string `json:"slot,omitempty"`
		Children    []struct {
			ID           string `json:"id"`
			Class        string `json:"class"`
			Claimed      bool   `json:"claimed"`
			Description  string `json:"description,omitempty"`
			Vendor       string `json:"vendor,omitempty"`
			Physid       string `json:"physid"`
			Version      string `json:"version,omitempty"`
			Date         string `json:"date,omitempty"`
			Units        string `json:"units,omitempty"`
			Size         int    `json:"size,omitempty"`
			Capacity     int    `json:"capacity,omitempty"`
			Capabilities struct {
				Pci                   string `json:"pci"`
				Upgrade               string `json:"upgrade"`
				Shadowing             string `json:"shadowing"`
				Cdboot                string `json:"cdboot"`
				Bootselect            string `json:"bootselect"`
				Socketedrom           string `json:"socketedrom"`
				Edd                   string `json:"edd"`
				Int13Floppynec        string `json:"int13floppynec"`
				Int13Floppytoshiba    string `json:"int13floppytoshiba"`
				Int13Floppy360        string `json:"int13floppy360"`
				Int13Floppy1200       string `json:"int13floppy1200"`
				Int13Floppy720        string `json:"int13floppy720"`
				Int13Floppy2880       string `json:"int13floppy2880"`
				Int5Printscreen       string `json:"int5printscreen"`
				Int9Keyboard          string `json:"int9keyboard"`
				Int14Serial           string `json:"int14serial"`
				Int17Printer          string `json:"int17printer"`
				Int10Video            string `json:"int10video"`
				Usb                   string `json:"usb"`
				Biosbootspecification string `json:"biosbootspecification"`
				Uefi                  string `json:"uefi"`
			} `json:"capabilities,omitempty"`
			Handle   string `json:"handle,omitempty"`
			Slot     string `json:"slot,omitempty"`
			Children []struct {
				ID          string `json:"id"`
				Class       string `json:"class"`
				Claimed     bool   `json:"claimed"`
				Handle      string `json:"handle"`
				Description string `json:"description"`
				Product     string `json:"product"`
				Vendor      string `json:"vendor"`
				Physid      string `json:"physid"`
				Serial      string `json:"serial"`
				Slot        string `json:"slot"`
				Units       string `json:"units"`
				Size        int64  `json:"size"`
				Width       int    `json:"width"`
				Clock       int64  `json:"clock"`
			} `json:"children,omitempty"`
			Configuration struct {
				Level string `json:"level"`
			} `json:"configuration,omitempty"`
			Product     string `json:"product,omitempty"`
			Businfo     string `json:"businfo,omitempty"`
			Serial      string `json:"serial,omitempty"`
			Width       int    `json:"width,omitempty"`
			Clock       int    `json:"clock,omitempty"`
			Logicalname string `json:"logicalname,omitempty"`
		} `json:"children,omitempty"`
		Units    string `json:"units,omitempty"`
		Capacity int    `json:"capacity,omitempty"`
	} `json:"children"`
}

type LSHW struct {
	Businfo string `json:"businfo,omitempty"`
	// Capabilities  *LSHWCAP    `json:"capabilities,omitempty"`
	Capacity int64   `json:"capacity,omitempty"`
	Children []*LSHW `json:"children,omitempty"`
	Claimed  bool    `json:"claimed,omitempty"`
	Class    string  `json:"class,omitempty"`
	Clock    int     `json:"clock,omitempty"`
	// Configuration *LSHWCONFIG `json:"configuration,omitempty"`
	Date        string      `json:"date,omitempty"`
	Description string      `json:"description,omitempty"`
	Dev         string      `json:"dev,omitempty"`
	Disabled    bool        `json:"disabled,omitempty"`
	Handle      string      `json:"handle,omitempty"`
	ID          string      `json:"id,omitempty"`
	Logicalname interface{} `json:"logicalname,omitempty"`
	Physid      string      `json:"physid,omitempty"`
	Product     string      `json:"product,omitempty"`
	Serial      string      `json:"serial,omitempty"`
	Slot        string      `json:"slot,omitempty"`
	Size        int64       `json:"size,omitempty"`
	Units       string      `json:"units,omitempty"`
	Vendor      string      `json:"vendor,omitempty"`
	Version     string      `json:"version,omitempty"`
	Width       int         `json:"width,omitempty"`
}

type address struct {
	Street  string `json:"street"`          // 街道
	Ste     string `json:"suite,omitempty"` // 单元（可以不存在）
	City    string `json:"city"`            // 城市
	State   string `json:"state"`           // 州/省
	Zipcode string `json:"zipcode"`         // 邮编
}

func main() {
	out, err := exec.Command("sudo", "lshw", "-json").Output()
	if err != nil {
		return
	}

	// fmt.Println(string(out))

	data := HWINFO{}

	err = json.Unmarshal([]byte(out), &data)
	if err != nil {
		return
	}
	fmt.Printf("%T\n", data)
	fmt.Println(data[0].ID)
	fmt.Println(data[0].Product)
	fmt.Println(data[0].Configuration.Boot)
	fmt.Println(data[0].Children[0].Children[0].Vendor)
	fmt.Println(data[0].Children[0].Children[1].Children[0].Vendor)
	fmt.Println(data[0].Children[0].Product)
	fmt.Println(data[0].Children[0].Children[0].ID)
	fmt.Println(data[0].Children[0].Children[1].Class)
}
