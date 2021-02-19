package util

import "strings"

func ReplaceSQLStr(str string) string{
	str = strings.Replace(str, "'","\\'",-1)
	return str
}