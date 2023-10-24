package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

func fileSha1(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha1.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil

}

func main() {
	sig, err := fileSha1("httpd.log.gz")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sig)
}
