package main

import (
	"fmt"
	"github.com/quicky-dev/generator/generator"
)

func main() {
	generator.Init("output", false)
	script, _ := generator.GenerateGeneric()
	fmt.Println(script)
}
