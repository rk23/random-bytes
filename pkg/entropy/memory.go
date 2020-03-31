package entropy

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/elastic/go-sysinfo"
)

// Mem puts random memory data onto the output channel
func Mem(ctx context.Context, out chan<- []byte) error {
	for {
		select {
		case <-ctx.Done():
			// TODO: Clean up work done here i.e. logging
			return ctx.Err()
		default:
			host, err := sysinfo.Host()
			if err != nil {
				return err
			}

			m, err := host.Memory()
			if err != nil {
				return err
			}

			out <- []byte(fmt.Sprintf("%d", m.Used))
			time.Sleep(time.Duration(rand.Intn(999)) * time.Millisecond)
		}
	}
}
