package generator

import (
    "os"

    "github.com/quicky-dev/generator/macos"
)


// Generic is the generic data structure for generating 
type Generic struct {
    Editor string
    Python bool
    Node bool
    NVM bool
    PATH string
}

// Init will initialize the modules filepath for where to save the generated files
func Init(filePath string) bool {
   return true
}

// GenerateGeneric will generate a generic developers setup for the user to
// to run locally
func GenerateGeneric() {
    genericScript := []string{}

    // Add shebang to the top of the file
    genericScript = append(genericScript, "#! /bin/bash")
    genericScript = macos.InstallXCode(genericScript)

    f, err := os.Create("test")

    if err != nil {
        panic("error")
    }

    for _, command := range genericScript {
        _, err := f.WriteString(command + "\n")

        if err != nil {
            f.Close()
            panic("error")
        }
    }
    f.Close()
}


