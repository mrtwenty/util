package util

import (
	"strings"
	"time"
)

//时间戳
func Time()int64{
	return time.Now().Unix()
}

//时间戳格式化成时间格式
//参数一是字符格式，
//参数二是需要格式的时间戳
func Date(str string,t int64) string{
	maps:=map[string]string{
		"Y":"2006","m":"01","d":"02",
		"H":"15","i":"04","s":"05",
		"y":"06","n":"1","j":"2",
		"h":"08",
	}
	for k,v:=range maps{
		str=strings.Replace(str,k,v,1)
	}
	return time.Unix(t,0).Format(str)
}
