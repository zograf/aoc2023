package main

import (
	"strings"
)

func FindNumbers(input string) (int64, int64) {
	var numbers map[string]int = map[string]int{"1": 1, "one": 1, "2": 2, "two": 2, "3": 3, "three": 3, "4": 4, "four": 4, "5": 5, "five": 5, "6": 6, "six": 6, "7": 7, "seven": 7, "8": 8, "eight": 8, "9": 9, "nine": 9}
	var first int64 = -1
	var last int64 = -1
	var first_index int64 = 99999999
	var last_index int64 = -1

	for key, val := range numbers {
		f := strings.Index(input, key)
		l := strings.LastIndex(input, key)
		if f != -1 {
			if int64(f) < first_index {
				first_index = int64(f)
				first = int64(val)
			}
			if int64(l) > last_index {
				last_index = int64(l)
				last = int64(val)
			}
		}
	}

	return first, last
}
