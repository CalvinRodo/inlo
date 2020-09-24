package halt

import "log"

// IfErr Panic if there is an Error
func IfErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
