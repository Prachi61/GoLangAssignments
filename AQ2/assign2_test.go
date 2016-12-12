package main

import (
	"testing"
	"fmt"
	"reflect"
)

func Testassign2(t *testing.T) {
	expected := map[string]int{"this":2,"is":2,"the":1,"rock":1,"rocks":1}
  fmt.Println(expected)
  actual := freqCount("This is the Is rock rocks THIS")
  if !reflect.DeepEqual(expected,actual) {
    t.Errorf("Test failed")
  }

}