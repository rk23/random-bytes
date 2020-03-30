package prng

import "crypto/sha512"

func PRNG(src <-chan []byte, out chan []byte) error {
	s := sha512.New()

	pool := []byte{}
	count := 0
	for {
		select {
		case input, more := <-src:
			if !more {
				return nil
			}

			s.Write(pool)
			s.Write(input)
			pool = s.Sum(pool)
			count++

			if count > 1 {
				out <- pool
				pool = []byte{}
			}
		}
	}

	return nil
}
