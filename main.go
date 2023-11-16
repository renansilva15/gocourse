package main

import (
	"fmt"
	"gocourse/l1packages"
	"gocourse/l2variables"
)

func main() {
	var (
		lesson int
	)

	packages := []func(){l1packages.Main, l2variables.Main}

	fmt.Println("Enter the number of the lesson you want to see:")
	fmt.Scanln(&lesson)

	packages[lesson-1]()
}
