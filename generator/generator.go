package generator

import (
    "os"
    "path"
    "strings"
    "errors"

    "github.com/quicky-dev/generator/macos"
    "github.com/google/uuid"
    "github.com/mitchellh/mapstructure"
)

var filePath string;
var debug bool;

type category struct{
    Description string
    Items []string
}

// InstallRequest is used for unmarshalling json objects directly into a request object
// that the script factory can then digest and work with
type InstallRequest struct{
    Languages category `json:"Languages"`
    Terminals category `json:"Terminals"`
    Shells    category `json:"Shells"`
    Browsers  category `json:"Browsers"`
    Editors   category `json:"Editors"`
    Tools     category `json:"Tools"`
    Databases category `json:"Databases"`
}


func supportedMacPkgs() InstallRequest {
    pkgCategories := []string{
        "Languages",
        "Terminals",
        "Shells",
        "Browsers",
        "Editors",
        "Tools",
        "Databases",
    }

    pkgDesc := []string{
        "Select all programming languages of your choice",
        "Select all Terminal emulators of your choice",
        "Select all terminal shells of your choice",
        "Select all web browsers of your choice",
        "Select all text editors of your choice",
        "Select all developer tools of your choice",
        "Select all databases of your choice",
    }
    
    pkgItems := []map[string]string{
        macos.SupportedLangs,
        macos.SupportedTerminals,
        macos.SupportedShells,
        macos.SupportedBrowsers,
        macos.SupportedEditors,
        macos.SupportedTools,
        macos.SupportedDatabases,
    }

    macPkgMap := map[string]category{}
    
    for index, currCategory := range pkgCategories {
        var currentItems []string

        for k := range pkgItems[index] {
            currentItems = append(currentItems, k)
        }

        categoryStruct := category{
            Description: pkgDesc[index],
            Items: currentItems,
        }

        macPkgMap[currCategory] = categoryStruct 
    }

    macPkgs := InstallRequest{}
    mapstructure.Decode(macPkgMap, &macPkgs)
    return macPkgs
}

// MacPkgs is a list of All available packages
var MacPkgs = supportedMacPkgs()


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

    genericPkgs := map[string]category{
        "Languages":{
            Description:"Select all programming languages of your choice",
            Items: []string{"python", "ruby", "yolo", "java"},
        },
        "Terminals":{
            Description: "testsjk",
            Items: []string{},
        },
        "Shells":{
            Description: "Select all Terminal shells of your choice",
            Items: []string{"hyper", "fake-terminal", "itemr2", "yeet"},
        },
        "Browsers":{
            Description: "Select the web browser of your choice",
            Items: []string{"google chrome", "opera"},
        },
        "Editors":{
            Description:"Select the text editor of your choice",
            Items: []string{"vim", "macvim", "sublime-text"},
        },
    }

    var install = InstallRequest{} 
    mapstructure.Decode(genericPkgs, &install)

    script := []string{}

    // Add shebang to the top of the file to ensure that bash
    // executes the file
    script = append(script, "#! /bin/bash\n")

    // Setup script
    macos.InstallXCode(commander(&script))
    macos.InstallBrew(commander(&script))
    macos.InstallLangs(commander(&script), install.Languages.Items)
    macos.InstallTerminals(commander(&script), install.Terminals.Items)
    macos.InstallShells(commander(&script), install.Shells.Items)
    macos.InstallBrowsers(commander(&script), install.Browsers.Items)
    macos.InstallEditors(commander(&script), install.Editors.Items)

    // Generate a new uuid4
    uuid, err := uuid.NewRandom(); if err != nil {
        return "", err
    }

    // Convert the uuid4 to a string
    fileName := uuid.String()
    file := path.Join(filePath, fileName)

    // Attempt to create a file
    f, err := os.Create(file); if err != nil {
        return "", err
    }

    // Iterate over the script and start writing it to a file 
    for _, command := range script {
        _, err := f.WriteString(command + "\n"); if err != nil {
            f.Close()
            return "", err
        }
    }
    f.Close()
    return file, nil
}


