package main

import "fmt"

//Go only has one type of for loop.
// func main() {
// 	i := 1
// 	for i <= 10 {
// 		fmt.Println(i)
// 		i = i + 1
// 	}
// }

// func main() {
// 	for i := 1; i <= 10; i++ {
// 		if i%2 == 0 {
// 			fmt.Println(i, "even")
// 		} else {
// 			fmt.Println(i, "odd")
// 		}
// 	}
// }

//a switch statement starts with the keyword switch followed by an expression (n this case i)
//and then a series of cases. The value of expression is compared to the expression following each
//case keyword.

// func main() {
// 	for i := 1; i <= 100; i++ {
// 		if i%3 == 0 {
// 			fmt.Println(i)
// 		}
// 	}
// }

func main() {
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		}
	}
}
