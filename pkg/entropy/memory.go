package entropy

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/elastic/go-sysinfo"
)

func Mem(out chan<- []byte) error {
	for {
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
