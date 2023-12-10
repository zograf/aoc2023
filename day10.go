package main

import "fmt"

func PipeMaze_Pt1(input []string) int {
    var ret int = 0
    var maze [][]string = ParsePipeMazeInput(input)
    var seen [][]int = make([][]int, 0)

    i, j := FindAnimal(maze)
    Traverse(i, j+1, &seen, maze)

    ret = len(seen) / 2
    return ret
}

func PipeMaze_Pt2(input []string) int {
    var ret int = 0
    var maze [][]string = ParsePipeMazeInput(input)
    var seen [][]int = make([][]int, 0)
    
    new_maze := ConvertMaze(maze)
    i, j := FindAnimal(new_maze)
    Traverse(i, j+1, &seen, new_maze)
    FloodFill(&new_maze, seen)

    for i := 0; i < len(new_maze); i+=2 {
        for j := 0; j < len(new_maze); j+=2 {
            if !IsContained(i, j, seen) && new_maze[i][j] != "+" {
                ret++
            }
        }
    }

    return ret
}

func ConvertMaze(maze [][]string) [][]string {
    var new_maze [][]string = make([][]string, (len(maze)+1) * 2) 

    for i := 0; i < len(new_maze); i++ {
        new_maze[i] = make([]string, (len(maze)+1) * 2)
        for j := 0; j < len(new_maze[i]); j++ {
            new_maze[i][j] = " "
        }
    }

    for i := 0; i < len(maze); i++ {
        for j := 0; j < len(maze[i]); j++ {
            if maze[i][j] == "S" {
                new_maze[i*2][j*2] = "S"
                new_maze[i*2][j*2+1] = "-"
                new_maze[i*2+1][j*2] = "|"
            } else if maze[i][j] == "|" {
                new_maze[i*2][j*2] = "|"
                new_maze[i*2+1][j*2] = "|"
            } else if maze[i][j] == "-" {
                new_maze[i*2][j*2] = "-"
                new_maze[i*2][j*2+1] = "-"
            } else if maze[i][j] == "J" {
                new_maze[i*2][j*2] = "J"
            } else if maze[i][j] == "F" {
                new_maze[i*2][j*2] = "F"
                new_maze[i*2][j*2+1] = "-"
                new_maze[i*2+1][j*2] = "|"
            } else if maze[i][j] == "L" {
                new_maze[i*2][j*2] = "L"
                new_maze[i*2][j*2+1] = "-"
            } else if maze[i][j] == "7" {
                new_maze[i*2][j*2] = "7"
                new_maze[i*2+1][j*2] = "|"
            } else if maze[i][j] == "." {
                new_maze[i*2][j*2] = "."
            }
        }
    }

    return new_maze
}

func FloodFill(maze *[][]string, loop [][]int) {
    var seen [][]int = make([][]int, 0)
    (*maze)[len(*maze)-1][len(*maze)-1] = "+"
    Fill(len(*maze)-1, len(*maze)-1, loop, &seen, maze)
}

func Fill(i, j int, loop [][]int, seen *[][]int, maze *[][]string) {
    (*seen) = append((*seen), []int{i, j})
    if !IsContained(i+1, j, *seen) && !IsContained(i+1, j, loop) && i+1 < len(*maze) {
        (*maze)[i+1][j] = "+"
        Fill(i+1, j, loop, seen, maze) 
    }
    if !IsContained(i-1, j, *seen) && !IsContained(i-1, j, loop) && i-1 >= 0 {
        (*maze)[i-1][j] = "+"
        Fill(i-1, j, loop, seen, maze) 
    }
    if !IsContained(i, j+1, *seen) && !IsContained(i, j+1, loop) && j+1 < len(*maze) {
        (*maze)[i][j+1] = "+"
        Fill(i, j+1, loop, seen, maze) 
    }
    if !IsContained(i, j-1, *seen) && !IsContained(i, j-1, loop) && j-1 >= 0 {
        (*maze)[i][j-1] = "+"
        Fill(i, j-1, loop, seen, maze) 
    }
}

func PrintMaze(maze [][]string) {
    for i := range maze {
        for j := range maze[i] {
            fmt.Print(maze[i][j])
        }
        fmt.Println()
    }
}

func Traverse(i, j int, seen *[][]int, maze [][]string) {
    (*seen) = append((*seen), []int{i, j})
    if maze[i][j] == "." || maze[i][j] == " " {
        return
    } else if maze[i][j] == "|" {
        if !IsContained(i-1, j, *seen) {
            Traverse(i-1, j, seen, maze)
        }
        if !IsContained(i+1, j, *seen) {
            Traverse(i+1, j, seen, maze)
        }
    } else if maze[i][j] == "-"{
        if !IsContained(i, j-1, *seen) {
            Traverse(i, j-1, seen, maze)
        }
        if !IsContained(i, j+1, *seen) {
            Traverse(i, j+1, seen, maze)
        }
    } else if maze[i][j] == "7" {
        if !IsContained(i, j-1, *seen) {
            Traverse(i, j-1, seen, maze)
        }
        if !IsContained(i+1, j, *seen) {
            Traverse(i+1, j, seen, maze)
        }
    } else if maze[i][j] == "F" {
        if !IsContained(i, j+1, *seen) {
            Traverse(i, j+1, seen, maze)
        }
        if !IsContained(i+1, j, *seen) {
            Traverse(i+1, j, seen, maze)
        }
    } else if maze[i][j] == "L" {
        if !IsContained(i, j+1, *seen) {
            Traverse(i, j+1, seen, maze)
        }
        if !IsContained(i-1, j, *seen) {
            Traverse(i-1, j, seen, maze)
        }
    } else if maze[i][j] == "J" {
        if !IsContained(i, j-1, *seen) {
            Traverse(i, j-1, seen, maze)
        }
        if !IsContained(i-1, j, *seen) {
            Traverse(i-1, j, seen, maze)
        }
    }
}

func IsContained(i, j int, list [][]int) bool {
    for _, arr := range list {
        if arr[0] == i && arr[1] == j {
            return true
        }
    }
    return false
}

func FindAnimal(maze [][]string) (int, int) {
    for i := range maze {
        for j, ch := range maze[i] {
            if ch == "S" {
                return i, j
            }
        }
    }
    return -1, -1
}


func ParsePipeMazeInput(input []string) [][]string {
    var ret [][]string = make([][]string, len(input))
    for i, line := range input { 
        ret[i] = make([]string, len(line))
        for j, ch := range line {
            ret[i][j] = string(ch)
        }
    }
    return ret
}
