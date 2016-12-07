package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	var a string
	fmt.Println("Enter string")
	reader := bufio.NewReader(os.Stdin)
    	a,_ = reader.ReadString('\n')
    	freqCount(a)
}

func freqCount(a string)map[string]int {
	field:= strings.Fields(a)
	c:= make(map[string]int)
    	for i := 0; i < len(field) ; i++ {
        	c[strings.ToLower(field[i])]++	
 	}
	fmt.Println(c)
	return c
}
