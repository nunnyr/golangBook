package main

//functions can take parameters in this order: name type, name type etc
// func average(xs []float64) float64 {
// 	total := 0.0
// 	//go compiler won't let us run things we define but never use
// 	//using for _v, we are saying that we don't need the iterator variable
// 	//bc we don't use it in our for block
// 	for _, v := range xs {
// 		total += v
// 	}
// 	return (total / float64(len(xs)))
// 	//	panic("Not implemented")

// }

// func main() {
// 	xs := []float64{98, 93, 77, 82, 83}
// 	fmt.Println(average(xs))
// }

//variables must be passed to functions
//functions don't have access to anything in the calling function unless it's
//passed in explicitly.

//functions form call stacks....

// func f() (int, int) {
// 	return 5, 6
// }

// func main() {
// 	x, y := f()
// }

//Variadic functions
//-to be able to called with multiple integers, ths is known as a variadic parameter.

// func add(args ...int) int {
// 	//more code
// }
// func main() {
// 	add := func(x, y int) int {
// 		return x + y
// 	}
// 	fmt.Println(add(1, 1))

// 	// xs := []int{1, 2, 3}
// 	// fmt.Println(add(xs...))
// }

//int64 cannt mix types

/*
1.sum is a function that takes a slice of numbers and adds them together
	what would its function look like in Go?
		func sum([xs [] int) int
		    function^ with the name(xs) and the type and outside is the return type

2.write a function that takes an integer and halves it and returns true if it was even or false
	if it was odd. For example, half(1) should return (0, false) and
	half(2) should return (1, true)

*/
//parameters and the return type are known as the function's signature

// func main(num int) (int, bool) {
// 	final := (num / 2)
// 	var result bool

// 	if final%2 == 0 {
// 		result = true
// 	} else {
// 		result = false
// 	}

// 	return (final + result)
// }

// main(4)

func main() {

	// fmt.Println(half(120))
	// fmt.Println(greatestNum(10, 20, 24, 5))
	// nextOdd := makeOddGenerator()
	// fmt.Println(nextOdd())
	// fmt.Println(nextOdd())
	// fmt.Println(nextOdd())
	// fmt.Println(nextOdd())
	// fmt.Println(nextOdd())
	// x := 1.5
	// square(&x)

	// fmt.Println(x)

}

func half(num int) (int, bool) {
	final := (num / 2)
	var result bool

	if final%2 == 0 {
		result = true
	} else {
		result = false
	}

	return final, result

}

func greatestNum(num ...int) int {
	// final := 0
	var final int
	for _, v := range num {
		if v > final {
			final = v
		}
	}
	return final
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func square(x *float64) {
	*x = *x * *x

}

/*
what are defer, panic, and recover?
-defer is a special statement that schedules a function call to be run after the function completes.
-panic will immediately stop the execution of the function. and you have to pair it with defer.

-panic generally indicates a programmer error or an exceptional condition that there's no easy way to recover from (hence the word panic)
-recover stops the panic and returns the value that was passed to the call to panic.


how do you get the memory address of a variable?
-we use the & operator to find the address of a variable
-an asterisk is used to dereference pointer variables. derefencing a pointer gives us accesss to the value the pointer points to.

how do you assign a value to a pointer?
-define a pointer variable like *x = 5
-using the & you can access the value


how do you create a new pointer?
-new takes a type as an argument, allocates enough memory to fit a value of that type, and returns a pointer to it.
- new function:x = new(int)

*/

func swap(x, y *int) {
	*x, *y = *y, *x
}
