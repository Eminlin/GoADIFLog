package goadiflog

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Eminlin/GoADIFLog/format"
)

type adfi struct {
	fileName string
}

func newAdfi(filename string) *adfi {
	return &adfi{
		fileName: filename,
	}
}

func (a *adfi) parse(line []string) []format.CQLog {
	var cqlog []format.CQLog
	for _, v := range line {
		if v == "" {
			continue
		}
		if !strings.Contains(v, "<") {
			fmt.Printf("strings not contains < : %s \n", v)
			continue
		}
		compileRegex := regexp.MustCompile("<(.*?)>")
		matchArr := compileRegex.FindAllStringSubmatch(v, -1)
		//[[<CALL:6> CALL:6] [<QSO_DATE:8> QSO_DATE:8] [<TIME_ON:6> TIME_ON:6] [<TIME_OFF:6> TIME_OFF:6] [<BAND:3> BAND:3] [<STATION_CALLSIGN:5> STATION_CALLSIGN:5] [<FREQ:8> FREQ:8] [<CONTEST_ID:2> CONTEST_ID:2] [<FREQ_RX:8> FREQ_RX:8] [<MODE:3> MODE:3] [<RST_RCVD:2> RST_RCVD:2] [<RST_SENT:2> RST_SENT:2] [<TX_PWR:3> TX_PWR:3] [<OPERATOR:5> OPERATOR:5] [<CQZ:2> CQZ:2] [<STX:1> STX:1] [<APP_N1MM_POINTS:1> APP_N1MM_POINTS:1] [<APP_N1MM_RADIO_NR:1> APP_N1MM_RADIO_NR:1] [<APP_N1MM_CONTINENT:2> APP_N1MM_CONTINENT:2] [<APP_N1MM_RUN1RUN2:1> APP_N1MM_RUN1RUN2:1] [<APP_N1MM_RADIOINTERFACED:1> APP_N1MM_RADIOINTERFACED:1] [<APP_N1MM_ISORIGINAL:4> APP_N1MM_ISORIGINAL:4] [<APP_N1MM_NETBIOSNAME:15> APP_N1MM_NETBIOSNAME:15] [<APP_N1MM_ISRUNQSO:1> APP_N1MM_ISRUNQSO:1] [<APP_N1MM_ID:32> APP_N1MM_ID:32] [<APP_N1MM_CLAIMEDQSO:1> APP_N1MM_CLAIMEDQSO:1] [<EOR> EOR]]
		var adfi format.CQLog
		for _, single := range matchArr {
			singleLen := len(single)
			if singleLen == 0 {
				fmt.Printf("single is empty \n")
				continue
			}
			if singleLen == 1 {
				fmt.Printf("single is empty \n")
				continue
			}
			a.dealSingle(v, single[len(single)-1], &adfi)
		}
		cqlog = append(cqlog, adfi)
	}
	return cqlog
}

//dealSingle 处理单个匹配到的内容
func (a *adfi) dealSingle(line, match string, adfi *format.CQLog) {
	//CALL:6
	temp := strings.Split(match, ":")
	if len(temp) != 2 {
		return
	}
	lower := strings.ToLower(temp[0])
	if strings.Contains(lower, "call") {
		adfi.Call = a.getTagData(line, temp)
	}
	if strings.Contains(lower, "mode") {
		adfi.Mode = a.getTagData(line, temp)
	}
	if strings.Contains(lower, "band") {
		adfi.Band = a.getTagData(line, temp)
	}
	if strings.Contains(lower, "qso_date") {
		adfi.QSODate = a.getTagData(line, temp)
		t, _ := time.Parse("20060102", adfi.QSODate)
		adfi.QSODateTimestamp = t.Unix()
	}
	if strings.Contains(lower, "freq") {
		adfi.Frequency = a.getTagData(line, temp)
	}
	if strings.Contains(lower, "station") {
		adfi.StationCallsign = a.getTagData(line, temp)
	}
	if strings.Contains(lower, "operator") {
		adfi.Oprator = a.getTagData(line, temp)
	}
	adfi.FileName = a.fileName
}

//getTagData 获取adif格式tag对应的数据
func (a *adfi) getTagData(line string, matchArray []string) string {
	typeIndex := strings.Index(line, matchArray[0])
	//len(STATION_CALLSIGN) + len(":") + len(temp[1]) + len(">")
	start := typeIndex + len(matchArray[0]) + len(matchArray[1]) + 2
	len, err := strconv.Atoi(matchArray[1])
	if err != nil {
		fmt.Printf("strconv.Atoi error : %s \n", err)
	}
	end := start + len
	return line[start:end]
}
