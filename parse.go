package goadiflog

import (
	"fmt"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/Eminlin/GoADIFLog/format"
	"github.com/Eminlin/GoADIFLog/tools"
)

func Parse(filename string) {
	fmt.Printf("Start Parse File %s \n", path.Base(filename))
	line, err := tools.ReadFileLine(filename)
	if err != nil {
		panic(err)
	}
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
		var adfi format.Adfi
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
			dealSingle(v, single[len(single)-1], &adfi)
		}
		fmt.Printf("%+v \n", adfi)
	}
}

//dealSingle 处理单个匹配到的内容
func dealSingle(line, match string, adfi *format.Adfi) {
	//CALL:6
	temp := strings.Split(match, ":")
	if len(temp) != 2 {
		return
	}
	lower := strings.ToLower(temp[0])
	if strings.Contains(lower, "call") {
		adfi.Call = getTagData(line, temp)
	}
	if strings.Contains(lower, "mode") {
		adfi.Mode = getTagData(line, temp)
	}
	if strings.Contains(lower, "band") {
		adfi.Band = getTagData(line, temp)
	}
	if strings.Contains(lower, "qso_date") {
		adfi.QSODate = getTagData(line, temp)
	}
	if strings.Contains(lower, "freq") {
		adfi.Frequency = getTagData(line, temp)
	}
	if strings.Contains(lower, "station") {
		adfi.StationCallsign = getTagData(line, temp)
	}
	if strings.Contains(lower, "operator") {
		adfi.Oprator = getTagData(line, temp)
	}
}

func getTagData(line string, matchArray []string) string {
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
