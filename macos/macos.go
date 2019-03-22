package macos


// Supported languages (Currently, only latest)
var SupportedLangs = map[string]string {
    "python":"python",
    "ruby": "ruby",
    "node": "node",
    "go": "go",
    "rust": "rust",
    "java": "java",
}

// Supported terminals
var SupportedTerminals = map[string]string {
    "iterm2": "iterm2",
    "hyper": "hyper",
}


// Supported shells
var SupportedShells = map[string]string {
    "zsh": "zsh",
    "fish": "fish",
}

// Supported browsers
var SupportedBrowsers = map[string]string {
    "google chrome": "google-chrome",
    "firefox": "firefox",
    "opera": "opera",
}

var SupportedEditors = map[string]string {
    "atom":"atom",
    "vscode": "visual-studio-code",
    "vim": "vim",
    "macvim": "macvim",
    "sublime-text": "sublime-text",
}

var SupportedTools = map[string]string {
    "basictex": "basictex",
    "cheatsheet": "cheatsheet",
    "docker": "docker",
    "heroku": "heroku",
    "insomnia": "insomnia",
    "mactext": "mactext",
    "postman": "postman",
    "vagrant": "vagrant",
    "caffeine": "caffeine",
}



var SupportedDatabases = map[string]string {

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
    
    if len(langs) == 0 {
        return 
    }

    addCmd("# Install all languages requested", 0)
    addCmd("echo \"Installing selected languages on to the system\"", 0)

    // Iterate over all the selected languages
    for _, lang := range langs {
       if langPkg, ok := SupportedLangs[lang]; ok {
            // Java needs to be installed with cask
            if lang == "java" {
                addCmd("brew cask install " + langPkg, 0)
            }else {
                addCmd("brew install " + langPkg, 0)
            }
       }
    }
    addCmd("", 0) 
}

// InstallTerminals will add all requested terminal emulator setup items to the script
func InstallTerminals(addCmd func(string, int), terminals []string){

    // If tehre are no terminals to install, return
    if len(terminals) == 0 {

    }

    addCmd("# Install all terminals requested", 0)
    addCmd("echo \"Installing selected terminal emulators on to the system\"", 0)

    // Iterate over all the selected terminals
    for _, terminal := range terminals {
        if terminalPkg, ok := SupportedTerminals[terminal]; ok {
            addCmd("brew cask install " + terminalPkg, 0)
        }
    }
    addCmd("", 0)
}

// InstallShells will add all requested shells setup items to the script
func InstallShells(addCmd func(string, int), shells []string) {

    // Temporary workaround for working with multiple shells, the first one is always the primary
    primary := shells[0]
    // If there are no shells to install, return
    if len(shells) == 0 {
        return
    }

    addCmd("# Install all shells requested", 0)
    addCmd("echo \"Installing selected shells on to the system\"", 0)

    for _, shell := range shells {
        if shellPkg, ok := SupportedShells[shell]; ok {
            addCmd("brew install " + shellPkg, 0)

            // Set the just installed shell to be the primary one
            if shellPkg == primary {
                addCmd("", 0)
                addCmd("# Configuring the shell you selected as your primary shell to be just that", 0)
                addCmd("echo \"Setting " + shellPkg + " to be your primary shell.\"", 0)
                addCmd("sudo -s \"echo /usr/local/bin/" + shellPkg + " >> /etc/shells\" && chsh -s /usr/local/bin/" + shellPkg, 0)
                addCmd("", 0)
            }
        }
    }
    addCmd("", 0)
}

// InstallBrowsers will add all requested browser setup items to the script
func InstallBrowsers(addCmd func(string, int), browsers []string) {
    if len(browsers) == 0 {
        return
    }

    addCmd("# Install all browsers requested", 0)
    addCmd("echo \"Installing selected browsers on to the system\"", 0)

    for _, browser := range browsers {
        if browserPkg, ok := SupportedBrowsers[browser]; ok {
            addCmd("brew cask install " + browserPkg, 0)
        }
    }
    addCmd("", 0)
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
            }else{
                addCmd("brew cask install " + editorPkg, 0)
            }
        }
    }
    addCmd("", 0)
}

// InstallTools will add all requested tool setup items to the script
func InstallTools(addCmd func(string, int), tools []string) {
    if len(tools) == 0 {
        return
    }

    addCmd("# Install all tools requested", 0)
    addCmd("echo \"Installing selected tools on to the system\"", 0)

    //for _, tool := range tools {
        //if toolPkg, ok := supportedTools[tool]; ok {
              
        //}
    // }
    addCmd("", 0)
}

// InstallDatabases will add all requested database items to the script
func InstallDatabases(addCmd func(string, int)) {

}

