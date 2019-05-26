package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsLetter("Alex"))      // true
	fmt.Println(IsLetter("alex"))      // true
	fmt.Println(IsLetter("alex.alex")) // true
	fmt.Println(IsLetter("123"))       // false
	/* 	fmt.Println(validationModule.CheckLocaleValue("en"))
	   	fmt.Println(validationModule.CheckLocaleValue("en11")) */
}
