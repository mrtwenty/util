package main

import (
	"fmt"
	"github.com/mrtwenty/util/util"
)

func main(){
	fmt.Println(util.CRC32("123456"))
	fmt.Println(util.MD5("123456"))
	fmt.Println(util.SHA1("123456"))
}
