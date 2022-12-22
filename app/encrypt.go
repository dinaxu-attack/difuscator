package app

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"io"
	"io/ioutil"
	"os"
)

func AES(text []byte, out string) error {

	f, err := ioutil.ReadFile(out)
	if err != nil {
		return err
	}

	key := sha512.Sum512_256(f)

	c, err := aes.NewCipher(key[:])
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	bCiphertext := gcm.Seal(nonce, nonce, text, nil)

	ciphertext := string([]byte(string(bCiphertext)))

	targ, err := os.OpenFile(out, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	_, err = targ.WriteString(ciphertext)
	if err != nil {
		return err
	}
	return nil
}
