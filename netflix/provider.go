package netflix

import (
	"fmt"
	"github.com/shchiv/speedtest/mod"
)

type Provider struct {
	Service Service
}

func InitProvider() *Provider {
	return &Provider{
		Service: &serviceImpl{},
	}
}

func (p *Provider) SpeedTest() (*mod.Measure, error) {
	ms, err := p.Service.Measure(false)
	if err != nil {
		return nil, err
	}
	return &mod.Measure{
		Down: fmt.Sprintf("%s Mbps", ms.Down),
		Up:   fmt.Sprintf("%s Mbps", ms.Up),
	}, nil
}
