package main

import (
	"fmt"
)

type Dimension struct {
	length int
	width  int
	height int
}

func (d *Dimension) Area() int {
	d.height = 8
	return d.length * d.width
}

func (d *Dimension) Volume() int {
	d.height = 6
	return d.length * d.width * d.height
}

// func dimensions(length, width, height int) (area int, volume int) {
// 	area = length * width
// 	volume = length * width * height

// 	return
// }

func main() {

	d := Dimension{10, 5, 6}

	fmt.Println("Area: ",d.Area())
	fmt.Println(d)
	fmt.Println("Volume: ",d.Volume())
	fmt.Println(d)

}
