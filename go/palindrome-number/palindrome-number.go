package main

import (
	"fmt"
	"strconv"
)

func toStr(x int)string{
    xStr := strconv.Itoa(x)
    return xStr
}

func reverse(str string) (result string) { 
    for _, v := range str { 
        result = string(v) + result 
    } 
    return
} 

func isPalindrome(x int) bool {
    xStr := toStr(x)
	xRev := reverse(xStr)
	// fmt.Println(xStr == xRev)
	return xStr == xRev
}

func main(){
	isPalindrome(121)
	fmt.Println(isPalindrome(121))
}