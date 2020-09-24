package folder

import (
	"github.com/mitchellh/go-homedir"
	"inlo/halt"
	"log"
	"os"
)

func MakeOrGetDir(path string) string {

	dir, err := homedir.Expand(path)
	halt.IfErr(err)

	// Check if directory is created
	if err := os.MkdirAll(dir, 0700 ); err != nil && !os.IsExist(err) {
		log.Panic(err)
	}

	return dir

}
