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

func main() {
	out, err := exec.Command("sudo", "lshw", "-json").Output()
	if err != nil {
		return
	}

	data := HWINFO{}

	err = json.Unmarshal([]byte(out), &data)
	if err != nil {
		return
	}

	//  useful website , find you want information JSON PATH and print
	//	JSON to Struct:			https://jsonpathfinder.com/
	// 	JSON PATH finder:		https://transform.tools/json-to-go
	//  Author : Andy Hu
	//  Date: 2023 / 2 / 15 07:59
	//  email:327656021@qq.com

	// some test case , use JSON path print
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
