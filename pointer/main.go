package main

import (
	"fmt"
)

type Dimension struct {
	length int
	width int
	height int

}

func (d Dimension) Area () int {
	return d.length * d.width
}

func (d Dimension) Volume () int {
	return d.length * d.width * d.height
}

// func dimensions(length, width, height int) (area int, volume int) {
// 	area = length * width
// 	volume = length * width * height

// 	return 
// }

func main () {
	x, y := 5, 10
	n := &x

	fmt.Println(*n)
	*n = 50
	fmt.Println(x)
	t := &y
	*t = *t + 30
	fmt.Println(*t)
	
	
	// d := Dimension{10, 34, 2}
	
	// fmt.Printf("Area: %d\nVolume: %d\n ",d.Area(),d.Volume())
}