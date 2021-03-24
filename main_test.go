package main

import (
	"reflect"
	"testing"
)

func Test_getConfiguration(t *testing.T) {
	tests := []struct {
		name    string
		want    *Configuration
		wantErr bool
	}{
		{
			name: "test",
			want: &Configuration{
				Host: "0.0.0.0",
				Port: 8000,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getConfiguration()
			if (err != nil) != tt.wantErr {
				t.Errorf("getConfiguration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConfiguration() = %v, want %v", got, tt.want)
			}
		})
	}
}
