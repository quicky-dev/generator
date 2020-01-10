package ubuntu

// SupportedShells is a map of supported macos Shells
var SupportedShells = map[string]string{
	"zsh":  "zsh",
	"fish": "fish",
}

// InstallShells will add all requested shells setup items to the script
func InstallShells(addCmd func(string, int), shells []string) {

	// If there are no shells to install, return
	if len(shells) == 0 || len(shells) > 1 {
		return
	}

	addCmd("# Install all shells requested", 0)
	addCmd("echo \"Installing selected shells on to the system\"", 0)

	for _, shell := range shells {
		if shellPkg, ok := SupportedShells[shell]; ok {
			if shellPkg == "zsh" {
				addCmd("sudo apt update && sudo apt upgrade -y", 0)
				addCmd("sudo apt install -y zsh", 0)
			} else if shellPkg == "fish" {
				addCmd("sudo apt install -y fish", 0)
			}

			addCmd("chsh -s /usr/bin/"+shellPkg, 0)
		}
	}
}
