package il

import (
	"fmt"
	"inlo/cmd/folder"
	"inlo/cmd/halt"
	"inlo/cmd/settings"
	"inlo/cmd/timestamp"
	"os"
	"path/filepath"
	"strings"
)

const (
	lineFormat = "15:04:05"
)

// PrintLine prints out a line to the log
func PrintLine(cat string, message string) {

	file := openOrCreateFile()
	defer file.Close()

	currentTime := timestamp.CurrentTime()
	line := fmt.Sprintf("%s|%s|%s  \n", currentTime.Format(lineFormat), cat, message)
	_, err := file.WriteString(line)
	halt.IfErr(err)

}

// PrintEvent prints a well known Event
// - INCIDENT ENDED
func PrintEvent(cat string) {
	PrintLine(cat, "")
}

// PrintStrings prints an array of strings as a string joined with a space
func PrintStrings(cat string, args []string) {
	PrintLine(cat, strings.Join(args, " "))
}

func FileNameForToday(dir string) string {
	date := timestamp.CurrentTime()
	path := filepath.Join(dir, date.Format(settings.Settings.DateLayout))
	return fmt.Sprintf("%s.md", path)
}

func openOrCreateFile() *os.File {

	dir := folder.MakeOrGetDir(settings.Settings.LogPath)
	fileName := FileNameForToday(dir)

	flags := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	var osPermissions os.FileMode = 0664

	f, err := os.OpenFile(fileName, flags, osPermissions)
	halt.IfErr(err)

	return f
}
