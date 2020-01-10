package ubuntu

// SupportedEditors is a mpa of the supported macos Editors
var SupportedEditors = map[string]string{
	"atom":         "atom",
	"vscode":       "visual-studio-code",
	"vim":          "vim",
	"sublime-text": "sublime-text",
}

func installAtom(addCmd func(string, int)) {
	addCmd("pushd $(mktemp -d)", 0)
	addCmd("wget -O atom-amd64.deb https://atom.io/download/deb", 0)
	addCmd("sudo dpkg -i atom-amd64.deb", 0)
	addCmd("popd  # mktemp -d", 0)
}

func installVisualStudio(addCmd func(string, int)) {
	addCmd("pushd $(mktemp -d)", 0)
	addCmd("sudo apt install -y software-properties-common apt-transport-https", 0)
	addCmd("wget -q https://packages.microsoft.com/keys/microsoft.asc -O- | sudo apt-key add -", 0)
	addCmd("sudo add-apt-repository \"deb [arch=amd64] https://packages.microsoft.com/repos/vscode stable main\"", 0)
	addCmd("popd  # mktemp -d", 0)
	addCmd("sudo apt update", 0)
	addCmd("sudo apt install -y code", 0)
}

func installSublimeText(addCmd func(string, int)) {
	addCmd("pushd $(mktemp -d)", 0)
	addCmd("sudo apt install -y apt-transport-https ca-certificates curl software-properties-common", 0)
	addCmd("curl -fsSL https://download.sublimetext.com/sublimehq-pub.gpg | sudo apt-key add -", 0)
	addCmd("sudo add-apt-repository \"deb https://download.sublimetext.com/ apt/stable/\"", 0)
	addCmd("sudo apt update", 0)
	addCmd("sudo apt install sublime-text", 0)
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
			if editorPkg == "atom" {
				installAtom(addCmd)
			} else if editorPkg == "visual-studio-code" {
				installVisualStudio(addCmd)
			} else if editorPkg == "vim" {
				addCmd("sudo apt install -y vim", 0)
			} else if editorPkg == "sublime-text" {
				installSublimeText(addCmd)
			}
		}
	}
	addCmd("", 0)
}
