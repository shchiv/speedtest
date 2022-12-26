package speedtest

import (
	"fmt"
	"github.com/shchiv/speedtest/mod"
	"github.com/shchiv/speedtest/netflix"
	"github.com/shchiv/speedtest/ookla"
)

// SpeedTest measures internet speed using provider data
func SpeedTest(p mod.Provider) (*mod.Measure, error) {
	if p == nil {
		return nil, fmt.Errorf("provider not implemented")
	}
	return p.SpeedTest()
}

// Netflix measures internet speed using fast.com
func Netflix() (*mod.Measure, error) {
	netflixProvider := netflix.InitProvider()
	return SpeedTest(netflixProvider)
}

// Ookla measures internet speed using speedtest.net
func Ookla() (*mod.Measure, error) {
	ooklaProvider := ookla.InitProvider()
	return SpeedTest(ooklaProvider)
}
