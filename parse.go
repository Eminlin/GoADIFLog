package goadiflog

import (
	"fmt"
	"path"
	"regexp"
	"strings"

	"github.com/Eminlin/GoADIFLog/tools"
)

func Parse(filename string) {
	fmt.Printf("Start Parse File %s \n", path.Base(filename))
	line, err := tools.ReadFileLine(filename)
	if err != nil {
		panic(err)
	}
	for _, v := range line {
		if !strings.Contains(v, "<") {
			fmt.Printf("strings not contains < : %s \n", v)
			continue
		}
		compileRegex := regexp.MustCompile("<(.*?)>")
		matchArr := compileRegex.FindAllStringSubmatch(v, -1)
		//[[<CALL:6> CALL:6] [<QSO_DATE:8> QSO_DATE:8] [<TIME_ON:6> TIME_ON:6] [<TIME_OFF:6> TIME_OFF:6] [<BAND:3> BAND:3] [<STATION_CALLSIGN:5> STATION_CALLSIGN:5] [<FREQ:8> FREQ:8] [<CONTEST_ID:2> CONTEST_ID:2] [<FREQ_RX:8> FREQ_RX:8] [<MODE:3> MODE:3] [<RST_RCVD:2> RST_RCVD:2] [<RST_SENT:2> RST_SENT:2] [<TX_PWR:3> TX_PWR:3] [<OPERATOR:5> OPERATOR:5] [<CQZ:2> CQZ:2] [<STX:1> STX:1] [<APP_N1MM_POINTS:1> APP_N1MM_POINTS:1] [<APP_N1MM_RADIO_NR:1> APP_N1MM_RADIO_NR:1] [<APP_N1MM_CONTINENT:2> APP_N1MM_CONTINENT:2] [<APP_N1MM_RUN1RUN2:1> APP_N1MM_RUN1RUN2:1] [<APP_N1MM_RADIOINTERFACED:1> APP_N1MM_RADIOINTERFACED:1] [<APP_N1MM_ISORIGINAL:4> APP_N1MM_ISORIGINAL:4] [<APP_N1MM_NETBIOSNAME:15> APP_N1MM_NETBIOSNAME:15] [<APP_N1MM_ISRUNQSO:1> APP_N1MM_ISRUNQSO:1] [<APP_N1MM_ID:32> APP_N1MM_ID:32] [<APP_N1MM_CLAIMEDQSO:1> APP_N1MM_CLAIMEDQSO:1] [<EOR> EOR]]
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
			dealSingle(single[len(single)-1])
		}
	}
}

//dealSingle 处理单个匹配到的内容
func dealSingle(single string) {

}
