package main

import (
	"fmt"
	"os"
)

func main() {

	e := NewEngine()
	err := e.Load(os.Args[1])
	if err != nil {
		panic(err)
	}
	errs := e.Start()
	for _, err := range errs {
		fmt.Println("Start Error:", err)
	}
	if len(errs) == 0 {
		fmt.Println("Pass")
	}
}
