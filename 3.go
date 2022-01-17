/*
	Написать генератор случайных чисел
*/
package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	var i int64
	for i = 0; i < 10; i++ {
		fmt.Println(Rander(i + 1234564574756756))
	}
}

// Get random int64 similar as init vector int64
func Rander(b int64) int64 {
	big := big.NewInt(b)
	num, ok := rand.Int(rand.Reader, big)
	if ok != nil {
		panic(ok)
	}
	return num.Int64()
}
