package main

import "fmt"

func LavaMirrors(input []string) int {
    var ret int = 0;

    matrix_arr := ParseLavaMirrors(input)
    for i := 0; i < len(matrix_arr); i++ {
        ret += CalculateReflectionScore(matrix_arr[i])
    }
    
    return ret
}

func CalculateReflectionScore(matrix [][]string) int {
    for j := 0; j < len(matrix[0])-1; j++ {
        if IsColumnReflection(matrix, j, j+1) {
            return j+1
        } 
    }
    
    for i := 0; i < len(matrix)-1; i++ {
        if IsRowReflection(matrix, i, i+1) {
            return (i+1)*100
        }
    }

    return 0
}

func IsColumnReflection(matrix [][]string, j, k int) bool {
    var smudge int = 0
    for j >= 0 && k < len(matrix[0]) {
        if !IsColEquals(matrix, j, k, &smudge) {
            if smudge > 1 {
                return false
            } 
        }
        j--
        k++
    }
    if smudge == 1 {
        return true
    }
    return false
}

func IsRowReflection(matrix [][]string, i, k int) bool {
    var smudge int = 0
    for i >= 0 && k < len(matrix) {
        if !IsRowEquals(matrix, i, k, &smudge) {
            return false
        }
        i--
        k++
    }
    if smudge == 1 {
        return true
    }
    return false
}

func IsColEquals(matrix [][]string, j, k int, smudge *int) bool {
    for i := 0; i < len(matrix); i++ {
        if matrix[i][j] != matrix[i][k] {
            (*smudge)++
            if (*smudge) > 1 {
                return false
            }
        }
    }
    return true
}

func IsRowEquals(matrix [][]string, i, k int, smudge *int) bool {
    for j := 0; j < len(matrix[0]); j++ {
        if matrix[i][j] != matrix[k][j] {
            (*smudge)++
            if (*smudge) > 1 {
                return false
            }
        }
    }
    return true
}

func PrintLavaMirrors(m [][][]string) {
    for i := 0; i < len(m); i++ {
        for j := 0; j < len(m[i]); j++ {
            for k := 0; k < len(m[i][j]); k++ {
                fmt.Print(m[i][j][k])
            }
            fmt.Println()
        }
        fmt.Println()
    }
}

func ParseLavaMirrors(input []string) [][][]string {
    var ret [][][]string = make([][][]string, 0)
    var temp [][]string = make([][]string, 0)
    for _, el := range input {
        if len(el) == 0 {
            ret = append(ret, temp) 
            temp = make([][]string, 0)
        } else {
            row := make([]string, 0)
            for _, ch := range el {
                row = append(row, string(ch))
            }
            temp = append(temp, row)
        }
    }
    ret = append(ret, temp) 
    return ret
}
