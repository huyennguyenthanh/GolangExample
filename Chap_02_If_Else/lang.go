package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {

	// 3 ways to declare variables
	// (1) var one int
	// one = 1
	// (2) var one = 1
	// (3) one := 1
	var num = 1

	fmt.Printf("%s\n", strconv.Itoa(num))
	fmt.Print(num, num+1)

	if num == 1 {
		fmt.Println("Num is 1")
	}

	japanHello := "Konichiwa"
	vietnameHello := "Xin Chao"
	engHello := "Hello"

	lang := flag.String("lang", "en", "greeting language")
	flag.Parse()

	if *lang == "en" {
		fmt.Println(engHello)
	} else if *lang == "ja" {
		fmt.Println(japanHello)
	} else if *lang == "vi" {
		fmt.Println(vietnameHello)
	} else {
		fmt.Println("Not support this language.")
	}

}
