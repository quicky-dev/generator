package ubuntu

// SupportedEditors is a mpa of the supported macos Editors
var SupportedEditors = map[string]string{
	"atom":         "atom",
	"vscode":       "visual-studio-code",
	"vim":          "vim",
	"macvim":       "macvim",
	"sublime-text": "sublime-text",
}

// InstallEditors will add all requested editor setup items to the script
func InstallEditors(addCmd func(string, int), editors []string) {
	if len(editors) == 0 {
		return
	}

	addCmd("# Install all editors requested", 0)
	addCmd("echo \"Installing selected editors on to the system\"", 0)

	for _, editor := range editors {
		if editorPkg, ok := SupportedEditors[editor]; ok {
			if editorPkg == "vim" {
				addCmd("brew install vim", 0)
			} else {
				addCmd("brew cask install "+editorPkg, 0)
			}
		}
	}
	addCmd("", 0)
}
