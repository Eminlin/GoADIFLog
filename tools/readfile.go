package tools

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// ReadADIFFileLine 逐行读取ADIF文件
func ReadADIFFileLine(filename string) ([]string, error) {
	var temp []string
	fi, err := os.Open(filename)
	if err != nil {
		return temp, err
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	findEOR := ""
	for {
		line, _, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return temp, err
		}
		if strings.Contains(strings.ToLower(string(line)), "eor") {
			findEOR += strings.Replace(string(line), "\n", "", -1)
			temp = append(temp, findEOR)
			findEOR = ""
		} else {
			findEOR += strings.Replace(string(line), "\n", "", -1)
		}

	}
	return temp, nil
}

func ReadAdifString(adif string) ([]string, error) {
	var temp []string
	br := bufio.NewReader(strings.NewReader(adif))
	findEOR := ""
	for {
		line, _, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return temp, err
		}
		if strings.Contains(strings.ToLower(string(line)), "eor") {
			findEOR += strings.Replace(string(line), "\n", "", -1)
			temp = append(temp, findEOR)
			findEOR = ""
		} else {
			findEOR += strings.Replace(string(line), "\n", "", -1)
		}

	}
	return temp, nil
}

// ReadCabrFileLine 逐行读取Cabr文件
func ReadCabrFileLine(filename string) ([]string, error) {
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
