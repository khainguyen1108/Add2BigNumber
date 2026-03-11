package main

import (
	add2num "ADD2NUM/add2num"
	"fmt"
)

func main() {
	myBigNumber := add2num.MyBigNumber{}
	fmt.Println(myBigNumber.Sum("456", "789"))
}
