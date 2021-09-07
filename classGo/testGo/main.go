// package main

// func main() {
// 	printState()
// }

package main

import "fmt"

// type book string

// func (b book) printTitle() {
// 	fmt.Println(b)
// }

// func main() {
// 	var b book = "Harry Potter"
// 	b.printTitle()
// }

type music string

func (m music) songTitle() {
	fmt.Println(m)
}

func main() {
	var m music = "Sound of Music"
	m.songTitle()
}
