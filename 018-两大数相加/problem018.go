package main

import (
	"strconv"
	"unsafe"
)

func mySum(x,y string) string {
	if len(x) < len(y) {
		x,y = y,x
	}
	difference := len(x) - len(y)
	carry := 0
	numBytes := make([]uint8,len(x)+1)
	for i := len(x)-1;i>=0;i-- {
		a,_ := strconv.Atoi(x[i:i+1])
		b := 0
		if i >= difference{
			b,_ = strconv.Atoi(y[i-difference:i+1-difference])
		}
		c := a + b + carry
		if c >= 10 {
			c = c - 10
			carry = 1
		}else {
			carry = 0
		}
		num := strconv.Itoa(c)[0]
		numBytes[i+1] = num
	}
	numBytes[0] = uint8('1')
	if carry == 0 {
		numBytes = numBytes[1:]
		return *(*string)(unsafe.Pointer(&numBytes))
	}
	return *(*string)(unsafe.Pointer(&numBytes))
}

func main()  {
	println(mySum("123456789","987654321"))
}