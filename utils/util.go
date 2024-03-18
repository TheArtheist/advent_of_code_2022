package utils

import (
	"bufio"
	"log"
	"os"
)

func GetFileScanner(path string) *bufio.Scanner {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	return scanner

}
