package main

import (
	"sort"
	"strconv"
	"strings"
)

var camel_card_map map[string]int = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 1,
	"Q": 12,
	"K": 13,
	"A": 14}

type Card struct {
	hand  string
	bid   int
	value int
}

func CamelCards(input []string) int {
	var cards []Card = make([]Card, 0)
	var ret int = 0

	for _, el := range input {
		cards = append(cards, ParseCamelCard(el))
	}

	sort.Slice(cards, func(i, j int) bool {
		if cards[i].value < cards[j].value {
			return true
		} else if cards[i].value > cards[j].value {
			return false
		} else {
			for k := 0; k < 5; k++ {
				if camel_card_map[string(cards[i].hand[k])] > camel_card_map[string(cards[j].hand[k])] {
					return false
				} else if camel_card_map[string(cards[i].hand[k])] < camel_card_map[string(cards[j].hand[k])] {
					return true
				}
			}
		}
		return true
	})

	for i := 0; i < len(cards); i++ {
		ret += cards[i].bid * (i + 1)
	}

	return ret
}

func ParseCamelCard(line string) Card {
	split := strings.Split(line, " ")
	num, _ := strconv.Atoi(split[1])
	card := Card{split[0], num, 0}
	CalculateCamelCardValue(&card)
	return card
}

func CalculateCamelCardValue(card *Card) {
	var hand_map map[string]int = make(map[string]int)
	var j int = 0

	for _, ch := range card.hand {
		if ch == 'J' {
			j++
			continue
		}
		if _, ok := hand_map[string(ch)]; ok {
			hand_map[string(ch)]++
		} else {
			hand_map[string(ch)] = 1
		}
	}
	if j == 5 {
		card.value = 7
	} else {
		card.value = GetCardValue(hand_map, j)
	}
}

func GetCardValue(hand_map map[string]int, j int) int {
	max_k, max_v := "", -1
	for k, v := range hand_map {
		if v > max_v {
			max_k = k
			max_v = v
		}
	}
	hand_map[max_k] += j

	if len(hand_map) == 5 {
		return 1
	} else if len(hand_map) == 4 {
		return 2
	} else if len(hand_map) == 3 {
		val := 3
		for _, v := range hand_map {
			if v == 3 {
				val = 4
				break
			}
		}
		return val
	} else if len(hand_map) == 2 {
		val := 5
		for _, v := range hand_map {
			if v == 4 {
				val = 6
				break
			}
		}
		return val
	}
	return 7
}
