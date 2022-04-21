package goadiflog

import (
	"reflect"
	"testing"

	"github.com/Eminlin/GoADIFLog/format"
)

func TestParse(t *testing.T) {
	type args struct {
		fileDir string
	}
	tests := []struct {
		name    string
		args    args
		want    []format.CQLog
		wantErr bool
	}{
		// TODO: Add test cases.
		// {
		// 	name: "Version 1.0.9057.0",
		// 	args: args{
		// 		fileDir: "./testFile/B0CRA.adi",
		// 	},
		// },
		// {
		// 	name: "Version 1.0.9057.0 all",
		// 	args: args{
		// 		fileDir: "./testFile/B8CRA2021.ADI",
		// 	},
		// },
		{
			name: "HRD Logbook version 6.7.0.30",
			args: args{
				fileDir: "./testFile/2.0/20210505.ADI",
			},
		},
		// {
		// 	name: "START-OF-LOG: 2.0",
		// 	args: args{
		// 		fileDir: "./testFile/2.0/BG0ATE.log",
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.fileDir)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
