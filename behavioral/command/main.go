package main

import "fmt"

type TvSwitcher interface {
	Execute()
}

type TvSwitchState struct {
	IsOn bool
}

func (t *TvSwitchState) Execute() {
	if t.IsOn {
		fmt.Println("switching tv off")
	} else {
		fmt.Println("switching tv on")
	}
	t.IsOn = !t.IsOn
}

type RemoteController struct {
	RedButton TvSwitcher
}

type TvBuiltInSwitchButton struct {
	RedButton TvSwitcher
}

func main()  {
	tvSwitch := TvSwitchState{
		IsOn: false,
	}

	remoteController := RemoteController{
		RedButton: &tvSwitch,
	}

	tvController := TvBuiltInSwitchButton{
		RedButton: &tvSwitch,
	}

	remoteController.RedButton.Execute()
	tvController.RedButton.Execute()
}
