package speedtest

import (
	"github.com/shchiv/speedtest/mock"
	"github.com/shchiv/speedtest/mod"
	"github.com/shchiv/speedtest/netflix"
	"github.com/shchiv/speedtest/ookla"
	"reflect"
	"testing"
)

func TestSpeedTest(t *testing.T) {
	expMeasure := &mod.Measure{
		Down: "100 Mbps",
		Up:   "100 Mbps",
	}

	var tests = []struct {
		name     string
		provider mod.Provider
		want     *mod.Measure
		wantErr  bool
	}{
		{
			name:     "mock provider",
			provider: &mock.Provider{},
			want:     expMeasure,
			wantErr:  false,
		},
		{
			name:     "nil provider",
			provider: nil,
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "netflix provider",
			provider: &netflix.Provider{Service: &mock.NetflixService{}},
			want:     expMeasure,
			wantErr:  false,
		},
		{
			name:     "netflix provider err",
			provider: &netflix.Provider{Service: &mock.NetflixServiceRetErr{}},
			want:     nil,
			wantErr:  true,
		},
		{
			name: "ookla provider",
			provider: &ookla.Provider{
				Service: &mock.OoklaService{},
			},
			want:    expMeasure,
			wantErr: false,
		},
		{
			name: "ookla provider download err",
			provider: &ookla.Provider{
				Service: &mock.OoklaServiceDownErr{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ookla provider upload err",
			provider: &ookla.Provider{
				Service: &mock.OoklaServiceDownErr{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Start(tt.provider)
			if (err != nil) != tt.wantErr {
				t.Errorf("SpeedTest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SpeedTest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkOokla(b *testing.B) {
	provider := ookla.InitProvider()
	for i := 0; i < b.N; i++ {
		_, err := Start(provider)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNetflix(b *testing.B) {
	provider := netflix.InitProvider()
	for i := 0; i < b.N; i++ {
		_, err := Start(provider)
		if err != nil {
			b.Fatal(err)
		}
	}
}
