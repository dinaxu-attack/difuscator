package internal

import (
	"crypto/rand"
	"fmt"
	"os"
)

func Garbage(file string) error {
	target, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("\nCan't find the file: " + file)
		os.Exit(0)
	}
	defer target.Close()
	targetSt, _ := target.Stat()
	targetSize := targetSt.Size()

	Old = targetSize

	if targetSize < 5 {
		fmt.Println("\nCan't find the file: " + file)
		os.Exit(0)
	}
	offs := targetSize + Random(5000, 10000)

	for i := 0; i <= 300; i++ {
		_, err = target.WriteString(GenGarbage(targetSize - offs/2)[:Random(1000, 3000)])
	}
	if err != nil {
		fmt.Println("\nGarbage error: " + err.Error())
		return err
	}

	return nil
}

func GenGarbage(size int64) string {
	garb := make([]byte, size)

	_, err := rand.Read(garb)

	if err != nil {
		panic(err)
	}

	return string(garb)
}
