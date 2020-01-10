package ubuntu

// InstallBrew grabs brew from github and installs it on the current machine
func InstallSetup(addCmd func(string, int)) {
	addCmd("sudo apt update", 0)
	addCmd("sudo apt install -y apt-transport-https software-properties-common", 0)
	addCmd("sudo apt install -y ca-certificates wget curl", 0)
}
