package entropy

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/elastic/go-sysinfo"
)

// CPU puts random CPU time data onto the output channel
func CPU(out chan<- []byte) error {
	for {
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
