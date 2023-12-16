package main

import (
	"fmt"
	"slices"
)


type Beam struct {
    i           int
    j           int
    direction   string
}

type Tile struct {
    value       string
    energized   bool
}

func FloorIsLava(input []string) int {
    var ret int = 0
    
    starting_beams := GenerateStartingBeams(input)
    for _, beam := range starting_beams {
        grid := ParseMirrorGrid(input)
        SimulateBeams(&grid, beam)
        if CountEnergized(grid) > ret {
            ret = CountEnergized(grid)
        }
    }
    
    return ret
}

func GenerateStartingBeams(input []string) []Beam {
    var ret []Beam = make([]Beam, 0)
    for i := 0; i < len(input); i++ {
        ret = append(ret, Beam{i, -1, "r"}) 
        ret = append(ret, Beam{i, len(input[0]), "l"}) 
    }
    for j := 0; j < len(input[0]); j++ {
        ret = append(ret, Beam{-1, j, "d"}) 
        ret = append(ret, Beam{len(input), j, "u"}) 
    }
    return ret
}

func CountEnergized(grid [][]Tile) int {
    var ret int = 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid); j++ {
            if grid[i][j].energized {
                ret++
            }
        }
    }
    return ret
}

func SimulateBeams(grid *[][]Tile, start_beam Beam) {
    var beams []Beam = []Beam{start_beam}
    var seen_beams []Beam = []Beam{}

    for len(beams) != 0 { 
        beam := beams[0]
        beams = beams[1:]
        if slices.Contains(seen_beams, beam) {
            continue
        }
        seen_beams = append(seen_beams, beam)
        counter := 0

        for {
            if counter > 1000 {
                break
            }
            counter++
            if beam.direction == "r" {
                if beam.j+1 >= len((*grid)[0]) {
                    break
                }
                beam.j++
                b := &(*grid)[beam.i][beam.j]
                if b.value == "." {
                    b.energized = true
                } else if b.value == "-" {
                    b.energized = true
                } else if b.value == "|" {
                    b.energized = true
                    beam.direction = "u"
                    beams = append(beams, Beam{beam.i, beam.j, "d"})
                } else if b.value == "\\" {
                    b.energized = true
                    beam.direction = "d"
                } else if b.value == "/" {
                    b.energized = true
                    beam.direction = "u"
                }

            } else if beam.direction == "l" {
                if beam.j-1 < 0 {
                    break
                }
                beam.j--
                b := &(*grid)[beam.i][beam.j]
                if b.value == "." {
                    b.energized = true
                } else if b.value == "-" {
                    b.energized = true
                } else if b.value == "|" {
                    b.energized = true
                    beam.direction = "d"
                    beams = append(beams, Beam{beam.i, beam.j, "u"})
                } else if b.value == "\\" {
                    b.energized = true
                    beam.direction = "u"
                } else if b.value == "/" {
                    b.energized = true
                    beam.direction = "d"
                }

            } else if beam.direction == "d" {
                if beam.i+1 >= len(*grid) {
                    break
                }
                beam.i++
                b := &(*grid)[beam.i][beam.j]
                if b.value == "." {
                    b.energized = true
                } else if b.value == "-" {
                    b.energized = true
                    beam.direction = "l"
                    beams = append(beams, Beam{beam.i, beam.j, "r"})
                } else if b.value == "|" {
                    b.energized = true
                } else if b.value == "\\" {
                    b.energized = true
                    beam.direction = "r"
                } else if b.value == "/" {
                    b.energized = true
                    beam.direction = "l"
                }

            } else if beam.direction == "u" {
                if beam.i-1 < 0 {
                    break
                }
                beam.i--
                b := &(*grid)[beam.i][beam.j]
                if b.value == "." {
                    b.energized = true
                } else if b.value == "-" {
                    b.energized = true
                    beam.direction = "r"
                    beams = append(beams, Beam{beam.i, beam.j, "l"})
                } else if b.value == "|" {
                    b.energized = true
                } else if b.value == "\\" {
                    b.energized = true
                    beam.direction = "l"
                } else if b.value == "/" {
                    b.energized = true
                    beam.direction = "r"
                }
            }
        }
    }
}

func PrintGrid(grid [][]Tile) {
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j].energized {
                fmt.Print("#")
            } else {
                fmt.Print(grid[i][j].value)
            }
        }
        fmt.Println()
    }
}

func ParseMirrorGrid(input []string) [][]Tile {
    var ret [][]Tile = make([][]Tile, 0)

    for _, el := range input {
        temp := make([]Tile, 0)
        for _, ch := range el {
            temp = append(temp, Tile{string(ch), false})
        }
        ret = append(ret, temp)
    }

    return ret
}
