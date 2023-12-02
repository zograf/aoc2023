package main

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(path string) []string {
	file, err := os.Open(path)
	ErrCheck(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ret := make([]string, 0)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	ErrCheck(scanner.Err())
	return ret
}

func ErrCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
