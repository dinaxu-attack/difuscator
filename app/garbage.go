package app

import (
	"crypto/rand"
	"math/big"
	"os"
)

func Garbage(file string) error {
	target, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer target.Close()
	targetSt, _ := target.Stat()
	targetSize := targetSt.Size()

	_, err = target.WriteString(GenGarbage(targetSize + Random(1000, 4099) ^ 2))
	if err != nil {
		return err
	}

	return nil
}

func Random(min, max int64) int64 {
	bg := big.NewInt(max - min)

	n, err := rand.Int(rand.Reader, bg)
	if err != nil {
		panic(err)
	}

	return n.Int64() + min
}

func GenGarbage(size int64) string {

	garb := make([]byte, size)

	_, err := rand.Read(garb)

	if err != nil {
		panic(err)
	}

	return string(garb)
}
