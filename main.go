package main

import (
	"fmt"
	"github.com/quicky-dev/generator/generator"
)

func main() {
	generator.Init("output", false)

	macos := generator.GetMacOSGenerator()
	ubuntu := generator.GetUbuntuGenerator()
	script, _ := macos.GenerateGenericScript()
	ubuntu_script, _ := ubuntu.GenerateGenericScript()
	fmt.Println(script)
	fmt.Println(ubuntu_script)
}
