package utils

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func FileRead(path string) (data []byte, err error) {
	data, err = ioutil.ReadFile(path)
	if err != nil {
		log.Printf("File reading error: %s", err)
	}
	return
}

func Getenv(key, fallback string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return fallback
	}
	return value
}

func FileCopy(src, dst string) (err error) {
	var srcfd *os.File
	var dstfd *os.File
	//var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		log.Printf("src file reading error: %s", err)
		return
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		log.Printf("dst file creating error: %s", err)
		return
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		log.Printf("copy error: %s", err)
		return
	}
	return
}
