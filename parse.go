package goadiflog

import (
	"errors"
	"path"
	"strings"

	"github.com/Eminlin/GoADIFLog/format"
	"github.com/Eminlin/GoADIFLog/tools"
)

// Parse 解析入口
// fileDir:./testFile/xxx.adi or ./testFile/xxx.log
func Parse(fileDir string) ([]format.CQLog, error) {
	var rtn []format.CQLog
	filename := path.Base(fileDir)
	fileSuffix := path.Ext(filename)
	mode := getFileMode(fileSuffix)
	if mode == format.UnknowMode {
		return rtn, errors.New("unkown file format file:" + filename)
	}
	if mode == format.ADIFMode {
		line, err := tools.ReadADIFFileLine(fileDir)
		if err != nil {
			return rtn, err
		}
		return newAdfi(filename).parse(line), nil
	}
	if mode == format.CabrilloMode {
		line, err := tools.ReadCabrFileLine(fileDir)
		if err != nil {
			return rtn, err
		}
		return newCabrillo(filename).parse(line), nil
	}
	return rtn, nil
}

func ParseAdifFromString(adif string) ([]format.CQLog, error) {
	var rtn []format.CQLog
	line, err := tools.ReadAdifString(adif)
	if err != nil {
		return rtn, err
	}
	return newAdfi("").parse(line), nil
}

// getFileMode 检查文件后缀格式
func getFileMode(suffix string) int {
	switch strings.ToLower(suffix) {
	case ".adi":
		return format.ADIFMode
	case ".adif":
		return format.ADIFMode
	case ".log":
		return format.CabrilloMode
	default:
		return format.UnknowMode
	}
}
