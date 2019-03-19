package generator

import (
    "os"
    "path"
    "strings"
    "errors"

    "github.com/quicky-dev/generator/macos"
    "github.com/google/uuid"
)

var filePath string;
var debug bool;

// Helper function for inserting commands into the overall script slice.
// Very useful for not having to call append manually every time
func commander(script *[]string) func(string, int) {
    return func (command string, indent int) {
        indents := strings.Repeat("\t", indent)
        *script = append(*script, indents + command) 
    }
}


// Init will initialize the modules filepath for where to save the generated files
func Init(path string, debugMode bool) bool {
   filePath = path
   debug = debugMode
   return true
}

// GenerateGeneric will generate a generic developers setup for the user to
// to run locally
func GenerateGeneric() (string, error) {
    if filePath == "" {
        return "", errors.New("The current file path isnt set")
    }
    genericScript := []string{}

    // Add shebang to the top of the file to ensure that bash
    // executes the file
    genericScript = append(genericScript, "#! /bin/bash\n")

    // Setup script
    macos.InstallXCode(commander(&genericScript))
    macos.InstallBrew(commander(&genericScript))

     
    fileName := uuid.New().String()
    file := path.Join(filePath, fileName)
    f, err := os.Create(file)

    if err != nil {
        return "", err
    }

    for _, command := range genericScript {
        _, err := f.WriteString(command + "\n")

        if err != nil {
            f.Close()
            return "", err
        }
    }
    f.Close()
    return file, nil
}


