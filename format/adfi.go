package format

var (
	UnknowMode   = 0
	ADIFMode     = 1
	CabrilloMode = 2
)

// CQLog 从日志获取主要的信息
type CQLog struct {
	Call             string //被呼 呼号 *
	Mode             string //模式*
	Band             string //米波段*
	QSODate          string //QSO日期
	TimeOn           string
	TimeOff          string
	QSODateTimestamp int64  //QSO日期时间戳格式
	Frequency        string //频率
	StationCallsign  string //操作台呼号
	Operator         string //操作员*
	FileName         string //来源文件
	QslRcvd          string
}
