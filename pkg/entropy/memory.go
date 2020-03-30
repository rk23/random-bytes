package entropy

import (
	"fmt"
	"time"

	"github.com/elastic/go-sysinfo"
)

func Mem(out chan<- []byte) error {
	for {
		host, err := sysinfo.Host()
		if err != nil {
			panic(err)
		}
		m, err := host.Memory()
		out <- []byte(fmt.Sprintf("%d", m.Used))
		time.Sleep(time.Duration(rand.Intn(999)) * time.Millisecond)
	}
}
