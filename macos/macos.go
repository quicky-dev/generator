package macos

// InstallXCode checks if xcode has been installed, which is needed for 
// the user to install any other dependency
func InstallXCode(commander func(string, int)) {
    // Xcode download link
    xCodeInstall := "https://itunes.apple.com/us/app/xcode/id497799835?mt=12"

    // Check if the user has xcode installed and attempt to install the xcode CLI tools
    // if they're not already installed
    commander("if command -v xcode-select &> /dev/null; then", 0)
    commander("echo \"Xcode installed, installing commandline tools if not already installed\"", 1)
    commander("xcode-select --install 2> /dev/null", 1)

    // Handle if its 
    commander("else", 0)
    commander("echo \"Xcode is not installed, would you like to install it now? (y/n)\"", 1)
    commander("read installXCode", 1)
    commander("if [ $installXCode = \"y\"] || [ $installXCode = \"yes\" ]; then", 1)
    commander("echo \"Opening up the browser to the app stores Xcode page\"", 2)
    commander("open " + xCodeInstall, 2)
    commander("echo \"After installing XCode, please rerun the script and xcode tools shall be installed upon next run\"", 2)
    commander("else", 1)
    commander("echo \" Xcode is needed for anything else to install. Shutting down.\"", 2)
    commander("exit", 1)
    commander("fi", 1)
    commander("fi", 0)
}

func InstallBrew(genericScript []string) []string {
   return genericScript 
}

