package main

import (
	"fmt"
	"math"
)

// // Prints a 5x5 square
// func main() {
// 	for row := 1; row <= 5; row++ {
// 		for i := 1; i <= 5; i++ {
// 			fmt.Printf("#")
// 		}
// 		fmt.Printf("\n")
// 	}
// }

// // Prints a half square (base is 5 at top)
// func main() {
// 	for row := 1; row <= 5; row++ {
// 		for i := 1; i <= 6-row; i++ {
// 			fmt.Printf("#")
// 		}
// 		fmt.Printf("\n")
// 	}
// }

// Prints a sideways triangle
func main() {
	for row := 1; row <= 7; row++ {
		for i := 1; i <= 4-int(math.Abs(float64(4-row))); i++ {
			fmt.Printf("#")
		}
		fmt.Printf("\n")
	}
}
