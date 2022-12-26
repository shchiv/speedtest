package netflix

import (
	"fmt"
	"github.com/adhocore/fast"
	"github.com/shchiv/speedtest/mod"
)

type Service interface {
	Measure(noUp bool) (*fast.Fast, error)
}

type serviceImpl struct{}

func (s *serviceImpl) Measure(noUp bool) (*fast.Fast, error) {
	return fast.Measure(noUp)
}

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
