package app

import "os"

var Old int64
var New int64

func Stat(file string) {

	target, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer target.Close()
	st, _ := target.Stat()
	res := st.Size()

	New = res

}
