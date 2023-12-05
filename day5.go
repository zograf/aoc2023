package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type ConversionMap struct {
	destination int64
	source      int64
	length      int64
}

type SeedRange struct {
	start int64
	end   int64
}

func FindLowestLocation_Pt2(input []string) {
	original_seeds := ExtractSeeds_Pt2(input[0])
	var seeds []SeedRange = make([]SeedRange, len(original_seeds))
	copy(seeds, original_seeds)

	for i := 3; i < len(input); i++ {
		var conversions []ConversionMap = make([]ConversionMap, 0)
		for line := input[i]; len(line) > 2 && i < len(input); i++ {
			parsed := ParseLine(line)
			conversions = append(conversions, ConversionMap{parsed[0], parsed[1], parsed[2]})
			if i+1 == len(input) {
				break
			}
			line = input[i+1]
		}
		i++
		sort.Slice(conversions, func(i, j int) bool {
			return conversions[i].source < conversions[j].source
		})
		new_seeds := ApplyMap_Pt2(conversions, seeds)
		seeds = new_seeds
	}

	var min int64 = math.MaxInt64
	for _, s := range seeds {
		if s.start <= min {
			min = s.start
		}
	}

	fmt.Println(min)
}

func ApplyMap_Pt2(conversions []ConversionMap, seeds []SeedRange) []SeedRange {
	var new_ranges []SeedRange = make([]SeedRange, 0)
	for _, seed := range seeds {
		found := false
		for _, c := range conversions {
			if seed.start >= c.source && seed.end <= c.source+c.length-1 {
				s := c.destination + seed.start - c.source
				e := c.destination + seed.end - c.source

				new_ranges = append(new_ranges, SeedRange{s, e})
				found = true
				break
			} else if seed.start < c.source && seed.end <= c.source+c.length-1 && seed.end >= c.source {
				seed.end = c.source - 1
				s := c.destination
				e := c.destination + seed.end - c.source

				new_ranges = append(new_ranges, SeedRange{s, e})
				new_ranges = append(new_ranges, SeedRange{seed.start, seed.end})
				found = true
			} else if seed.start >= c.source && seed.end > c.source+c.length-1 && seed.start <= c.source+c.length-1 {
				seed.start = c.source + c.length - 1

				s := c.destination + seed.start - c.source
				e := c.destination + c.length - 1

				new_ranges = append(new_ranges, SeedRange{s, e})
				new_ranges = append(new_ranges, SeedRange{seed.start, seed.end})
				found = true
			} else if seed.start < c.source && seed.end > c.source+c.length-1 {
				new_seed_first := SeedRange{seed.start, c.source - 1}
				new_seed_second := SeedRange{c.source + c.length, seed.end}
				s := c.destination
				e := c.destination + c.length - 1

				new_ranges = append(new_ranges, SeedRange{s, e})
				new_ranges = append(new_ranges, new_seed_first)
				new_ranges = append(new_ranges, new_seed_second)
				found = true
				seeds = append(seeds, new_seed_first)
				seeds = append(seeds, new_seed_second)
				break
			}
		}
		if !found {
			new_ranges = append(new_ranges, SeedRange{seed.start, seed.end})
		}
	}
	return new_ranges
}

func ExtractSeeds_Pt2(line string) []SeedRange {
	var ret []SeedRange = make([]SeedRange, 0)
	seeds := ParseLine(strings.Split(line, ": ")[1])
	for i := 0; i < len(seeds); i += 2 {
		range_start := seeds[i]
		range_size := seeds[i+1]
		ret = append(ret, SeedRange{range_start, range_start + range_size - 1})
	}
	return ret
}

func FindLowestLocation_Pt2_BruteForce(input []string) {
	original_seeds := ExtractSeeds_Pt2_BruteForce(input[0])
	var seeds []int64 = make([]int64, len(original_seeds))
	copy(seeds, original_seeds)

	for i := 3; i < len(input); i++ {
		var conversions []ConversionMap = make([]ConversionMap, 0)
		for line := input[i]; len(line) > 2 && i < len(input); i++ {
			parsed := ParseLine(line)
			conversions = append(conversions, ConversionMap{parsed[0], parsed[1], parsed[2]})
			if i+1 == len(input) {
				break
			}
			line = input[i+1]
		}
		i++
		ApplyMap_Pt1(conversions, seeds)
	}

	fmt.Println(slices.Min(seeds))
}

func ExtractSeeds_Pt2_BruteForce(line string) []int64 {
	var ret []int64 = make([]int64, 0)
	seeds := ParseLine(strings.Split(line, ": ")[1])
	for i := 0; i < len(seeds); i += 2 {
		range_start := seeds[i]
		range_size := seeds[i+1]
		for j := range_start; j < range_start+range_size; j++ {
			ret = append(ret, j)
		}
	}
	return ret
}

func FindLowestLocation_Pt1(input []string) {
	original_seeds := ExtractSeeds_Pt1(input[0])
	var seeds []int64 = make([]int64, len(original_seeds))
	copy(seeds, original_seeds)

	for i := 3; i < len(input); i++ {
		var conversions []ConversionMap = make([]ConversionMap, 0)
		for line := input[i]; len(line) > 2 && i < len(input)-1; i++ {
			parsed := ParseLine(line)
			conversions = append(conversions, ConversionMap{parsed[0], parsed[1], parsed[2]})
			line = input[i+1]
		}
		i++
		ApplyMap_Pt1(conversions, seeds)
	}
	fmt.Println(slices.Min(seeds))
}

func ApplyMap_Pt1(conversions []ConversionMap, seeds []int64) {
	for i := 0; i < len(seeds); i++ {
		seed := seeds[i]
		for _, c := range conversions {
			if seed >= c.source && seed <= c.source+c.length-1 {
				seeds[i] = c.destination + seed - c.source
			}
		}
	}
}

func ExtractSeeds_Pt1(line string) []int64 {
	return ParseLine(strings.Split(line, ": ")[1])
}

func ParseLine(line string) []int64 {
	var ret []int64 = make([]int64, 0)
	nums_str := strings.Split(line, " ")
	for _, seed := range nums_str {
		val, _ := strconv.ParseInt(seed, 10, 64)
		ret = append(ret, val)
	}
	return ret
}
