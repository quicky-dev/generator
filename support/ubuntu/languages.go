package ubuntu

// SupportedLangs is a map of supported macos programming languages(Currently, only latest)
var SupportedLangs = map[string]string{
	"python": "python",
	"ruby":   "ruby",
	"node":   "node",
	"go":     "go",
	"rust":   "rust",
	"java":   "java",
}

// InstallLangs installs all languages that are supported by the script factory
func InstallLangs(addCmd func(string, int), langs []string) {

	if len(langs) == 0 {
		return
	}

	addCmd("# Install all languages requested", 0)
	addCmd("echo \"Installing selected languages on to the system\"", 0)

	// Iterate over all the selected languages
	for _, lang := range langs {
		if langPkg, ok := SupportedLangs[lang]; ok {
			// Java needs to be installed with cask
			if lang == "java" {
				addCmd("brew cask install "+langPkg, 0)
			} else {
				addCmd("brew install "+langPkg, 0)
			}
		}
	}
	addCmd("", 0)
}
