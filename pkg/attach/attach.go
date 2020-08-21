package attach

import (
	"fmt"
	"inlo/consts"
	"inlo/halt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// CopyFile creates a file in the log from a reader
func CopyFile(fileName string, reader io.Reader) {

	to, err := os.OpenFile(toPath(fileName), os.O_RDWR|os.O_CREATE, 0666)
	halt.IfErr(err)
	defer to.Close()

	_, err = io.Copy(to, reader)
	halt.IfErr(err)
}

// MdFileLink Prints out a markdown file link
func MdFileLink(fileName string) string {
	return fmt.Sprintf("[%s](%s)", fileName, toPath(fileName))
}

func toPath(fileName string) string {
	return filepath.Join(viper.GetString(consts.LOGDIR), fileName)
}
