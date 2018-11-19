package basics1

import (
	"math/rand"
	"time"
)

func init() {
	// seeding to get different values for each iteration of rand.Intn
	rand.Seed(time.Now().Unix())
}

func GenerateRandomString(sizeOfString int) string {

	byteArray := make([]byte, sizeOfString)

	for i := 0; i < sizeOfString; i++ {
		//65 to 90 because will return A to Z
		byteArray[i] = byte(GenerateRandInteger(65, 90))
	}

	return string(byteArray)
}

func GenerateRandInteger(min int, max int) int {
	return min + rand.Intn(max-min)
}
