package ookla

import (
	"fmt"
	"github.com/showwin/speedtest-go/speedtest"
)

type Service interface {
	DownloadTest(savingMode bool) error
	UploadTest(savingMode bool) error
	DLSpeed() float64
	ULSpeed() float64
}

type serviceImpl struct {
	Server *speedtest.Server
}

func (s *serviceImpl) DownloadTest(savingMode bool) error {
	return s.Server.DownloadTest(savingMode)
}

func (s *serviceImpl) UploadTest(savingMode bool) error {
	return s.Server.UploadTest(savingMode)
}

func (s *serviceImpl) DLSpeed() float64 {
	return s.Server.DLSpeed
}

func (s *serviceImpl) ULSpeed() float64 {
	return s.Server.ULSpeed
}

// getServer gather all required information for running server
func getServer() (*speedtest.Server, error) {
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
