package function

import "strconv"

// AllNumConvertBIN
// @description 所有数转换二进制
// @param binNum
// @return string
func AllNumConvertBIN(binNum int64) string {
	var bin string = strconv.FormatInt(binNum, 2)
	return bin
}
