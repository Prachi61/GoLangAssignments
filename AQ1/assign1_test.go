package main

import (
	"testing"
	"reflect"
)

func Testassign1(t *testing.T) {
	expected := []int{2,3,2,1,1,1,0,0,0,1}
  actual := set("13412159020")
  if !reflect.DeepEqual(expected,actual) {
    t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
  }
}