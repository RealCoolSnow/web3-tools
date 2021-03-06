package str

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strconv"
)

// GetMD5Encode 返回一个32位md5加密后的字符串
func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func IsGBK(data []byte) bool {
	length := len(data)
	var i int = 0
	for i < length {
		if data[i] <= 0x7f {
			//编码0~127,只有一个字节的编码，兼容ASCII码
			i++
			continue
		} else {
			//大于127的使用双字节编码，落在gbk编码范围内的字符
			if data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i+1] >= 0x40 &&
				data[i+1] <= 0xfe &&
				data[i+1] != 0xf7 {
				i += 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func Int64ToStr(num int64) string {
	return strconv.FormatInt(num, 10)
}

func IntToStr(num int) string {
	return strconv.FormatInt(int64(num), 10)
}

func Float32ToStr(num float32) string {
	return fmt.Sprintf("%f", num)
}

func Float64ToStr(num float64) string {
	return fmt.Sprintf("%f", num)
}

func StrToFloat64(str string) (uint64, error) {
	bint, ok := new(big.Int).SetString(str, 0)
	if !ok {
		return 0, errors.New("SetString err")
	}
	return bint.Uint64(), nil
}
