package main

import (
	"strconv"
	"strings"
)

type SpringRecord struct {
    springs []string
    groups []int
}

func HotSprings(input []string) int {
    var records []SpringRecord
    var ret int = 0

    records = ParseSpringRecords(input)
    for _, record := range records {
        var memo map[string]int = make(map[string]int)
        ret += SpringArrangements(record.springs, 0, record.groups, &memo)
    }

    return ret
}

func SpringArrangements(chars []string, group int, remaining []int, memo *map[string]int) int {
    if len(chars) == 0 {
        if group != 0 && len(remaining) == 1 && group == remaining[0] {
            // There is a group and it matches the last count
            return 1
        }
        if group == 0 && len(remaining) == 0 {
            // There is no group but also no count 
            return 1
        }
        // There is either a group that doesnt match a count
        // or no group but a count exists
        return 0
    }

    if len(remaining) == 0 && group != 0 {
        // There is a group but no more groups allowed
        return 0
    }

    key := strconv.Itoa(len(chars)) + "," + strconv.Itoa(group) + "," + strconv.Itoa(len(remaining))
    if val, ok := (*memo)[key]; ok {
        return val
    }

    val := 0
    if chars[0] == "#" {
        // If broken, increase group count
        val = SpringArrangements(chars[1:], group+1, remaining, memo)
    } else if chars[0] == "." {
        // If working, there are 3 cases
        if group == 0 {
            // 1 - There was no group so just increase the pointer
            val = SpringArrangements(chars[1:], 0, remaining, memo)
        } else if group == remaining[0] {
            // 2 - There was a group and we're closing it
            val = SpringArrangements(chars[1:], 0, remaining[1:], memo)
        }
        // 3 - There was a group but it didn't match remaining[0] so it's impossible
    } else if chars[0] == "?" {
        // If ? there are 2 cases
        if group == 0 {
            // 1 - There is no group so it can act as "#" or "."
            val = SpringArrangements(chars[1:], 1, remaining, memo) + SpringArrangements(chars[1:], 0, remaining, memo)
        } else {
            // 2 - There was a group
            if group == remaining[0] {
                // Closing the group (act as ".")
                val = SpringArrangements(chars[1:], 0, remaining[1:], memo)
            } else {
                // Continuing the group (act as "#")
                val = SpringArrangements(chars[1:], group+1, remaining, memo) 
            }
        }
    } 
    
    (*memo)[key] = val
    return val
}

func ParseSpringRecords(input []string) []SpringRecord {
    var ret []SpringRecord = make([]SpringRecord, 0)

    for _, el := range input {
        str_split := strings.Split(el, " ")
        springs := make([]string, 0)
        groups := make([]int, 0)

        for i := 0; i < 5; i++ {
            for _, ch := range str_split[0] {
                springs = append(springs, string(ch))
            }
            // This took 30min to debug jesus
            if (i != 4) {
                springs = append(springs, "?")
            }
        }

        for i := 0; i < 5; i++ {
            for _, num_str := range strings.Split(str_split[1], ",") {
                num, _ := strconv.Atoi(num_str)
                groups = append(groups, num)
            }
        }
        
        ret = append(ret, SpringRecord{springs, groups})
    }

    return ret
}
