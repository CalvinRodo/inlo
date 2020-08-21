package halt

// IfErr Panic if there is an Error
func IfErr(err error) {
	if err != nil {
		panic(err)
	}
}
