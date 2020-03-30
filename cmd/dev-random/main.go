package main

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/sync/errgroup"

	"github.com/rk23/symbiont/pkg/entropy"
	"github.com/rk23/symbiont/pkg/prng"
)

func main() {
	ctx := context.Background()
	input := make(chan []byte)
	output := make(chan []byte)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error { return entropy.Mem(input) })
	g.Go(func() error { return entropy.CPU(input) })
	g.Go(func() error { return prng.PRNG(ctx, input, output) })
	g.Go(func() error { return prng.Write(ctx, output) })

	err := g.Wait()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

}
