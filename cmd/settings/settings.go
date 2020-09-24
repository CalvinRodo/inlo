package settings

import "github.com/spf13/viper"

type settings struct {
	LogPath string
	FilesFolder string
	Theme string
	DateLayout string
}

var Settings settings

func Initialize() {
	Settings.LogPath = get("LogDir", "logs")
	Settings.FilesFolder = get("Files", "files")
	Settings.Theme = get("Theme", "nOTTY")
	Settings.DateLayout = get("DateLayout", "2006-01-02")
}

func get(k string, d string) string {
	if viper.IsSet(k) {
		return viper.GetString(k)
	}
	return d
}