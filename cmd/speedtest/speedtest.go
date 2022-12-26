package main

import (
	"fmt"
	"github.com/shchiv/speedtest/mod"
	"github.com/shchiv/speedtest/netflix"
	"github.com/shchiv/speedtest/ookla"
	"log"
)

func main() {
	netflixProvider := netflix.InitProvider()
	netflixSpeed, err := SpeedTest(netflixProvider)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Netflix Download: %v Upload: %v", netflixSpeed.Down, netflixSpeed.Up)

	ooklaProvider := ookla.InitProvider()
	ooklaSpeed, err := SpeedTest(ooklaProvider)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Ookla Download: %v Upload: %v", ooklaSpeed.Down, ooklaSpeed.Up)
}

// SpeedTest measures internet speed using provider data
func SpeedTest(p mod.Provider) (*mod.Measure, error) {
	if p == nil {
		return nil, fmt.Errorf("provider not implemented")
	}
	return p.SpeedTest()
}
