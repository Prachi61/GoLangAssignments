package main

import (
	"fmt"
	"strconv"
)

type Month int

const (
January Month = 1
February
March
April
May
June 
July
August
September
October
November
December
)

func (i Month) String() string {
return "Month " + strconv.Itoa(int(i))
}

func main() {
	fmt.Printf("%s",Month(January))
	//for i:=1;i<13;i++ {
	//	fmt.Printf("%s",Month[i])
	//}
}