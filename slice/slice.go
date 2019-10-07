package main

import (
	//"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	
	for i := 0; i < dy; i++ {
		pic[i] = make([]uint8, dx)
	}
	
	for y, row := range pic {
		for x := range row {
			row[x] = uint8((x+y)/2)
		}
	}
			return pic
}

func main() {
	pic.Show(Pic)
}