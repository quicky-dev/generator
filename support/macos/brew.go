package macos

// InstallBrew grabs brew from github and installs it on the current machine
func InstallBrew(addCmd func(string, int)) {
	addCmd("# Installing brew pkg manager", 0)
	addCmd("/usr/bin/ruby -e \"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)\"", 0)
	addCmd("brew tap caskroom/cask", 0)
	addCmd("", 0)
}
