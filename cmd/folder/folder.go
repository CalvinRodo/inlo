package folder

import (
	"github.com/mitchellh/go-homedir"
	"inlo/cmd/halt"
	"log"
	"os"
)

// ExpandDir expand the path
func ExpandDir(path string) string {

	dir, err := homedir.Expand(path)
	halt.IfErr(err)

	return dir
}

// MakeOrGetDir create the directory path if it's not already there
func MakeOrGetDir(path string) string {

	dir := ExpandDir(path)

	if err := os.MkdirAll(dir, 0700 ); err != nil && !os.IsExist(err) {
		log.Panic(err)
	}

	return dir

}
