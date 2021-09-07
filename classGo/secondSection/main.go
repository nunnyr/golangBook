package main

import "fmt"

func main() {
	fmt.Println("Hi there!")
}

/*
how do we run the code in our project?
go run main.go : go run compiles and executes
go command is our portal to running go
gives us the ability to compile and run
go build: just compiles it but does not execute it. ls
go install and go get : are used to compile and install apackage and downaloads the raw source code of someoen elses package
*/

/*
what does package main mean
- the first line must declare which file it belongs to
- package main
- a package can have many related files inside of it
- there are 2 different types of packages: we have an executable
type which generates a file that we can run. reusable: code used as 'helpers'


what does 'import "fmt" ' mean
-fmt is short for format and is a library. it is used to print out information to the terminal

what is func?
-func is short for function,

how are the files organized?
-at the very top we are having our package declaration.
-underneath that you will list the other imports that you need or create custom packages
-after you get to the body of the file. this is where your logic lives.
-


-at the very top we are having our package declaration.
-underneath that you will list the other imports that you need or create custom packages
-after you get to the body of the file. this is where your logic lives.
-

1. what is the purpose of a package in Go? To group together code with a similar purpose.
2. what is the special name we use for a package to tell Go that we want it to be turned into
a file that can be executed? the package would be "main"
3. the one requirement of packages named "main" is that we....? define a function named
"main" which is ran automatically when the program runs.
4. why do we use "import" statements? to give our package access to code written in another package




*/
