package ubuntu

// SupportedBrowsers is a map of supported macos Browsers
var SupportedBrowsers = map[string]string{
	"google chrome": "google-chrome",
	"tor":           "tor",
}

func installChrome(addCmd func(string, int)) {
	addCmd("echo \"deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main\"  > /etc/apt/sources.list.d/google-chrome.list", 0)
	addCmd("pushd $(mktemp -d)", 0)
	addCmd("wget https://dl.google.com/linux/linux_signing_key.pub", 0)
	addCmd("sudo apt-key add linux_signing_key.pub", 0)
	addCmd("popd  # mktemp -d", 0)
	addCmd("sudo apt update", 0)
	addCmd("sudo apt install -y google-chrome-stable", 0)
}

func installTor(addCmd func(string, int)) {
	addCmd("sudo apt install apt-transport-https", 0)
	addCmd("sudo add-apt-repository universe", 0)
	addCmd("pushd $(mktemp -d)", 0)
	addCmd("wget -q -0 https://deb.torproject.org/torproject.org/A3C4F0F979CAA22CDBA8F512EE8CBC9E886DDD89.asc | sudo apt-key add -", 0)
	addCmd("echo \"deb https://deb.torproject.org/torproject.org $(lsb_release -cs) main\" | sudo tee -a /etc/apt/sources.list", 0)
	addCmd("sudo apt update", 0)
	addCmd("sudo apt install -y tor deb.torproject.org-keyring torbrowser-launcher", 0)
}

// InstallBrowsers will add all requested browser setup items to the script
func InstallBrowsers(addCmd func(string, int), browsers []string) {
	if len(browsers) == 0 {
		return
	}

	addCmd("# Install all browsers requested", 0)
	addCmd("echo \"Installing selected browsers on to the system\"", 0)

	for _, browser := range browsers {
		if _, ok := SupportedBrowsers[browser]; ok {
			if browser == "google chrome" {
				installChrome(addCmd)
			} else if browser == "tor" {
				installTor(addCmd)
			}
		}
	}
	addCmd("", 0)
}
