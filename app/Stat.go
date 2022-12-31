package app

import (
	"os"
	"strconv"
)

var Old int64

func Stat(file string) string {

	target, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer target.Close()
	st, _ := target.Stat()
	res := st.Size()

	var sizeKb int64 = 1024.0

	return strconv.Itoa(int(Old/sizeKb)) + " KB -> " + strconv.Itoa(int(res/sizeKb)) + " KB"

}
