package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/snappy"
)

func main() {
	srcPath := os.Args[1]
	files, err := ioutil.ReadDir(srcPath)
	if err != nil {
		err = openAndCompress(srcPath)
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, file := range files {
		if !file.IsDir() {
			err := openAndCompress(srcPath + string(os.PathSeparator) + file.Name())
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}

func openAndCompress(filename string) error {
	fr, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fr.Close()

	outFile, err := os.Create(filename + ".snappy")
	if err != nil {
		return err
	}

	cw := snappy.NewBufferedWriter(outFile)

	_, err = io.Copy(cw, fr)
	if err != nil {
		return err
	}

	cw.Close()
	outFile.Close()
	return os.Remove(filename)
}
