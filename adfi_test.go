package goadiflog

import "testing"

func Test_adfi_getStationCallFromFileName(t *testing.T) {
	tests := []struct {
		name string
		a    *adfi
		want string
	}{
		// TODO: Add test cases.
		{
			name: "01",
			a: &adfi{
				fileName: "temp/B8CRA2021活动日志汇总",
			},
			want: "B8CRA",
		},
		{
			name: "02",
			a: &adfi{
				fileName: "temp/2021_B7CRA_02A.adi",
			},
			want: "B7CRA",
		},
		{
			name: "03",
			a: &adfi{
				fileName: "temp/DXPED-HF-MIXED_2021_STN1@B7CRA.adi",
			},
			want: "B7CRA",
		},
		{
			name: "04",
			a: &adfi{
				fileName: "temp/2021 B2CRA ALL LOG.adi",
			},
			want: "B2CRA",
		},
		{
			name: "05",
			a: &adfi{
				fileName: "temp/B5CRA (OP BA5HAM) 20210505.adi",
			},
			want: "B5CRA",
		},
		{
			name: "06",
			a: &adfi{
				fileName: "temp/wsjtx_log_B7CRA_BA7NO.adi",
			},
			want: "B7CRA",
		},
		{
			name: "07",
			a: &adfi{
				fileName: "temp/BG4HYK-20-40-CW-B5CRA-2.adi",
			},
			want: "B5CRA",
		},
		{
			name: "08",
			a: &adfi{
				fileName: "temp/B4CRA-0504-BA4WI、BH4RRG、BI4VIP、BH8PHG、BD4RCC-4.adi",
			},
			want: "B4CRA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.getStationCallFromFileName(); got != tt.want {
				t.Errorf("adfi.getStationCallFromFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}
