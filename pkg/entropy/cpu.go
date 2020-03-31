package entropy

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/elastic/go-sysinfo"
)

// CPU puts random CPU time data onto the output channel
func CPU(ctx context.Context, out chan<- []byte) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			host, err := sysinfo.Host()
			if err != nil {
				return err
			}

			cpu, err := host.CPUTime()
			if err != nil {
				return err
			}

			out <- []byte(fmt.Sprintf("%d", cpu.Total()))
			time.Sleep(time.Duration(rand.Intn(999)) * time.Millisecond)
		}
	}
}
