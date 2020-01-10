package ubuntu

// SupportedDatabases is a map of the supported macos databases
var SupportedDatabases = map[string]string{
	"mongoDB": "mongodb",
}

// InstallDatabases will add all requested database items to the script
func InstallDatabases(addCmd func(string, int), databases []string) {
	if len(databases) == 0 {
		return
	}

	addCmd("# Install all databases requested", 0)
	addCmd("echo \"Installing selected tools on to the system\"", 0)

	for _, database := range databases {
		if databasePkg, ok := SupportedDatabases[database]; ok {
			addCmd("brew install "+databasePkg, 0)

			// Handle mongodb install specially
			if databasePkg == "mongodb" {
				addCmd("", 0)
				addCmd("mkdir -p /data/db", 0)
				addCmd("sudo chown -R $USER /data/db", 0)
			}
		}
	}
}
