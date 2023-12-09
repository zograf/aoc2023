package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Oasis(input []string) int {
    var ret int = 0
    
    for _, el := range input {
        reading := ParseOasisHistory(el)
        ret += ParseOasisReading(reading)
    }
    
    return ret
}

func ParseOasisReading(reading []int) int {
    var ret int = 0
    var last []int = []int{reading[len(reading)-1]}
    var temp []int

    for {
        temp = make([]int, 0)
        for i := 0; i < len(reading)-1; i++ {
            temp = append(temp, reading[i+1] - reading[i])
        }

        if len(temp) == 0 {
            last = append(last, reading[0])
        }

        if IsAllZeroes(temp){
            break
        }
        last = append(last, temp[len(temp)-1])

        reading = temp
    }

    for i := len(last)-1; i >= 0; i-- {
        ret += last[i]
    }

    fmt.Println(ret)
    return ret 
}

func IsAllZeroes(arr []int) bool {
    for _, el := range arr {
        if el != 0 {
            return false
        }
    }
    return true
}

func ParseOasisHistory(line string) []int {
    var ret []int = make([]int, 0)
    for _, el := range strings.Split(line, " ") {
        num, _ := strconv.Atoi(el)
        ret = append(ret, num)
    }
    
    // Pt2 (reverse)
    var r []int = make([]int, 0)
    for i := len(ret)-1; i >= 0; i-- {
        r = append(r, ret[i]) 
    }
    // end pt2

    return r
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
