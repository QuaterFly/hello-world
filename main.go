package main

import (
	"fmt"
	"getostype"
	_ "getostype"
)


func main(){
	fmt.Println("Hello World!/n")
	fmt.Println("OS type:" + getostype.GetOSType())
}
	