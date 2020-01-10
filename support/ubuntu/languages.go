package ubuntu

// SupportedLangs is a map of supported macos programming languages(Currently, only latest)
var SupportedLangs = map[string]string{
	"python": "python3.6",
	"ruby":   "ruby-full",
	"node":   "nodejs",
	"go":     "go",
	"rust":   "rust",
	"java":   "java",
}

// install the go 1.13.3 executable and configure it's path.
func installGolang(addCmd func(string, int)) {
	addCmd("pushd $(mktemp -d)", 0)
	addCmd("wget https://dl.google.com/go/go1.13.3.linux-amd64.tar.gz", 0)
	addCmd("sudo tar -xvf go1.13.3.linux-amd64.tar.gz", 0)
	addCmd("sudo mv go /usr/local", 0)
	addCmd("popd  # mktemp", 0)
	addCmd("export GOROOT=/usr/local/go", 0)
	addCmd("export PATH=$GOPATH/bin:$GOROOT/bin:$PATH", 0)
}

// Install rust via rustup.
func installRust(addCmd func(string, int)) {
	addCmd("sudo apt install build-essential", 0)
	addCmd("curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y", 0)
	addCmd("source $HOME/.cargo/env", 0)
}

// install the default java runtime and developer kit.
func installJava(addCmd func(string, int)) {
	addCmd("sudo apt install default-jre", 0)
	addCmd("sudo apt install default-jdk", 0)
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
			if lang == "go" {
				installGolang(addCmd)
			} else if lang == "rust" {
				installRust(addCmd)
			} else if lang == "java" {
				installJava(addCmd)
			} else {
				addCmd("sudo apt install "+langPkg, 0)
			}
		}
	}
	addCmd("", 0)
}
