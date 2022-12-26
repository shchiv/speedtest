package speedtest

import (
	"fmt"
	"github.com/shchiv/speedtest/mod"
	"github.com/shchiv/speedtest/netflix"
	"github.com/shchiv/speedtest/ookla"
)

// Start speed test using provider data
func Start(p mod.Provider) (*mod.Measure, error) {
	if p == nil {
		return nil, fmt.Errorf("provider not implemented")
	}
	return p.SpeedTest()
}

// Netflix prepared provider to measure internet speed using fast.com
func Netflix() mod.Provider {
	return netflix.InitProvider()
}

// Ookla prepared provider to measure internet speed using speedtest.net
func Ookla() mod.Provider {
	return ookla.InitProvider()
}
