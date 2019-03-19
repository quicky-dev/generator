package main

import ( 
    "github.com/quicky-dev/generator/generator"
)

func main() {
    generator.Init("output", true)
    generator.GenerateGeneric()
}
