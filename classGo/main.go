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

*/
