package utils

import (
	"crypto/md5"
	"encoding/hex"
)

/**
 *对一个字符串数据进行MD5哈希计算
 */
func MD5Hashstring(data string)string{
	md5Hash := md5.New()
	md5Hash.Write([]byte(data))
	bytes := md5Hash.Sum(nil)
	return hex.EncodeToString(bytes)//把脱敏的密码的md5值重新赋值为密码进行存储

}
