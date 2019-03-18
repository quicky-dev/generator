package macos

// InstallXCode checks if xcode has been installed, which is needed for 
// the user to install any other dependency
func InstallXCode(genericScript []string) []string {
    // Xcode download link
    xCodeInstall := "https://itunes.apple.com/us/app/xcode/id497799835?mt=12"

    // setup found flag, for if xcode is found
    genericScript = append(genericScript, "found=false\n")
    
    // Setup loop to keep running if xcode isn't found
    genericScript = append(genericScript, "while [ $found = \"false\" ]; do")
    
    // Check if xcode has been installed, then install command line tools if so 
    genericScript = append(genericScript, "\t" + "if command -v xcode-select &> /dev/null; then")
    genericScript = append(genericScript, "\t\t" + "found=true")
    genericScript = append(genericScript, "\t\t" + "echo \"Xcode installed, installing commandline tools if not already installed\"")
    genericScript = append(genericScript, "\t\t" + "xcode-select --install 2> /dev/null") 
    genericScript = append(genericScript, "\t" + "else")
    genericScript = append(genericScript, "\t\t" + "printf \"Xcode is not installed, would you like to install it now? (y/n): \"")
    genericScript = append(genericScript, "\t\t" + "read installXCode")
    
    // Check if the user had entered yes for installing xcode, if so link them to xcode
    genericScript = append(genericScript, "\t\t" + "if [ $installXCode = \"y\"] || [ $installXCode = \"yes\" ]; then")
    genericScript = append(genericScript, "\t\t\t" + "echo \"Opening up the browser to the app stores Xcode page\"")
    genericScript = append(genericScript, "\t\t\t" + "open " + xCodeInstall)
    genericScript = append(genericScript, "\t\t" + "else")
    genericScript = append(genericScript, "\t\t\t" + "echo \" Xcode is needed for anything else to install. Shutting down.\"")
    genericScript = append(genericScript, "\t\t\t" + "exit")
    genericScript = append(genericScript, "\t\t" + "fi") 
    genericScript = append(genericScript, "\t" + "fi")
   
    // End of the while loop
    genericScript = append(genericScript, "done")

    return genericScript
}

