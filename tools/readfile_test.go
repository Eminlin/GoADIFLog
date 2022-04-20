package tools

import (
	"reflect"
	"testing"
)

func TestReadFileLine(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				filename: "../testFile/B0CRA.adi",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFileLine(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFileLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFileLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
