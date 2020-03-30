package main

import (
	"fmt"
	"os"

	"github.com/rk23/symbiont/pkg/entropy"
	"github.com/rk23/symbiont/pkg/prng"
)

func main() {

	input := make(chan []byte)
	output := make(chan []byte)

	go entropy.Mem(input)
	go entropy.CPU(input)
	go prng.PRNG(input, output)

	for {
		select {
		case input, more := <-output:
			if !more {
				os.Exit(0)
			}
			fmt.Println(string(input))
		}
	}

}
