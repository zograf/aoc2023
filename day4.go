package main

import (
	"strconv"
	"strings"
)

type Scratchcard struct {
	number      int
	first_half  []int
	second_half []int
}

func CalculatePoints(input []string) int {
	var cards []Scratchcard = ParseCards(input)
	var score map[int]int = InitMap(cards)

	CalculateScore(score, cards)
	ret := SumScore(score)
	return ret
}

func SumScore(score map[int]int) int {
	var ret int = 0

	for _, v := range score {
		ret += v
	}
	return ret
}

func CalculateScore(score map[int]int, cards []Scratchcard) {
	for _, el := range cards {
		count := CalculateMatches(el)
		for i := el.number + 1; i < el.number+count+1; i++ {
			if _, ok := score[i]; ok {
				score[i] += score[el.number]
			}
		}
	}

}

func InitMap(cards []Scratchcard) map[int]int {
	var ret map[int]int = make(map[int]int)

	for _, el := range cards {
		ret[el.number] = 1
	}
	return ret
}

func ParseCards(input []string) []Scratchcard {
	var ret []Scratchcard = make([]Scratchcard, 0)

	for _, el := range input {
		card_split := strings.Split(el, ": ")
		str := strings.Split(card_split[0], " ")
		number, _ := strconv.Atoi(str[len(str)-1])

		halves := strings.Split(card_split[1], " | ")
		first_half := ReadNumbers(halves[0])
		second_half := ReadNumbers(halves[1])

		ret = append(ret, Scratchcard{number, first_half, second_half})
	}
	return ret
}

func ReadNumbers(input string) []int {
	var ret []int = make([]int, 0)
	for _, el := range strings.Split(input, " ") {
		val, _ := strconv.Atoi(el)
		if val != 0 {
			ret = append(ret, val)
		}
	}
	return ret
}

func CalculateMatches(card Scratchcard) int {
	var ret int = 0

	for _, i := range card.first_half {
		for _, j := range card.second_half {
			if i == j {
				ret++
			}
		}
	}
	return ret
}
