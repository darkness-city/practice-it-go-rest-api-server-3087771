package main

import (
	"example.com/helloworld"
)

func main() {
	a := helloworld.App{}
	a.Port = ":9003"
	a.Initialize()
	a.Run()
	// helloworld.Run(":9003")
}
