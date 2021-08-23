package main

import (
	"design_patterns/utils"
	"fmt"
	"go.uber.org/dig"
)

type Bottomer interface {
	Bottom()
}

type BottomLevel struct {
	Something string
}

func(b BottomLevel) Bottom() {
	fmt.Println("hello from bottom level")
}

func NewBottomLevel() Bottomer {
	return BottomLevel{
		Something: "bottom level test",
	}
}

// ------------------------- another one ----------------------

type Middler interface {
	Middle()
}

type MiddleLevel struct {
	BottomLevel Bottomer
}

func (m MiddleLevel) Middle() {
	fmt.Println("> hi from middle level")
	m.BottomLevel.Bottom()
}

func NewMiddleLevel(b Bottomer) Middler {
	return MiddleLevel{
		BottomLevel: b,
	}
}

// ------------------------- another one ----------------------

type Config struct {
	A string
	B int
}

func NewConfig() Config {
	return Config{
		A: "aaa",
		B: 10,
	}
}

type TopLevel struct {
	Middle Middler
	Config Config
}

func (m TopLevel) Top() {
	fmt.Println(">> salem from top level")
	m.Middle.Middle()
}

func NewTopLevel(m Middler, c Config) TopLevel {
	return TopLevel{
		Middle: m,
		Config: c,
	}
}

// ------------------- main part -----------------

func main() {

	config := NewConfig()

	bottom := NewBottomLevel()
	middle := NewMiddleLevel(bottom)
	top := NewTopLevel(middle, config)

	top.Top()

	fmt.Println()

	// -------------------- container ---------------
	container := dig.New()

	utils.IfErrorPanic(container.Provide(NewBottomLevel))
	utils.IfErrorPanic(container.Provide(NewMiddleLevel))
	utils.IfErrorPanic(container.Provide(NewTopLevel))
	utils.IfErrorPanic(container.Provide(NewConfig))

	utils.IfErrorPanic(
		container.Invoke(func(top TopLevel) {
			top.Top()
		}),
	)
}