package generator

import (
	"errors"
	"github.com/mitchellh/mapstructure"
	"github.com/quicky-dev/generator/support/ubuntu"
	"github.com/quicky-dev/generator/util"
	"strings"
)

type UBUNTU_GENERATOR struct {
	AvailablePackages SupportedPackages
	genericPackages   map[string]category
}

// supportedmacPkgs gets us the list of supported packages
func getSupportedUbuntuPackages() SupportedPackages {
	// All the available categories we have
	packageCategories := []string{
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
		"Select all programming languages of your choice:",
		"Select all Terminal emulators of your choice:",
		"Select all terminal shells of your choice:",
		"Select all web browsers of your choice:",
		"Select all text editors of your choice:",
		"Select all developer tools of your choice:",
		"Select all databases of your choice:",
	}

	// All the available packages we have for the items so far
	pkgItems := []map[string]string{
		ubuntu.SupportedLangs,
		ubuntu.SupportedTerminals,
		ubuntu.SupportedShells,
		ubuntu.SupportedBrowsers,
		ubuntu.SupportedEditors,
		ubuntu.SupportedTools,
		ubuntu.SupportedDatabases,
	}

	// the pkg map to map strings (category names) to category structs
	macPkgMap := map[string]category{}

	// Iterate over all the categories
	for index, currCategoryName := range packageCategories {
		var currentItems []string

		// Append all keys (package names) into the current items string array
		for k := range pkgItems[index] {
			currentItems = append(currentItems, k)
		}

		// Create the category struct which will keep the current packages description
		// and it's Items bundled together
		categoryStruct := category{
			Description: pkgDesc[index],
			Items:       currentItems,
		}

		// Map the current category name to the category struct
		macPkgMap[currCategoryName] = categoryStruct
	}

	// Create the install request and bind it into the mapStructure
	macPackages := SupportedPackages{}
	mapstructure.Decode(macPkgMap, &macPackages)
	return macPackages
}

// GenerateGeneric will generate a generic developers setup for the user to
// to run locally
func (generator UBUNTU_GENERATOR) GenerateGenericScript() (Script, error) {
	if filePath == "" {
		return Script{}, errors.New("The current file path isnt set")
	}

	// Convert the map of categories into
	var install = SupportedPackages{}
	mapstructure.Decode(generator.genericPackages, &install)

	script := []string{}

	// Add shebang to the top of the file to ensure that bash
	// executes the file
	script = append(script, "#! /bin/bash\n")

	// Setup script
	ubuntu.InstallSetup(commander(&script))
	ubuntu.InstallLangs(commander(&script), install.Languages.Items)
	ubuntu.InstallTerminals(commander(&script), install.Terminals.Items)
	ubuntu.InstallShells(commander(&script), install.Shells.Items)
	ubuntu.InstallBrowsers(commander(&script), install.Browsers.Items)
	ubuntu.InstallEditors(commander(&script), install.Editors.Items)
	ubuntu.InstallTools(commander(&script), install.Tools.Items)
	ubuntu.InstallDatabases(commander(&script), install.Databases.Items)

	if debug == true {
		uuid, err := util.CreateFile(filePath, script)
		if err != nil {
			return Script{}, nil
		}

		// Bind script uuid and payload together
		scriptInfo := Script{
			UUID:    uuid,
			Payload: strings.Join(script[:], "\n"),
		}

		return scriptInfo, nil
	}

	uuid, err := util.GenerateUUID()
	if err != nil {
		return Script{}, err
	}

	scriptInfo := Script{
		UUID:    uuid,
		Payload: strings.Join(script[:], "\n"),
	}

	return scriptInfo, nil

}

// GenerateDynamic generates a script based on what the user has entered
func (generator UBUNTU_GENERATOR) GenerateDynamicScript(install InstallRequest) (Script, error) {
	script := []string{}
	script = append(script, "#!/bin/bash\n")

	// Setup the setup script
	ubuntu.InstallXCode(commander(&script))
	ubuntu.InstallBrew(commander(&script))
	ubuntu.InstallLangs(commander(&script), install.Languages)
	ubuntu.InstallTerminals(commander(&script), install.Terminals)
	ubuntu.InstallShells(commander(&script), install.Shells)
	ubuntu.InstallBrowsers(commander(&script), install.Browsers)
	ubuntu.InstallEditors(commander(&script), install.Editors)
	ubuntu.InstallTools(commander(&script), install.Tools)
	ubuntu.InstallDatabases(commander(&script), install.Databases)

	// If in debug mode
	if debug == true {

		uuid, err := util.CreateFile(filePath, script)
		if err != nil {
			return Script{}, err
		}

		generatedScript := Script{
			UUID:    uuid,
			Payload: strings.Join(script[:], ""),
		}

		return generatedScript, nil
	}

	// Generate a new UUID
	uuid, err := util.GenerateUUID()
	if err != nil {
		return Script{}, err
	}

	// bind the script UUID to it's payload
	scriptInfo := Script{
		UUID:    uuid,
		Payload: strings.Join(script[:], "\n"),
	}

	return scriptInfo, nil
}

func GetUbuntuGenerator() UBUNTU_GENERATOR {
	var genericUbuntuPkgs = map[string]category{
		"Languages": {
			Description: "Select all programming languages of your choice",
			Items:       []string{"python", "ruby", "java"},
		},
		"Terminals": {
			Description: "Select all terminal emulators of your choice",
			Items:       []string{},
		},
		"Shells": {
			Description: "Select all Terminal shells of your choice",
			Items:       []string{"hyper", "iterm2", "yeet"},
		},
		"Browsers": {
			Description: "Select the web browser of your choice",
			Items:       []string{"google chrome", "opera"},
		},
		"Editors": {
			Description: "Select the text editor of your choice",
			Items:       []string{"vim", "macvim", "sublime-text"},
		},
		"Tools": {
			Description: "Select the toolings of your choice",
			Items:       []string{"vagrant", "heroku", "docker"},
		},
		"Databases": {
			Description: "Select the databases of your choice",
			Items:       []string{"mongoDB"},
		},
	}

	return UBUNTU_GENERATOR{getSupportedUbuntuPackages(), genericUbuntuPkgs}
}
