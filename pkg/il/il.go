package il

import (
	"fmt"
	"inlo/pkg/timestamp"
	"log"
	"os"
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
	line := fmt.Sprintf("%s|%s - %s\n", currentTime.Format(lineFormat), cat, message)
	if _, err := file.WriteString(line); err != nil {
		log.Fatal(err)
	}

}

func openOrCreateFile() *os.File {

	date := timestamp.CurrentTime()
	fileName := fmt.Sprintf("%s.md", date.Format(layoutISO))
	flags := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	var osPermissions os.FileMode = 0664

	f, err := os.OpenFile(fileName, flags, osPermissions)
	if err != nil {
		log.Fatal(err)
	}

	return f
}
