package tools

import (
	"bufio"
	"io"
	"os"
)

//ReadFileLine 逐行读取文件
func ReadFileLine(filename string) ([]string, error) {
	var temp []string
	fi, err := os.Open(filename)
	if err != nil {
		return temp, err
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	for {
		line, _, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return temp, err
		}
		temp = append(temp, string(line))
	}
	return temp, nil
}
