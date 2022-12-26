package mod

type Provider interface {
	SpeedTest() (*Measure, error)
}

// Measure represents speed metrics
type Measure struct {
	Down string `json:"download"`
	Up   string `json:"upload"`
}
