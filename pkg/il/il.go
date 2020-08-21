package il

import (
	"fmt"
	"inlo/consts"
	"inlo/pkg/timestamp"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

const (
	layoutISO  = "2006-01-02"
	lineFormat = "15:04:05"
)

// PrintLine prints out a line to the log
func PrintLine(cat string, message string) {

	file := openOrCreateFile()
	defer file.Close()

	currentTime := timestamp.CurrentTime()
	line := fmt.Sprintf("%s|%s|%s  \n", currentTime.Format(lineFormat), cat, message)
	if _, err := file.WriteString(line); err != nil {
		log.Fatal(err)
	}

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

func openOrCreateFile() *os.File {

	date := timestamp.CurrentTime()
	path := filepath.Join(viper.GetString(consts.LOGDIR), date.Format(layoutISO))

	fileName := fmt.Sprintf("%s.md", path)
	flags := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	var osPermissions os.FileMode = 0664

	f, err := os.OpenFile(fileName, flags, osPermissions)
	if err != nil {
		log.Fatal(err)
	}

	return f
}
