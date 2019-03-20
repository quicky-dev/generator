package macos

// List of supported browsers that can be added to our script
var supportedBrowsers = map[string]string {
    "google-chrome": "google-chrome",
    "firefox": "firefox",
    "opera": "opera",
}

// supportedLangs is a list of supported languages and versions,
// Redundant for now but will be more useful in the future when language versioning comes into
// play
var supportedLangs = map[string]string {
    "python":"python",
    "ruby": "ruby",
    "node": "node",
    "go": "go",
    "rust": "rust",
}

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

// InstallBrew grabs brew from github and installs it on the current machine
func InstallBrew(addCmd func(string, int)) {
    addCmd("# Installing brew pkg manager", 0)
    addCmd("/usr/bin/ruby -e \"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)\"", 0)
    addCmd("brew tap caskroom/cask", 0)
    addCmd("", 0)
}

// InstallLangs installs all languages that are supported by the script factory
func InstallLangs(addCmd func(string, int), langs []string) {

   addCmd("# Install all languages requested", 0)
   addCmd("echo \"Installing selected languages on to the system\"", 0)

   // Iterate over all the selected languages
   for _, lang := range langs {
       if lang, ok := supportedLangs[lang]; ok {
            addCmd("brew install " + lang, 0)
       }
   }
   addCmd("", 0) 
}

// InstallTerminals will add all requested terminal emulator setup items to the script
func InstallTerminals(addCmd func(string, int)) {

}

// InstallShells will add all requested shells setup items to the script
func InstallShells(addCmd func(string, int)) {
    
}

// InstallBrowsers will add all requested browser setup items to the script
func InstallBrowsers(addCmd func(string, int)) {

}

// InstallEditors will add all requested editor setup items to the script
func InstallEditors(addCmd func(string, int)) {

}

// InstallTools will add all requested tool setup items to the script
func InstallTools(addCmd func(string, int)) {

}

// InstallDatabases will add all requested database items to the script
func InstallDatabases(addCmd func(string, int)) {

}

