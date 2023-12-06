package main

import (
	"regexp"
	"strconv"
	"strings"
)

func BoatRace(input []string, part_one bool) int {
	var time []string
	var record []string
	var ret int = 1

	if part_one {
		time, record = ParseBoatRace_Pt1(input)
	} else {
		time, record = ParseBoatRace_Pt2(input)
	}

	for i := 0; i < len(time); i++ {
		count := 0
		t, _ := strconv.Atoi(time[i])
		d, _ := strconv.Atoi(record[i])
		for j := 1; j < t; j++ {
			if j*(t-j) > d {
				count++
			}
		}
		ret *= count
	}
	return ret
}

func ParseBoatRace_Pt2(input []string) ([]string, []string) {
	r, _ := regexp.Compile("[0-9]+")
	time := r.FindAllString(strings.ReplaceAll(input[0], " ", ""), -1)
	record := r.FindAllString(strings.ReplaceAll(input[1], " ", ""), -1)
	return time, record
}

func ParseBoatRace_Pt1(input []string) ([]string, []string) {
	r, _ := regexp.Compile("[0-9]+")
	time := r.FindAllString(input[0], -1)
	record := r.FindAllString(input[1], -1)
	return time, record
}
