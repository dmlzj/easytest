package main

import (
	"os"
)

func main() {

	e := NewEngine()
	e.Load(os.Args[1:]...)
	e.Start()
}
