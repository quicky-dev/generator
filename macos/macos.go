package macos

// InstallXCode checks if xcode has been installed, which is needed for 
// the user to install any other dependency
func InstallXCode(addCmd func(string, int)) {
    // Xcode download link
    xCodeInstall := "https://itunes.apple.com/us/app/xcode/id497799835?mt=12"

    // Check if the user has xcode installed and attempt to install the xcode CLI tools
    // if they're not already installed
    addCmd("# Install XCode / XCode CLI tools", 0)
    addCmd("if command -v xcode-select &> /dev/null; then", 0)
    addCmd("echo \"Xcode installed, installing commandline tools if not already installed\"", 1)
    addCmd("xcode-select --install 2> /dev/null", 1)
    addCmd("echo \"Successfully installed xcode-cli tools\"", 1)

    // Handle if its not installed
    addCmd("else", 0)
    addCmd("echo \"Xcode is not installed, would you like to install it now? (y/n)\"", 1)
    addCmd("read installXCode", 1)
    addCmd("if [ $installXCode = \"y\"] || [ $installXCode = \"yes\" ]; then", 1)
    addCmd("echo \"Opening up the browser to the app stores Xcode page\"", 2)
    addCmd("open " + xCodeInstall, 2)
    addCmd("echo \"After installing XCode, please rerun the script and xcode tools shall be installed upon next run\"", 2)
    addCmd("else", 1)
    addCmd("echo \" Xcode is needed for anything else to install. Shutting down.\"", 2)
    addCmd("exit", 1)
    addCmd("fi", 1)
    addCmd("fi", 0)
    addCmd("", 0)
}

// Install brew grabs brew from github and installs it on the current machine
func InstallBrew(addCmd func(string, int)) {
    addCmd("# Installing brew pkg manager", 0)
    addCmd("/usr/bin/ruby -e \"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)\"", 0)
}


