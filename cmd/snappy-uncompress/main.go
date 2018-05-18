package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/golang/snappy"
)

func main() {
	srcPath := os.Args[1]
	files, err := ioutil.ReadDir(srcPath)
	if err != nil {
		err = openAndUncompress(srcPath)
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, file := range files {
		if !file.IsDir() {
			err := openAndUncompress(srcPath + string(os.PathSeparator) + file.Name())
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}

func openAndUncompress(filename string) error {
	fr, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fr.Close()

	outFile, err := os.Create(strings.Replace(filename, ".snappy", "", 1))
	if err != nil {
		return err
	}
	sr := snappy.NewReader(fr)

	_, err = io.Copy(outFile, sr)
	if err != nil {
		return err
	}
	err = os.Remove(filename)
	return err
}
