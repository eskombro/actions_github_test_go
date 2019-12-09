package main

import (
	"testing"
)

func TestCheckAge(t *testing.T) {
	if !checkAge(person{Age: 19}.Age) {
		t.Errorf("Error in checkAge")
	}
	if !checkAge(person{Age: 18}.Age) {
		t.Errorf("Error in checkAge")
	}
	if checkAge(person{Age: 17}.Age) {
		t.Errorf("Error in checkAge")
	}
}

func TestSellBeer(t *testing.T) {
	resultStr := sellBeer(person{Name: "Sam", Age: 19})
	expectedStr := "Sam can buy beer"

	if expectedStr != resultStr {
		t.Errorf("Error in sellBeer.\nExpected:\n%s\nGot\n%s", expectedStr, resultStr)
	}

	resultStr = sellBeer(person{Name: "Sam", Age: 17})
	expectedStr = "Sam CAN'T buy beer"

	if expectedStr != resultStr {
		t.Errorf("Error in sellBeer.\nExpected:\n%s\nGot\n%s", expectedStr, resultStr)
	}
}
