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
	var UpdatesStruct_var UpdatesStruct
	var OsChecker_var = OsChecker()

	if OsChecker_var == "centos" {
		UpdatesStruct_var = Centos()
	} else if OsChecker_var == "ubuntu" {
		UpdatesStruct_var = UbuntuDebian()
	}

	var final_output_all_updates_int = UpdatesStruct_var.AllUpdates
	var final_output_all_updates_string = strconv.Itoa(UpdatesStruct_var.AllUpdates)
	var final_output_security_updates_int = UpdatesStruct_var.SecurityUpdates
	var final_output_security_updates_string = strconv.Itoa(UpdatesStruct_var.SecurityUpdates)

	if final_output_all_updates_int > 1 && final_output_security_updates_int > 1 {
		fmt.Println(" 🟡 There are " + final_output_all_updates_string + " updates available.")
		fmt.Println(" 🔴 Including " + final_output_security_updates_string + " security updates!")
	} else if final_output_all_updates_int > 1 {
		fmt.Println(" 🟡 There are " + final_output_all_updates_string + " updates available.")
	} else if final_output_security_updates_int > 1 {
		fmt.Println(" 🔴 There are " + final_output_security_updates_string + " security updates waiting to be installed!")
	} else if final_output_all_updates_int == 1 && final_output_security_updates_int == 1 {
		fmt.Println(" 🔴 There is " + final_output_all_updates_string + " security update waiting to be installed!")
	} else {
		fmt.Println(" 🟢 The system is up to date. Great work!")
	}
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
	} else if final_output == "ID=ubuntu" || final_output == "ID=pop" {
		final_output = "ubuntu"
	} else {
		log.Fatal(1, "Sorry, but your OS is not supported!")
	}

	return final_output
}
