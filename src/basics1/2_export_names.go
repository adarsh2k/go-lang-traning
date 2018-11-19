package basics1

import (
	"math"
)

// can be exported

func TestImport() float64 {
	return math.Pi
}

// can't be exported

func testImport() float64 {
	return math.Pi
}