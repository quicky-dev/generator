package generator

import (
    "strings"
    "errors"

    "github.com/quicky-dev/generator/macos"
    "github.com/quicky-dev/generator/util"
    "github.com/mitchellh/mapstructure"
)

var filePath string;
var debug bool;

// InstallRequest is used for unmarshalling json objects directly from a user request into an install request
type InstallRequest struct{
    Languages []string `json:"Languages"`
    Terminals []string `json:"Terminals"`
    Shells    []string `json:"Shells"`
    Browsers  []string `json:"Browsers"`
    Editors   []string `json:"Editors"`
    Tools     []string `json:"Tools"`
    Databases []string `json:"Databases"`
}

type category struct{
    Description string `json:"Description"`
    Items []string `json:"Items"`
}

// SupportedPackages is a struct for maintaining supported packages within our factory
type SupportedPackages struct{
    Languages category `json:"Languages"`
    Terminals category `json:"Terminals"`
    Shells    category `json:"Shells"`
    Browsers  category `json:"Browsers"`
    Editors   category `json:"Editors"`
    Tools     category `json:"Tools"`
    Databases category `json:"Databases"`
}

// supportedmacPkgs gets us the list of supported packages
func getSupportedMacPkgs() SupportedPackages {
    // All the available categories we have
    pkgCategories := []string{
        "Languages",
        "Terminals",
        "Shells",
        "Browsers",
        "Editors",
        "Tools",
        "Databases",
    }

    // Descriptions for all the categories we have so far
    pkgDesc := []string{
        "Select all programming languages of your choice",
        "Select all Terminal emulators of your choice",
        "Select all terminal shells of your choice",
        "Select all web browsers of your choice",
        "Select all text editors of your choice",
        "Select all developer tools of your choice",
        "Select all databases of your choice",
    }
   
    // All the available packages we have for the items so far
    pkgItems := []map[string]string{
        macos.SupportedLangs,
        macos.SupportedTerminals,
        macos.SupportedShells,
        macos.SupportedBrowsers,
        macos.SupportedEditors,
        macos.SupportedTools,
        macos.SupportedDatabases,
    }

    // the pkg map to map strings (category names) to category structs
    macPkgMap := map[string]category{}
    
    // Iterate over all the categories
    for index, currCategoryName := range pkgCategories {
        var currentItems []string

        // Append all keys (package names) into the current items string array
        for k := range pkgItems[index] {
            currentItems = append(currentItems, k)
        }

        // Create the category struct which will keep the current packages description
        // and it's Items bundled together
        categoryStruct := category{
            Description: pkgDesc[index],
            Items: currentItems,
        }

        // Map the current category name to the category struct
        macPkgMap[currCategoryName] = categoryStruct 
    }

    // Create the install request and bind it into the mapStructure
    macPkgs := SupportedPackages{}
    mapstructure.Decode(macPkgMap, &macPkgs)
    return macPkgs
}

// List of all the generic macpkgs provided
var genericMacPkgs = map[string]category{
        "Languages":{
            Description:"Select all programming languages of your choice",
            Items: []string{"python", "ruby", "java"},
        },
        "Terminals":{
            Description: "Select all terminal emulators of your choice",
            Items: []string{},
        },
        "Shells":{
            Description: "Select all Terminal shells of your choice",
            Items: []string{"hyper", "iterm2", "yeet"},
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

// MacPkgs is a list of All available packages
var MacPkgs = getSupportedMacPkgs()

// Script binds the scripts payload to a uuid for storing and retrieiving from aws
type Script struct{
    UUID    string
    Payload string
}

// Init will initialize the modules filepath for where to save the generated files
func Init(path string, debugMode bool) bool {
   filePath = path
   debug = debugMode
   return true
}

// Helper function for inserting commands into the overall script slice.
// Very useful for not having to call append manually every time
func commander(script *[]string) func(string, int) {
    return func (command string, indent int) {
        indents := strings.Repeat("\t", indent)
        *script = append(*script, indents + command) 
    }
}

// GenerateGeneric will generate a generic developers setup for the user to
// to run locally
func GenerateGeneric() (Script, error) {
    if filePath == "" {
        return Script{}, errors.New("The current file path isnt set")
    }

    // Convert the map of categories into 
    var install = SupportedPackages{} 
    mapstructure.Decode(genericMacPkgs, &install)

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

    if debug == true {
        uuid, err := util.CreateFile(filePath, script); if err != nil {
            return Script{}, nil
        }

        // Bind script uuid and payload together
        scriptInfo := Script{
            UUID: uuid,
            Payload: strings.Join(script[:], "\n"),
        }

        return scriptInfo, nil
    }

    uuid, err := util.GenerateUUID(); if err != nil {
        return Script{}, err
    }
    
    scriptInfo := Script{
        UUID: uuid,
        Payload: strings.Join(script[:], "\n"),
    }

    return scriptInfo, nil

}

// GenerateDynamic generates a script based on what the user has entered
func GenerateDynamic(install InstallRequest) (Script, error) {
    script := []string{}
    script = append(script, "#!/bin/bash\n")

    // Setup the setup script
    macos.InstallXCode(commander(&script))
    macos.InstallBrew(commander(&script))
    macos.InstallLangs(commander(&script), install.Languages)
    macos.InstallTerminals(commander(&script), install.Terminals)
    macos.InstallShells(commander(&script), install.Shells)
    macos.InstallBrowsers(commander(&script), install.Browsers)
    macos.InstallEditors(commander(&script), install.Editors)

    // If in debug mode
    if debug == true {

        uuid, err := util.CreateFile(filePath, script); if err != nil {
            return Script{}, err
        }

        generatedScript := Script{
            UUID: uuid,
            Payload: strings.Join(script[:], ""),
        }

        return generatedScript, nil
    }

    // Generate a new UUID
    uuid, err := util.GenerateUUID(); if err != nil {
        return Script{}, err
    }

    // bind the script UUID to it's payload
    scriptInfo := Script{
        UUID: uuid,
        Payload: strings.Join(script[:], "\n"),
    }

    return scriptInfo, nil
}
