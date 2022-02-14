package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type UpdatesStruct struct {
	AllUpdates      int
	SecurityUpdates int
}

func main() {
	OsChecker()
	// var UpdatesStruct_var = UbuntuDebian()
	var UpdatesStruct_var = Centos()
	fmt.Println(" 🟢 There are " + strconv.Itoa(UpdatesStruct_var.AllUpdates) + " updates available.")
	fmt.Println(" 🔴 Including " + strconv.Itoa(UpdatesStruct_var.SecurityUpdates) + " security updates!")
}

func UbuntuDebian() UpdatesStruct {
	//Set vars
	var UpdatesStruct_var UpdatesStruct

	//All updates list
	refresh_cmd := "sudo apt-get update"
	var _, _ = exec.Command("bash", "-c", refresh_cmd).Output()

	all_updates_cmd := "sudo apt-get dist-upgrade -s | grep Inst"
	var all_updates_out, _ = exec.Command("bash", "-c", all_updates_cmd).Output()
	all_updates_output := strings.Split(string(all_updates_out), "\n")

	var all_updates_list []string

	for _, item := range all_updates_output {
		if item != "" {
			all_updates_list = append(all_updates_list, item)
		}
	}

	UpdatesStruct_var.AllUpdates = len(all_updates_list)

	//Sec updates list
	security_updates_cmd := "sudo apt-get dist-upgrade -s | grep Inst | grep security"
	var security_updates_out, _ = exec.Command("bash", "-c", security_updates_cmd).Output()

	security_updates_output := strings.Split(string(security_updates_out), "\n")

	var security_updates_list []string

	for _, item := range security_updates_output {
		if item != "" {
			security_updates_list = append(security_updates_list, item)
		}
	}

	UpdatesStruct_var.SecurityUpdates = len(security_updates_list)

	return UpdatesStruct_var
}

func Centos() UpdatesStruct {
	//Set vars
	var UpdatesStruct_var UpdatesStruct

	//All updates list
	refresh_cmd := "sudo yum makecache fast"
	var _, _ = exec.Command("bash", "-c", refresh_cmd).Output()

	all_updates_cmd := "sudo yum --cacheonly check-update | grep -v \"Loaded plugins: fastestmirror\" | grep -vG \"^$\" | grep -v \"updateinfo info done\""
	var all_updates_out, _ = exec.Command("bash", "-c", all_updates_cmd).Output()
	all_updates_output := strings.Split(string(all_updates_out), "\n")

	var all_updates_list []string

	for _, item := range all_updates_output {
		if item != "" {
			all_updates_list = append(all_updates_list, item)
		}
	}
	UpdatesStruct_var.AllUpdates = len(all_updates_list)

	//Sec updates list
	security_updates_cmd := "sudo yum --cacheonly updateinfo info security | grep -v \"Loaded plugins: fastestmirror\" | grep -vG \"^$\" | grep -v \"updateinfo info done\""
	var security_updates_out, _ = exec.Command("bash", "-c", security_updates_cmd).Output()

	security_updates_output := strings.Split(string(security_updates_out), "\n")

	var security_updates_list []string

	for _, item := range security_updates_output {
		if item != "" {
			security_updates_list = append(security_updates_list, item)
		}
	}

	UpdatesStruct_var.SecurityUpdates = len(security_updates_list)

	return UpdatesStruct_var
}

func OsChecker() string {
	cmd := "cat /etc/os-release | grep \"ID=\" | grep -v \"VERSION_ID=\""
	var os_release, _ = exec.Command("bash", "-c", cmd).Output()
	final_output := string(os_release)
	final_output = strings.ReplaceAll(final_output, "\n", "")

	if final_output == "ID=\"centos\"" {
		final_output = "centos"
	} else if final_output == "ID=ubuntu" {
		final_output = "ubuntu"
	} else {
		log.Fatal(1, "Sorry, but your OS is not supported!")
	}

	fmt.Println(final_output)
	return final_output
}
