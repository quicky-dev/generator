package main

import (
	"fmt"
	"github.com/quicky-dev/generator/generator"
)

func main() {
	generator.Init("output", false)

	macos := generator.GetMacOSGenerator()
	script, _ := macos.GenerateGenericScript()
	fmt.Println(script)
}
