/*
  2) Даны n каналов типа chan int. Надо написать функцию, которая смерджит все данные из этих каналов в один и вернет его.
Нужно , чтобы результат работы функции выглядел так:

for num := range joinChannels(a, b, c) {
    fmt.Println(num)
}

*/
package main

import (
	"fmt"
	"time"
)

const MULTIPLEXY_DELAY time.Duration = time.Second * 3

func main() {

	// For example we can use next bunch of channels:
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)
	d := make(chan int)

	go func() {
		a <- 1
		b <- 2
		c <- 3
		d <- 456
	}()

	for num := range joinChannels(a, b, c, d) {
		fmt.Println(num)
	}
}

func joinChannels(inputs ...<-chan int) <-chan int {
	c := make(chan int)
	for _, in := range inputs {
		go func(nchan <-chan int) {
			for {
				select {
				case s := <-nchan:
					c <- s
				default:
					time.After(MULTIPLEXY_DELAY)
				}
			}
		}(in)
	}
	return c
}
