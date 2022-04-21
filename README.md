# GoADIFLog
Parse ADI file like [adif.org](https://adif.org)

Support cabrillo log file like https://www.cqwpx.com/cabrillo.htm

# Usage

```sh
go get -u github.com/Eminlin/GoADIFLog
```
   
```go 
import goadiflog "github.com/Eminlin/GoADIFLog"

func main(){
    logContent, err := goadiflog.Parse(path)
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Printf("%+v\n", logContent)
}
```

# Struct

```go
type CQLog struct {
	Call             string //被呼 呼号 *
	Mode             string //模式*
	Band             string //米波段*
	QSODate          string //QSO日期
	QSODateTimestamp int64  //QSO日期时间戳格式
	Frequency        string //频率
	StationCallsign  string //操作台呼号
	Operator         string //操作员*
	FileName         string //来源文件
}
``