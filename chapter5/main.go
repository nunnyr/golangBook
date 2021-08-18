package main

import "fmt"

// func main() {
// 	var x [5]int
// 	x[4] = 100
// 	fmt.Println(x)
// }

// func main() {
// 	var x [5]float64
// 	x[0] = 98
// 	x[1] = 93
// 	x[2] = 77
// 	x[3] = 82
// 	x[4] = 83

// 	var total float64 = 0
// 	for i := 0; i < len(x); i++ {
// 		total += x[i]
// 	}
// 	fmt.Print(total / (len(x)))
// }

// func main() {
// 	x := make([]float64, 5)
// }

// func main() {
// 	x := make(map[string]int)
// 	// var x map[string]int
// 	x["key"] = 10
// 	fmt.Println(x)
// }

//runtime errors happen when you run the program while compile-time errors
//happen when you try to compile the program

// func main() {
// 	//elements := make(map[string]string)
// 	// elements["H"] = "Hydrogen"
// 	elements := map[string]string {
// 		"H": "Hydrogen",

// 	}

// 	name, ok := elements["Un"]
// 	fmt.Println(name, ok)
// }

//write a program that finds the smallest number in this list
func main() {
	// final := nil
	// for i := 0; i < len(x); i++ {
	// }
	// for key, element := range x {
	// 	fmt.Println("Key:", key, "=>", "Element", element)
	// }
	var final int
	x := []int{
		48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17,
	}

	for i, value := range x {
		if i == 0 || value < final {
			final = value
		}
		fmt.Println(final)
	}

}
