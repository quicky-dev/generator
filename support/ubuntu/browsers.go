package ubuntu

// SupportedBrowsers is a map of supported macos Browsers
var SupportedBrowsers = map[string]string{
	"google chrome": "google-chrome",
	"firefox":       "firefox",
	"opera":         "opera",
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
			addCmd("brew cask install "+browserPkg, 0)
		}
	}
	addCmd("", 0)
}
