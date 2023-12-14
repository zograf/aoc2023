package main

import (
	"fmt"
	"slices"
)

func ParabolicReflectorDish(input []string) int {
    var ret int = 0
    var total int = 1000000000
    var values []int = make([]int, 0)
    var keys []string = make([]string, 0)

    dish := ParseReflectorDish(input)
    
    for i := 0; i < 950; i++ {
        FullCycle(&dish) 
    }

    for {
        val := FullCycle(&dish)
        key := CreateDishKey(dish)

        if slices.Contains(keys, key) {
            break
        }

        values = append(values, val)
        keys = append(keys, key)
    }

    ret = values[(total-950) % len(values)-1]
    return ret
}

func CreateDishKey(dish [][]string) string {
    var ret string = "" 
    for i := 0; i < len(dish); i++ {
        for j := 0; j < len(dish[i]); j++ {
            ret += dish[i][j]
        }
    }
    return ret
}

func FullCycle(dish *[][]string) int {
    TiltNorth(dish) 
    TiltWest(dish) 
    TiltSouth(dish) 
    TiltEast(dish) 
    return CalculateBeamLoad(*dish);
}

func CalculateBeamLoad(dish [][]string) int {
    var ret int = 0
    for i := 0; i < len(dish); i++ {
        for j := 0; j < len(dish[i]); j++ {
            if dish[i][j] == "O" {
                ret += len(dish) - i
            }
        }
    }
    return ret
}

func TiltNorth(dish *[][]string) {
    for i := 0; i < len(*dish); i++ {
        for j := 0; j < len((*dish)[i]); j++ {
            if (*dish)[i][j] == "O" {
                c := 0
                for k := i-1; k >= 0; k-- {
                    if (*dish)[k][j] == "." {
                        c++
                        continue
                    }
                    break
                }

                if (c == 0) {
                    continue
                }
                (*dish)[i-c][j] = "O"
                (*dish)[i][j] = "."
            }
        }
    }
}

func TiltSouth(dish *[][]string) {
    for i := len(*dish)-1; i >= 0; i-- {
        for j := len((*dish)[i])-1; j >= 0; j-- {
            if (*dish)[i][j] == "O" {
                c := 0
                for k := i+1; k < len(*dish); k++ {
                    if (*dish)[k][j] == "." {
                        c++
                        continue
                    }
                    break
                }

                if (c == 0) {
                    continue
                }
                (*dish)[i+c][j] = "O"
                (*dish)[i][j] = "."
            }
        }
    }
}

func TiltWest(dish *[][]string) {
    for i := 0; i < len(*dish); i++ {
        for j := 0; j < len((*dish)[i]); j++ {
            if (*dish)[i][j] == "O" {
                c := 0
                for k := j-1; k >= 0; k-- {
                    if (*dish)[i][k] == "." {
                        c++
                        continue
                    }
                    break
                }

                if (c == 0) {
                    continue
                }
                (*dish)[i][j-c] = "O"
                (*dish)[i][j] = "."
            }
        }
    }
}

func TiltEast(dish *[][]string) {
    for i := len(*dish)-1; i >= 0; i-- {
        for j := len((*dish)[i])-1; j >= 0; j-- {
            if (*dish)[i][j] == "O" {
                c := 0
                for k := j+1; k < len((*dish)[i]); k++ {
                    if (*dish)[i][k] == "." {
                        c++
                        continue
                    }
                    break
                }

                if (c == 0) {
                    continue
                }
                (*dish)[i][j+c] = "O"
                (*dish)[i][j] = "."
            }
        }
    }
}

func PrintReflectorDish(dish [][]string) {
    for _, el := range dish {
        for _, ch := range el {
            fmt.Print(ch)
        }
        fmt.Println()
    }
}

func ParseReflectorDish(input []string) [][]string {
    var ret [][]string = make([][]string, 0)
        
    for i := 0; i < len(input); i++ {
        temp := make([]string, 0)
        for j := 0; j < len(input[i]); j++ {
            temp = append(temp, string(input[i][j]))
        }
        ret = append(ret, temp)
    }

    return ret
}
