package entropy

import (
	"bytes"
	"context"
	"testing"
	"time"
)

func TestMem(t *testing.T) {
	ctx := context.Background()
	input := make(chan []byte)

	go Mem(ctx, input)

	// Needs to be longer than 3x max random sleep time to be able to sample at least two events
	// Sleeping in tests isn't ideal - tests should be fast. Given more time would create an entropy struct
	// with a sleep time. Main would set it to random while tests would set to zero
	time.Sleep(3 * time.Second)
	ctx.Done()

	x, y := <-input, <-input

	if bytes.Equal(x, y) {
		t.Errorf("failed to produce random value from Memory")
	}
}
