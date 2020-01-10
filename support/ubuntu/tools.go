package ubuntu

// SupportedTools is a map of the supported macos toolings
var SupportedTools = map[string]string{
	"docker":     "docker",
	"heroku":     "heroku",
	"vlc":        "vlc",
	"insomnia":   "insomnia",
	"redshift":   "redshift",
	"postman":    "postman",
	"vagrant":    "vagrant",
	"virtualbox": "virtualbox",
	"caffeine":   "caffeine",
}

func installDocker(addCmd func(string, int)) {
	addCmd("sudo apt remove docker docker-engine docker.io", 0)
	addCmd("sudo apt install -y docker.io", 0)
	addCmd("sudo systemctl start docker", 0)
	addCmd("sudo systemctl enable docker", 0)
}

func installVagrant(addCmd func(string, int)) {
	addCmd("pushd $(mktemp -d)", 0)
	addCmd("curl -O https://releases.hashicorp.com/vagrant/2.2.6/vagrant_2.2.6_x86_64.deb", 0)
	addCmd("sudo apt install ./vagrant_2.2.6_x86_64.deb", 0)
	addCmd("popd  # mktemp -d", 0)
}

// InstallTools will add all requested tool setup items to the script
func InstallTools(addCmd func(string, int), tools []string) {
	if len(tools) == 0 {
		return
	}

	addCmd("# Install all tools requested", 0)
	addCmd("echo \"Installing selected tools on to the system\"", 0)

	for _, tool := range tools {
		if toolPkg, ok := SupportedTools[tool]; ok {
			if toolPkg == "heroku" {
				addCmd("sudo snap install --classic heroku", 0)
			} else if toolPkg == "vlc" {
				addCmd("sudo snap install vlc", 0)
			} else if toolPkg == "insomnia" {
				addCmd("sudo snap install insomnia", 0)
			} else if toolPkg == "postman" {
				addCmd("sudo snap install postman", 0)
			} else if toolPkg == "redshift" {
				addCmd("sudo apt install -y redshift redshift-gtk", 0)
			} else if toolPkg == "virtualbox" {
				addCmd("sudo apt install -y virtualbox", 0)
			} else if toolPkg == "caffeine" {
				addCmd("sudo apt install -y caffeine", 0)
			} else if toolPkg == "vagrant" {
				installVagrant(addCmd)
			} else if toolPkg == "docker" {
				installDocker(addCmd)
			}
		}
	}
	addCmd("", 0)
}
