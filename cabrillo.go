package goadiflog

import (
	"fmt"
	"strings"

	"github.com/Eminlin/GoADIFLog/format"
)

type cabrillo struct {
	filename string
}

func newCabrillo(filename string) *cabrillo {
	return &cabrillo{
		filename: filename,
	}
}

func (c *cabrillo) parse(line []string) []format.CQLog {
	var calog []format.CQLog
	var stationCallsign, oprator string
	for _, v := range line {
		if v == "" {
			continue
		}
		var cabr format.CQLog
		lineLower := strings.ToLower(v)
		if !strings.HasPrefix(lineLower, "qso") && strings.Contains(lineLower, ":") {
			c.dealBaseInfo(v, lineLower, &stationCallsign, &oprator)
			continue
		}
		cabr.StationCallsign = stationCallsign
		cabr.Oprator = oprator
		c.dealSingle(lineLower, &cabr)
		calog = append(calog, cabr)
	}
	fmt.Printf("%+v\n", calog)
	return calog
}

func (c *cabrillo) dealBaseInfo(v, lineLower string, stationCallsign, oprator *string) {
	if strings.Contains(lineLower, "callsign") {
		lineInfo := strings.Split(v, ":")
		if len(lineInfo) == 2 {
			*stationCallsign = strings.TrimSpace(lineInfo[1])
		}
	}
	if strings.Contains(lineLower, "operator") {
		lineInfo := strings.Split(v, ":")
		if len(lineInfo) == 2 {
			*oprator = strings.TrimSpace(lineInfo[1])
		}
	}
}

// ---------------------info sent-------------           -------info rcvd--------
// QSO:  freq    mo   date         time   call            rst   exch     call            rst   exch     t
// QSO:  *****   **   yyyy-mm-dd   nnnn   *************   nnn   ******   *************   nnn   ******   n
// QSO:   3799   PH   1999-03-06   0711   HC8N            59    001      W1AW            59    001      0
// 00000 00001 1 11 1 1111122222 2 2222 3 3333333334444 4 444 4 455555 5 5555666666666 6 777 7 777777 8 8
// 123456789012345678901234567890123456789012345678901234567890123456789012345678901
func (c *cabrillo) dealSingle(line string, cabr *format.CQLog) {

	cabr.Frequency = line[6-1 : 5+1]

	cabr.Mode = line[12-1 : 13+1]

	cabr.QSODate = line[15-1 : 24+1]

	cabr.Call = line[56-1 : 68+1]

}
