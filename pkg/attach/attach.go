package attach

import (
	"bufio"
	"fmt"
	"inlo/consts"
	"inlo/halt"
	"inlo/pkg/il"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// File Attach a file to the log
func File(fileName string, from *os.File) {

	to := openFile(toPath(fileName))
	defer to.Close()

	_, err := io.Copy(to, from)
	halt.IfErr(err)

	il.PrintLine("FILEATTACHED", fileLink(fileName))
	fmt.Printf("File %s attached\n", fileName)
}

// Buffer Attach a buffer to the log
func Buffer(fileName string, buffer *bufio.Reader) {

	to := openFile(toPath(fileName))
	defer to.Close()

	_, err := io.Copy(to, buffer)
	halt.IfErr(err)

	il.PrintLine("STDIN", fileLink(fileName))

}

func fileLink(fileName string) string {
	return fmt.Sprintf("[%s](%s)", fileName, toPath(fileName))
}

func toPath(fileName string) string {
	return filepath.Join(viper.GetString(consts.LOGDIR), fileName)
}

func openFile(path string) *os.File {
	flags := os.O_RDWR | os.O_CREATE

	to, err := os.OpenFile(path, flags, 0666)
	halt.IfErr(err)

	return to
}
