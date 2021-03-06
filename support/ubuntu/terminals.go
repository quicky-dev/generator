package ubuntu

// SupportedTerminals is a map of supported macos terminals
var SupportedTerminals = map[string]string{
	"terminator": "terminator",
	"guake":      "guake",
	"tilda":      "tilda",
	"eterm":      "eterm",
}

// InstallTerminals will add all requested terminal emulator setup items to the script
func InstallTerminals(addCmd func(string, int), terminals []string) {

	// If tehre are no terminals to install, return
	if len(terminals) == 0 {
		return
	}

	addCmd("# Install all terminals requested", 0)
	addCmd("echo \"Installing selected terminal emulators on to the system\"", 0)

	// Iterate over all the selected terminals
	for _, terminal := range terminals {
		if terminalPkg, ok := SupportedTerminals[terminal]; ok {
			addCmd("sudo apt install -y "+terminalPkg, 0)
		}
	}
}
