package main

import (
	"fmt"
)

var EXPANSION_COEF = 1000000

func CosmicExpansion(input []string) int {
    var ret int = 0

    cosmos := ParseCosmicExpansion(input)
    galaxies := FindGalaxies(cosmos)
    expanded_rows := ExpandedRows(cosmos)
    expanded_cols := ExpandedCols(cosmos)

    for i := 0; i < len(galaxies)-1; i++ {
        for j := i+1; j < len(galaxies); j++ {
            ret += ManhattanDistance(galaxies[i], galaxies[j], expanded_rows, expanded_cols)
        }
    }

    return ret
}

func ManhattanDistance(a, b, rows, cols []int) int {
    count := 0
    max_i := a[0]
    min_i := b[0]
    max_j := a[1]
    min_j := b[1]

    if a[0] <= b[0] {
        max_i = b[0]
        min_i = a[0]
    }
    if a[1] <= b[1] {
        max_j = b[1]
        min_j = a[1]
    }

    for i := 0; i < len(rows); i++ {
        if rows[i] > min_i && rows[i] < max_i {
            count++
        }
    }
    for i := 0; i < len(cols); i++ {
        if cols[i] > min_j && cols[i] < max_j {
            count++
        }
    }

    return Abs(a[0] - b[0]) + Abs(a[1] - b[1]) + count * (EXPANSION_COEF-1)
}

func FindGalaxies(cosmos [][]string) [][]int {
    var ret [][]int = make([][]int, 0)
    for i := 0; i < len(cosmos); i++ {
        for j := 0; j < len(cosmos[i]); j++ {
            if cosmos[i][j] == "#" {
                ret = append(ret, []int{i, j})
            }
        }
    }
    return ret
}

func PrintCosmos(cosmos [][]string) {
    for i := range cosmos {
        for j := range cosmos[i] {
            fmt.Print(cosmos[i][j])
        }
        fmt.Println()
    }
}

func ParseCosmicExpansion(input []string) [][]string {
    var ret [][]string = make([][]string, 0)
   
    for i := 0; i < len(input); i++ {
        ret = append(ret, make([]string, 0))
        for j := 0; j < len(input[i]); j++ {
            ret[i] = append(ret[i], string(input[i][j]))
        }
    }
    return ret
}

func ExpandedCols(cosmos [][]string) []int {
    var ret []int = make([]int, 0)

    for j := 0; j < len(cosmos[0]); j++ {
        found := true
        for i := 0; i < len(cosmos); i++ {
            if cosmos[i][j] == "#" {
                found = false
                break
            }
        }
        if found {
            ret = append(ret, j)
        }
    }
    return ret
}

func ExpandedRows(cosmos [][]string) []int {
    var ret []int = make([]int, 0)

    for i := 0; i < len(cosmos); i++ {
        found := true
        for j := 0; j < len(cosmos[i]); j++ {
            if cosmos[i][j] == "#" {
                found = false
                break
            }
        }
        if found {
            ret = append(ret, i)
        }
    }
    return ret
}
