package mock

import (
	"fmt"
	"github.com/adhocore/fast"
	"github.com/shchiv/speedtest/mod"
)

// Provider mock
type Provider struct{}

// SpeedTest mock test
func (m *Provider) SpeedTest() (*mod.Measure, error) {
	return &mod.Measure{
		Down: "100 Mbps",
		Up:   "100 Mbps",
	}, nil
}

type NetflixService struct{}

func (ns *NetflixService) Measure(_ bool) (*fast.Fast, error) {
	return &fast.Fast{Up: "100", Down: "100"}, nil
}

type NetflixServiceRetErr struct{}

func (ns *NetflixServiceRetErr) Measure(_ bool) (*fast.Fast, error) {
	return nil, fmt.Errorf("failed")
}

type OoklaService struct {
}

func (os *OoklaService) DownloadTest(_ bool) error {
	return nil
}

func (os *OoklaService) UploadTest(_ bool) error {
	return nil
}

func (os *OoklaService) DLSpeed() float64 {
	return float64(100)
}

func (os *OoklaService) ULSpeed() float64 {
	return float64(100)
}

type OoklaServiceDownErr struct {
}

func (os *OoklaServiceDownErr) DownloadTest(_ bool) error {
	return fmt.Errorf("failed")
}

func (os *OoklaServiceDownErr) UploadTest(_ bool) error {
	return nil
}

func (os *OoklaServiceDownErr) DLSpeed() float64 {
	return float64(100)
}

func (os *OoklaServiceDownErr) ULSpeed() float64 {
	return float64(100)
}

type OoklaServiceUpErr struct {
}

func (os *OoklaServiceUpErr) DownloadTest(_ bool) error {
	return nil
}

func (os *OoklaServiceUpErr) UploadTest(_ bool) error {
	return fmt.Errorf("failed")
}

func (os *OoklaServiceUpErr) DLSpeed() float64 {
	return float64(100)
}

func (os *OoklaServiceUpErr) ULSpeed() float64 {
	return float64(100)
}
