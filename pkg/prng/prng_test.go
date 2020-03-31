package prng

import (
	"context"
	"testing"
)

func TestPRNG(t *testing.T) {
	// TODO: Could turn this into a table test where you try different byte strings,
	// check for distribution of output, let it run for more iterations,
	// and try other methods to test for failures.

	ctx := context.Background()
	input := make(chan []byte)

	// Need a buffered channel so it won't block while waiting to be read
	output := make(chan []byte, 10)

	go PRNG(ctx, input, output)

	input <- []byte("non-random string")
	input <- []byte("non-random string")
	input <- []byte("non-random string")
	close(input)

	uniquer := map[string]int{}

U:
	for {
		select {
		case o, more := <-output:
			if !more {
				break U
			}

			if len(o) != 64 {
				t.Errorf("returned unexpected amount of bytes: %d vs 64", len(o))
			}

			if count, ok := uniquer[string(o)]; !ok {
				uniquer[string(o)] = 0
			} else {
				if count > 1 {
					t.Errorf("PRNG failed to create unique values")
				}
			}
			uniquer[string(o)]++
		}
	}

}
