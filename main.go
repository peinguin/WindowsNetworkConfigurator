package main

import (
    "github.com/lxn/walk"
    . "github.com/lxn/walk/declarative"
    "os/exec"
)

func main() {
    var outTE *walk.TextEdit
    MainWindow{
        Title:   "Network Configurator",
        MinSize: Size{640, 480},
        Layout:  VBox{},
        Children: []Widget{
            PushButton{
		Text: "Rezerv",
                OnClicked: func() {
		    out, err := exec.Command(
		        "netsh",
			"interface",
			"ip",
			"set",
			"address",
			"name=\"Ethernet\"",
			"source=static",
			"addr=192.168.48.37",
			"mask=255.255.255.0",
			"gateway=192.168.48.100").Output()
		    if err != nil {
		    	outTE.SetText(err.Error())
		    }else{
		        outTE.SetText(string(out))
		    }
                },
            },
            PushButton{
                Text: "VPN podklychenie",
                OnClicked: func() {
		    out, err := exec.Command(
		        "netsh",
			"interface",
			"ip",
			"set",
			"address",
			"name=\"Ethernet\"",
			"source=static",
			"addr=192.168.48.37",
			"mask=255.255.255.0",
			"gateway=192.168.48.4",
		    ).Output()
		    if err != nil {
		    	outTE.SetText(err.Error())
		    }else{
			outTE.SetText(string(out))
		    }
                },
            },
            PushButton{
                Text: "STANDART",
                OnClicked: func() {
		    out, err := exec.Command(
			"C:\\\"Program Files (x86)\"\\\"Cisco Systems\"\\\"VPN Client\"\\\vpngui.exe",
			"-c",
			"-user",
			"YLazorenko",
			"-pwd",
			"Kfpjhtyrj093").Output()
		    if err != nil {
		    	outTE.SetText(err.Error())
		    }else{
			outTE.SetText(string(out))
		    }
                },
            },
	    TextEdit{AssignTo: &outTE, ReadOnly: true},
        },
    }.Run()
}