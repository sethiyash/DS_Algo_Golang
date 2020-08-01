// Go Program to check wheather the given string is an IP address or not

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	checkIPAddress("12.245.73.12")
	checkIPAddress("12.255.73.12")
	checkIPAddress("12.265.732.12")
	checkIPAddress("102.245.73.240")

}

func InvalidIPAddress(s string) bool {
	fmt.Printf("Given string %s is not an valid IP Address\n", s)
	return false
}

func validIPAddress(s string) bool {
	fmt.Printf("Given string is an IP Address - %s \n", s)
	return true
}

func checkIPAddress(s string) bool {
	s = strings.Trim(s, "")
	splitedStrings := strings.Split(s, ".")
	if len(splitedStrings) == 4 {
		for _, str := range splitedStrings {
			strint, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Error occured in converting string to int - ", err)
				return false
			}
			if strint < 0 || strint > 255 {
				return InvalidIPAddress(s)
			}
		}

	} else {
		return InvalidIPAddress(s)
	}
	return validIPAddress(s)
}
