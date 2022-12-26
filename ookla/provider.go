package ookla

import (
	"fmt"
	"github.com/shchiv/speedtest/mod"
	"github.com/showwin/speedtest-go/speedtest"
	"log"
)

type Provider struct {
	Service Service
	Server  *speedtest.Server
}

func InitProvider() *Provider {
	srv, err := getServer()
	if err != nil {
		log.Fatal(err)
	}
	return &Provider{
		Service: &serviceImpl{Server: srv},
		Server:  srv,
	}
}

func (p *Provider) SpeedTest() (*mod.Measure, error) {
	if err := p.Service.DownloadTest(false); err != nil {
		return nil, err
	}
	if err := p.Service.UploadTest(false); err != nil {
		return nil, err
	}
	return &mod.Measure{
		Down: fmt.Sprintf("%.0f Mbps", p.Service.DLSpeed()),
		Up:   fmt.Sprintf("%.0f Mbps", p.Service.ULSpeed()),
	}, nil
}
