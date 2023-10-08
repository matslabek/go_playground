package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	resultSlice := make([][]uint8, 0)
	for i := 0; i < dy ; i++ {
		innerSlice := make([]uint8, 0)
		for j:= 0; j < dx; j++ {
			innerSlice = append(innerSlice, uint8((i^j)))
		}
		resultSlice = append(resultSlice, innerSlice)
	}
	return resultSlice
}

func main() {
	pic.Show(Pic)
}
