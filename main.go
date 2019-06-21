package main

import (
	"fmt"
	"github.com/mrtwenty/util/v2/util"
)

func main(){
	fmt.Println(util.CRC32("123456"))
	fmt.Println(util.MD5("123456"))
	fmt.Println(util.SHA1("123456"))
	fmt.Println(util.PasswordHash("123456"))
	fmt.Println(util.Date("Y-m-d H:i:s",util.Time()))
}
