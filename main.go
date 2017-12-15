package main

import (
	"fmt"
	"os"
)

func main() {

	e := NewEngine()
	e.Load(os.Args[1:]...)
	e.Start()

	fmt.Println("Pass")

}
