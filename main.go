package main

import (
	"os"
)

func main() {

	e := NewEngine(os.Args[1])
	e.Load(os.Args[2:]...)
	e.Start()
}
