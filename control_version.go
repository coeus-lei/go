package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

func main() {
	// v_initial := strconv.ltoa(2.2)
	var v_initial = "2.2"
	version := flag.String("version", "str_version", "version_id")
	// str_version = strconv.Itoa(version)
	var month = time.Now().Month()
	var digit_month = int(month)
	var str_digit_month = strconv.Itoa(digit_month)
	day := time.Now().Day()
	var str_day = strconv.Itoa(day)

	flag.Parse()
	fmt.Println(v_initial + "." + str_digit_month + "." + str_day + "." + *version)
}
