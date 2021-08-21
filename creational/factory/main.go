package main

import (
	"design_patterns/utils"
	"fmt"
	"os"
	"strings"
)

var (
	stateOpen = "open"
	stateClosed = "closed"
)

type PopUpper interface {
	Open()
	Close()
}

// --------------------------- WindowsPopUp --------------------

type WindowsPopUp struct {
	state string
	name string
}

func (w WindowsPopUp) Open() {
	fmt.Println("windows pop up is open!")
}

func (w WindowsPopUp) Close() {
	fmt.Println("we are sorry to say that windows pop up is closed")
}

// --------------------------- LinuxPopUp --------------------

type LinuxPopUp struct {
	noCores int
	name string
	state string
}

func (w LinuxPopUp) Open() {
	fmt.Println("linux pop up is open!")
}

func (w LinuxPopUp) Close() {
	fmt.Println("we are sorry to say that linux pop up is closed")
}

// --------------------------- PopUp ------------------------------

type PopUpManager struct {
	popUp PopUpper
}

func (p *PopUpManager) Open() {
	systemType := os.Getenv("OS-TEMPORARY")
	switch strings.ToLower(systemType) {
	case "windows":
		p.popUp = WindowsPopUp{
			state: stateClosed,
			name:  "windows pop up",
		}
	default:
		p.popUp = LinuxPopUp{
			noCores: 8,
			name:    "linux pop up",
			state:   stateClosed,
		}
	}
	p.popUp.Open()
}

func (p *PopUpManager) Close() {
	p.popUp.Close()
}

func main() {
	popUpManager := PopUpManager{}

	utils.IfErrorPanic(os.Setenv("OS-TEMPORARY", "windows"))
	popUpManager.Open()
	popUpManager.Close()

	utils.IfErrorPanic(os.Setenv("OS-TEMPORARY", "linux"))
	popUpManager.Open()
	popUpManager.Close()
}