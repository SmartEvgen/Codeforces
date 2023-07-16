package main

import (
    "fmt"
    "strings"
)

func ToLower(s string) string {
	var lower string
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			lower += string(c + 32)
			} else {
				lower += string(c)
				}
		}
	return lower
	}

func main() {
	var a1 string 
	var a2 string 
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	var q1 string = ToLower(a1)
	var q2 string = ToLower(a2)
	fmt.Println(strings.Compare(q1, q2))
}
