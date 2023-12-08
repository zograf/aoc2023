package main

import (
	"slices"
	"strings"
)

func HauntedWasteland(input []string) int64 {
	var steps int = 0
	var ret int64 = 1
	var step_list []int = make([]int, 0)
	instructions := ParseHauntedInstructions(input[0])
	network := ParseHauntedNetwork(input)
	nodes := GetStartingNodes(network)

	for _, node := range nodes {
		i := 0
		steps = 0
		for {
			if i == len(instructions) {
				i = 0
			}
			step := instructions[i]
			steps++

			node = network[node][step]
			if node[2] == 'Z' {
				break
			}

			i++
		}
		step_list = append(step_list, steps)
	}

	factors := FindFactors(step_list)
	for _, f := range factors {
		ret *= int64(f)
	}

	return ret
}

func FindFactors(arr []int) []int {
	var ret []int = make([]int, 0)
	for i := 0; i < len(arr); i++ {
		for j := 2; j < arr[i]/2; j++ {
			if arr[i]%j == 0 && !slices.Contains(ret, j) {
				ret = append(ret, j)
			}
		}
	}
	return ret
}

func IsHauntedEnd(nodes []string) bool {
	for _, node := range nodes {
		if node[2] != 'Z' {
			return false
		}
	}
	return true
}

func GetStartingNodes(network map[string][]string) []string {
	var ret []string = make([]string, 0)

	for k := range network {
		if k[2] == 'A' {
			ret = append(ret, k)
		}
	}

	return ret
}

func ParseHauntedNetwork(lines []string) map[string][]string {
	var ret map[string][]string = make(map[string][]string)

	for i := 2; i < len(lines); i++ {
		strs := strings.Split(lines[i], " = ")
		paths := strings.Split(strs[1], ", ")
		ret[strs[0]] = []string{paths[0][1:], paths[1][:3]}
	}

	return ret
}

func ParseHauntedInstructions(line string) []int {
	var ret []int = make([]int, 0)

	for _, ch := range line {
		if ch == 'L' {
			ret = append(ret, 0)
		} else {
			ret = append(ret, 1)
		}
	}

	return ret
}
