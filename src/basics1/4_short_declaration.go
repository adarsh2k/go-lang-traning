package basics1

import (
	"fmt"
	"math/cmplx"
)

var (
	a, b  = 1, 2
	z = cmplx.Sqrt(1-1i)
)

func SimplePrintValues() {
	k := 3
	fmt.Println(a,b,k)
	fmt.Printf("Type %T Value: %v \n", z, z)
}