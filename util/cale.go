package util

import (
	"fmt"
	"strconv"
	"time"
)

// String2int 字符串转int
func String2int(str string) int {
	intValue, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return intValue
}

// SecondsToHms 将秒数转换为小时、分钟、秒的格式
func SecondsToHms(seconds int) string {
	t := time.Duration(seconds) * time.Second
	h := t / time.Hour
	t -= h * time.Hour
	m := t / time.Minute
	t -= m * time.Minute
	s := t / time.Second
	return fmt.Sprintf("%02d-%02d-%02d", h, m, s)
}

// CalcPercentage 计算percentage相对于total的百分比
func CalcPercentage(percentage int, total int) int {
	return int(float64(percentage) / float64(total) * 100)
}
