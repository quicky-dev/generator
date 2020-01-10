package macos

// SupportedShells is a map of supported macos Shells
var SupportedShells = map[string]string{
	"zsh":  "zsh",
	"fish": "fish",
}

// InstallShells will add all requested shells setup items to the script
func InstallShells(addCmd func(string, int), shells []string) {

	// If there are no shells to install, return
	if len(shells) == 0 {
		return
	}
	// Temporary workaround for working with multiple shells, the first one is always the primary
	primary := shells[0]

	addCmd("# Install all shells requested", 0)
	addCmd("echo \"Installing selected shells on to the system\"", 0)

	for _, shell := range shells {
		if shellPkg, ok := SupportedShells[shell]; ok {
			addCmd("brew install "+shellPkg, 0)

			// Set the just installed shell to be the primary one
			if shellPkg == primary {
				addCmd("", 0)
				addCmd("# Configuring the shell you selected as your primary shell to be just that", 0)
				addCmd("echo \"Setting "+shellPkg+" to be your primary shell.\"", 0)
				addCmd("sudo -s \"echo /usr/local/bin/"+shellPkg+" >> /etc/shells\" && chsh -s /usr/local/bin/"+shellPkg, 0)
				addCmd("", 0)
			}
		}
	}
	addCmd("", 0)
}
