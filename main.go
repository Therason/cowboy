package main

import (
	"github.com/therason/cowboy/core"
)

func main() {
	core.App = &core.Cowboy{}
	if err := core.App.Init(); err != nil {
		panic(err)
	}
}
