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
		// 	name: "正常识别 1.0",
		// 	args: args{
		// 		fileDir: "./testFile/B8CRA2021.ADI",
		// 	},
		// },

		// {
		// 	name: "测试没有station_callsign，读文件名",
		// 	args: args{
		// 		fileDir: "./testFile/3.0/2021 B2CRA ALL LOG.ADI",
		// 	},
		// },

		// {
		// 	name: "呼号识别错误：N:5>B9",
		// 	args: args{
		// 		fileDir: "./testFile/3.0/青海B9CRA_20210505_SELECTED_QSO.adi",
		// 	},
		// },
		// {
		// 	name: "带D识别 ",
		// 	args: args{
		// 		fileDir: "./testFile/3.0/2021 B2CRA ALL LOG.ADI",
		// 	},
		// },
		{
			name: "识别2.0 ",
			args: args{
				fileDir: "./testFile/2.0/20210505.ADI",
			},
		},
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
