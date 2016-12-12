package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string
	fmt.Println("Enter sequence")
    _, err := fmt.Scanf("%s", &a)
    if err != nil {
        fmt.Println(err)
    }	
    set(a)
}

func set(a  string)[]int {
	c:= make([]int, 10)
    	for i := 0; i < len(a) ; i++ {
    		s, e := strconv.Atoi(a[i:i+1])
    		if e != nil {
        		fmt.Println(e)
    		}
		c[s]++;	
 	}
	fmt.Println(c);
	return c
}
