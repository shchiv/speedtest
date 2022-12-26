package netflix

import "github.com/adhocore/fast"

type Service interface {
	Measure(noUp bool) (*fast.Fast, error)
}

type serviceImpl struct{}

func (s *serviceImpl) Measure(noUp bool) (*fast.Fast, error) {
	return fast.Measure(noUp)
}
