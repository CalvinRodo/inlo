package il

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	layoutISO  = "2006-01-02"
	lineFormat = "15:04:05"
)

func printTimeLine(file *os.File, currentTime time.Time, message string) {
	line := fmt.Sprintf("%s|TIMELINE - %s\n", currentTime.Format(lineFormat), message)
	if _, err := file.WriteString(line); err != nil {
		log.Fatal(err)
	}
}

func openOrCreateFile(date time.Time) *os.File {

	fileName := fmt.Sprintf("%s.md", date.Format(layoutISO))
	flags := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	var osPermissions os.FileMode = 0664

	f, err := os.OpenFile(fileName, flags, osPermissions)
	if err != nil {
		log.Fatal(err)
	}

	return f
}
