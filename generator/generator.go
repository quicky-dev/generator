package generator

import (
	"strings"
)

var filePath string
var debug bool

// InstallRequest is used for unmarshalling json objects directly from a user request into an install request
type InstallRequest struct {
	Languages []string `json:"Languages"`
	Terminals []string `json:"Terminals"`
	Shells    []string `json:"Shells"`
	Browsers  []string `json:"Browsers"`
	Editors   []string `json:"Editors"`
	Tools     []string `json:"Tools"`
	Databases []string `json:"Databases"`
}

type category struct {
	Description string   `json:"Description"`
	Items       []string `json:"Items"`
}

// SupportedPackages is a struct for maintaining supported packages within our factory
type SupportedPackages struct {
	Languages category `json:"Languages"`
	Terminals category `json:"Terminals"`
	Shells    category `json:"Shells"`
	Browsers  category `json:"Browsers"`
	Editors   category `json:"Editors"`
	Tools     category `json:"Tools"`
	Databases category `json:"Databases"`
}

// Script binds the scripts payload to a uuid for storing and retrieiving from aws
type Script struct {
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
	return func(command string, indent int) {
		indents := strings.Repeat("\t", indent)
		*script = append(*script, indents+command)
	}
}
