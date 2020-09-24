package attach

import (
	"bufio"
	"fmt"
	"inlo/cmd/folder"
	"inlo/cmd/halt"
	"inlo/cmd/settings"
	"io"
	"os"
	"path/filepath"
)

// CopyFile creates a file in the log from a reader
func CopyFile(fileName string, reader io.Reader) {

	to, err := os.OpenFile(toPath(fileName), os.O_RDWR|os.O_CREATE, 0666)
	defer to.Close()
	halt.IfErr(err)

	_, err = io.Copy(to, reader)
	halt.IfErr(err)
}

func ReadFile(fileName string) []string {

	file, err := os.Open(fileName)
	defer file.Close()
	halt.IfErr(err)

	var fileContents []string

	lineReader := bufio.NewScanner(file)

	for lineReader.Scan() {
		fileContents = append(fileContents, lineReader.Text())
	}
	halt.IfErr(lineReader.Err())

	return fileContents
}


// MdFileLink Prints out a markdown file link
func MdFileLink(fileName string) string {
	return fmt.Sprintf("[%s](%s)", fileName, toPath(fileName))
}

func toPath(fileName string) string {
	filesPath := filepath.Join(settings.Settings.LogPath, settings.Settings.FilesFolder)
	return filepath.Join(folder.MakeOrGetDir(filesPath), fileName)
}
