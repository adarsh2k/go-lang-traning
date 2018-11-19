package basics1

import "fmt"

func Sqrt(numberToFindRoot float64) (i float64) {
	precision := float64(0.00001)
	for i = float64(1); i*i < numberToFindRoot; i++ {}
	for i--; i*i < numberToFindRoot; i += precision {}
	return
}

func ForLoopInGo() {
	for i := 0; i < 20 ; i++ {
		if a := i / 2 ; a < 10 {
			fmt.Printf("%v ",i * 2)
		}
		fmt.Printf("%v ",i)
	}
	fmt.Printf("\n")
}

func WhileTypeLoopInGo()  {
	i := 0
	for i < 2 {
		fmt.Printf("%v ", i+2)
		i++
	}
	fmt.Printf("\n")
}
