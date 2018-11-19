package main

import (
	"basics1"
	"fmt"
)

func main() {
	//basics 1 1_packages.go
	fmt.Println(basics1.GenerateRandomString(10))

	//basics 1 2_export_names.go
	fmt.Println(basics1.TestImport())

	//basics 1 3_named_return.go
	fmt.Println(basics1.SplitNumbers(28))

	//basics 1 4_short_declaration.go
	basics1.SimplePrintValues()

	//basics 1 5_looping_and_conditional_fns.go
	basics1.ForLoopInGo()
	basics1.WhileTypeLoopInGo()
	fmt.Println(basics1.Sqrt(2))
}