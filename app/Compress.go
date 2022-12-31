package app

import (
	"bytes"
	"compress/gzip"
)

func Compress(text []byte) ([]byte, error) {
	var compressed bytes.Buffer
	writer, err := gzip.NewWriterLevel(&compressed, gzip.BestCompression)

	_, err = writer.Write(text)
	writer.Close()

	if err != nil {
		return []byte(""), err
	}

	return compressed.Bytes(), nil
}
