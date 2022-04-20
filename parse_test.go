package goadiflog

import "testing"

func TestParse(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		// {
		// 	name: "Version 1.0.9057.0",
		// 	args: args{
		// 		filename: "./testFile/B0CRA.adi",
		// 	},
		// },
		{
			name: "Version 1.0.9057.0 all",
			args: args{
				filename: "./testFile/B8CRA2021.ADI",
			},
		},
		// {
		// 	name: "HRD Logbook version 6.7.0.30",
		// 	args: args{
		// 		filename: "./testFile/2.0/20210505.ADI",
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Parse(tt.args.filename)
		})
	}
}
