package macos

// SupportedTools is a map of the supported macos toolings
var SupportedTools = map[string]string{
	"basictex":   "basictex",
	"cheatsheet": "cheatsheet",
	"docker":     "docker",
	"heroku":     "heroku",
	"insomnia":   "insomnia",
	"mactex":     "mactex",
	"postman":    "postman",
	"vagrant":    "vagrant",
	"caffeine":   "caffeine",
	"flux":       "flux",
	"ngrok":      "ngrok",
}

var supportedCaskTools = map[string]string{
	"docker":     "docker",
	"vlc":        "vlc",
	"basictex":   "basictext",
	"cheatsheet": "cheatsheet",
	"vagrant":    "vagrant",
	"caffeine":   "caffeine",
	"flux":       "flux",
	"mactex":     "mactext",
	"ngrok":      "ngrok",
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
			cmd := "brew"
			if _, ok := supportedCaskTools[toolPkg]; ok {
				cmd += " cask"
			}
			cmd += " install " + toolPkg
			addCmd(cmd, 0)
		}
	}
	addCmd("", 0)
}
