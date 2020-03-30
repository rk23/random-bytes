package prng

import (
	"context"
	"crypto/sha512"
	"fmt"
)

func Write(ctx context.Context, input <-chan []byte) error {
	for {
		select {
		case o, more := <-input:
			if !more {
				return nil
			}

			// TODO: Dependency injection, allow for writing elsewhere
			fmt.Println(string(o))
		case <-ctx.Done():
			return ctx.Err()
		}
	}

}

func PRNG(ctx context.Context, src <-chan []byte, output chan<- []byte) error {
	s := sha512.New()
	pool := []byte{}
	count := 0

	for {
		select {
		case input, more := <-src:
			if !more {
				return nil
			}

			_, err := s.Write(pool)
			if err != nil {
				return err
			}
			_, err = s.Write(input)
			if err != nil {
				return err
			}

			pool = s.Sum(pool)

			s.Reset()
			count++

			if count > 2 {
				output <- pool
				pool = []byte{}
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}

}
