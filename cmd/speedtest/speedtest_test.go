package main

import (
	"github.com/shchiv/speedtest/mod"
	"reflect"
	"testing"
)

func TestSpeedTest(t *testing.T) {
	type args struct {
		p mod.Provider
	}
	tests := []struct {
		name    string
		args    args
		want    *mod.Measure
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SpeedTest(tt.args.p)
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
