package ookla

import (
	"fmt"
	"github.com/shchiv/speedtest/mod"
	"github.com/showwin/speedtest-go/speedtest"
)

type Service interface {
	DownloadTest(savingMode bool) error
	UploadTest(savingMode bool) error
	GetServer() (*speedtest.Server, error)
}

type serviceImpl struct {
}

func (s *serviceImpl) DownloadTest(savingMode bool) error {
	return s.Server.DownloadTest(savingMode)
}

func (s *serviceImpl) UploadTest(savingMode bool) error {
	return s.Server.UploadTest(savingMode)
}

// GetServer gather all required information for running server
func (s *serviceImpl) GetServer() (*speedtest.Server, error) {
	user, err := speedtest.FetchUserInfo()
	if err != nil {
		return nil, err
	}
	serverList, err := speedtest.FetchServers(user)
	if err != nil {
		return nil, err
	}
	targets, err := serverList.FindServer([]int{})
	if err != nil {
		return nil, err
	}
	if len(targets) >= 1 {
		// for simplicity selected first server
		return targets[0], nil
	}
	return nil, fmt.Errorf("can't find available servers")
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
	s, err := p.Service.GetServer()
	if err != nil {
		return nil, err
	}
	if err = p.Service.DownloadTest(false); err != nil {
		return nil, err
	}
	if err = p.Service.UploadTest(false); err != nil {
		return nil, err
	}
	return &mod.Measure{
		Down: fmt.Sprintf("%.2f Mbps", s.DLSpeed),
		Up:   fmt.Sprintf("%2.f Mbps", s.ULSpeed),
	}, nil
}
