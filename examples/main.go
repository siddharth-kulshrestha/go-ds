package main

import (
	"fmt"
)

var exampleRunners []Runner

func init() {
	exampleRunners = []Runner{}
}

type Runner struct {
	Name    string
	RunFunc func()
}

func RegisterRunner(runF func(), name string) {
	exampleRunners = append(exampleRunners, Runner{name, runF})
}

func main() {
	for _, runner := range exampleRunners {
		fmt.Printf("======================= %s =====================================\n", runner.Name)
		runner.RunFunc()
		fmt.Printf("===========================================================================\n\n")
	}
}
