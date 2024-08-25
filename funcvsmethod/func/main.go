package main

import (
	"fmt"
)

func dimensions(length, width, height int) (area int, volume int) {
	area = length * width
	volume = length * width * height

	return 
}

func main () {
	area, volume := dimensions( 10, 10, 20)
	fmt.Printf("Area: %d\nVolume: %d\n ",area,volume)
}