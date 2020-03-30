package entropy

import (
	"fmt"
	"time"

	"github.com/elastic/go-sysinfo"
)

func CPU(out chan<- []byte) error {
	for {
		host, err := sysinfo.Host()
		if err != nil {
			panic(err)
		}
		cpu, err := host.CPUTime()

		out <- []byte(fmt.Sprintf("%d", cpu.Total()))
		time.Sleep(time.Duration(rand.Intn(999)) * time.Millisecond)
	}
}
