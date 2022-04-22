package goadiflog

import (
	"strconv"
	"strings"
	"time"

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
		cabr.Operator = oprator
		c.dealSingle(lineLower, &cabr)
		calog = append(calog, cabr)
		// fmt.Printf("%+v\n", cabr)
	}
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
// QSO:   freq    mo   date         time   call            rst   exch     call            rst   exch     t
// QSO:   *****   **   yyyy-mm-dd   nnnn   *************   nnn   ******   *************   nnn   ******   n
// QSO:    3799   PH   1999-03-06   0711   HC8N            59    001      W1AW            59    001      0
// 0000 0 00001 1 11 1 1111122222 2 2222 3 3333333334444 4 444 4 455555 5 5555666666666 6 777 7 777777 8 8
// 123456789012345678901234567890123456789012345678901234567890123456789012345678901
func (c *cabrillo) dealSingle(line string, cabr *format.CQLog) {

	cabr.Frequency = strings.TrimSpace(cabr.Frequency)
	fre, _ := strconv.Atoi(cabr.Frequency)
	if fre > 0 {
		cabr.Band = c.freqToBand(fre)
	}

	cabr.Mode = strings.TrimSpace(strings.ToUpper(line[12-1 : 13+1]))

	cabr.QSODate = strings.ReplaceAll(strings.TrimSpace(line[15-1:24+1]), "-", "")
	t, _ := time.Parse("20060102", strings.TrimSpace(cabr.QSODate))
	cabr.QSODateTimestamp = t.Unix()
	if t.Unix() < 0 {
		cabr.QSODateTimestamp = 0
	}

	cabr.Call = strings.TrimSpace(line[55-1 : 68+1])
	//cabrillo log file missing exchange
	if len(strings.TrimSpace(cabr.Call)) <= 3 {
		cabr.Call = strings.TrimSpace(strings.ToUpper(line[49-1 : 55+1]))
	}

	cabr.FileName = c.filename
}

//freqToBand 频率对照米波段
func (c *cabrillo) freqToBand(freq int) string {
	if freq > 1800 && freq < 2000 {
		return "160M"
	}
	if freq > 3500 && freq < 4000 {
		return "80M"
	}
	if freq > 7000 && freq < 7300 {
		return "40M"
	}
	if freq > 10100 && freq < 10150 {
		return "30M"
	}
	if freq > 14000 && freq < 14350 {
		return "20M"
	}
	if freq > 18068 && freq < 18168 {
		return "17M"
	}
	if freq > 21000 && freq < 21450 {
		return "15M"
	}
	if freq > 24890 && freq < 24990 {
		return "12M"
	}
	if freq > 28000 && freq < 29700 {
		return "10M"
	}
	if freq > 50000 && freq < 54000 {
		return "6M"
	}
	if freq > 144000 && freq < 148000 {
		return "2M"
	}
	if freq > 222000 && freq < 225000 {
		return "1.25M"
	}
	if freq > 420000 && freq < 450000 {
		return "70CM"
	}
	if freq > 902000 && freq < 928000 {
		return "33CM"
	}
	if freq > 1200000 && freq < 1300000 {
		return "23CM"
	}
	if freq > 2400000 && freq < 2484000 {
		return "13CM"
	}
	return "UNKNOWN"
}
