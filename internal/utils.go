package internal

import (
	"crypto/rand"
	"math/big"
	random "math/rand"
	"os"
	"path/filepath"
)

var Names = make([]string, 0)

func UniqueName() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, Random(4, 10))
	for i := range s {
		s[i] = letters[random.Intn(len(letters))]
		for _, i := range Names {
			if string(s) == i {
				UniqueName()
			} else {
				Names = append(Names, string(s))
			}
		}
	}
	return string(s)

}

func Random(min, max int64) int64 {
	bg := big.NewInt(max - min)

	n, err := rand.Int(rand.Reader, bg)
	if err != nil {
		panic(err)
	}

	return n.Int64() + min
}

func GetMain(path string) {
	li, _ := os.ReadDir(path)

	for _, i := range li {
		if i.IsDir() {
			GetMain(path + "/" + i.Name())
		} else {
			path, _ := filepath.Abs(path)

			if filepath.Ext(filepath.Join(path, i.Name())) == ".go" {

				readf, _ := os.ReadFile(filepath.Join(path, i.Name()))

				mainfile := GetMainFile(readf, filepath.Join(path, i.Name()))

				MainFile = mainfile

				if mainfile == filepath.Join(path, i.Name()) {
					return
				}

			}

		}

	}

}
