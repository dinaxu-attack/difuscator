package app

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

var names = []string{}
var target = []string{}

func ReadFiles() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	dir, err := ioutil.ReadDir(cwd)
	if err != nil {
		return err
	}

	for _, i := range dir {
		d, _ := filepath.Abs(i.Name())
		if i.IsDir() == false {
			if strings.Contains(i.Name(), ".go") {
				target = append(target, d)
			}
		} else {
			readDir(d)
		}

	}
	return nil
}

func readDir(path string) {

	dir, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, i := range dir {
		d, _ := filepath.Abs(i.Name())
		if i.IsDir() == false {
			if strings.Contains(i.Name(), ".go") {
				target = append(target, path+"/"+i.Name())
			}
		} else {
			readDir(d)
		}
	}

}

func Obfuscate(file string) error {

	ReadFiles()

	for _, i := range target {
		d := trim(i)
		if d != file {
			res := strings.ReplaceAll(i, d, uniqueName()+".go")
			os.Rename(i, res)
		}
	}
	return nil
}

func uniqueName() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, Random(4, 10))
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
		for _, i := range names {
			if string(s) == i {
				uniqueName()
			} else {
				names = append(names, string(s))
			}
		}
	}
	return string(s)

}

func trim(s string) (prefix string) {
	var strs = strings.SplitAfterN(s, "/", -1)

	d := len(strs)

	return strs[d-1]
}
