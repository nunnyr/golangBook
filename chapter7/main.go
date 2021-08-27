package main

import "fmt"

func main() {
	fmt.Println("test")
}

/*
a struct is a type that contains named fields.
for example

type Circle struct {
	x float64
	y float64
	r float64
}

the type keyword introduces a new type. it's followed by the name of the type (Circle),
the keyword struct to indicate that we are defining a struct type, and a list of fields inside of curly braces.

Fields are like a set of grouped variables. Each field has a name and at ype and is stored adjacent to
the other fields in the struct. Like with functions, we can collapse fields that have the same type.

type Circle struct {
	x,y,r float64
}

we create an instance of our new Circle type in a variety of ways
              var c Circle
the default of that variable is set to zero. for a struct zero means each of the fields is set to their corresponding zero value
we can also use the new function...... c := new(Circle)
-this allocates memory for all the fields, set each of them ot their zero value, and returns a pointer to the struct (*Circle). Pointers are often used with structs so that functions can modifiy their contents.
-using new in this way is somewhat uncommon. more typically, we want to give each of the fields an initial value. we can do this in two ways.




*/

/*
what is the difference between a method and a function?
-a method is a function that has a defined receiver.
-the receiver is like a parameter - it has a name and a type - but by creating the function in this way, it allos us to call the function
	using the . operator
*/
package main

import ("fmt"; "math")

func distance(x1, y1, x2, y2 float64) float64 {
	

}