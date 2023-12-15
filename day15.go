package main

import (
	"strings"
)

type MirrorLens struct {
    value string
    lens int
}

func HASHAlgorithm(input []string) int {
    var ret int = 0
    var hashmap [][]MirrorLens = make([][]MirrorLens, 256)

    for i := 0; i < 256; i++ {
        hashmap[i] = make([]MirrorLens, 0)
    }

    for _, el := range strings.Split(input[0], ",") {
        if el[len(el)-2] == '=' {
            box := HASH(el[:len(el)-2])
            found := false
            for i := 0; i < len(hashmap[box]); i++ {
                if hashmap[box][i].value == el[:len(el)-2] {
                    hashmap[box][i].lens = int(el[len(el)-1] - '0')
                    found = true
                    break
                }
            }
            if (found) {
                continue
            }
            hashmap[box] = append(hashmap[box], MirrorLens{el[:len(el)-2], int(el[len(el)-1] - '0')})
        } else {
            box := HASH(el[:len(el)-1])
            found := false
            index := 0

            for i := 0; i < len(hashmap[box]); i++ {
                if hashmap[box][i].value == el[:len(el)-1] {
                    index = i
                    found = true
                    break
                }
            }
            if (found) {
               hashmap[box] = append(hashmap[box][:index], hashmap[box][index+1:]...) 
            }
        }
    }

    ret = CalculateLensStrength(hashmap)

    return ret
}

func CalculateLensStrength(hashmap [][]MirrorLens) int {
    var ret int = 0

    for i := 0; i < len(hashmap); i++ {
        for j := 0; j < len(hashmap[i]); j++ {
            ret += (i+1) * (j+1) * hashmap[i][j].lens
        }
    }

    return ret
}

func HASH(line string) int {
    var ret int = 0

    for _, ch := range line {
        ret += int(ch)
        ret *= 17
        ret %= 256
    }

    return ret
}
