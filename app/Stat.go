package app

import (
	"os"
	"strconv"
)

var Old int64
var New int64

func Stat(file string) string {

	target, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer target.Close()
	st, _ := target.Stat()
	res := st.Size()

	New = res

	var sizeKb int64 = 1024.0
	var sizeMb int64 = sizeKb * sizeKb
	var sizeGb int64 = sizeKb * sizeMb

	var result string
	if res < sizeMb {
		result = strconv.Itoa(int(res/sizeKb)) + " KB -> " + strconv.Itoa(int(New/sizeKb)) + " KB"

	} else if res < sizeGb {
		result = strconv.Itoa(int(res/sizeMb)) + " MB -> " + strconv.Itoa(int(New/sizeMb)) + " MB"
	} else {
		result = "File size error"
	}
	return result
}
