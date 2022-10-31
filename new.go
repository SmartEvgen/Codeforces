package main

import (
  "fmt"
  "strings"
)

func main() {
	var s string
	var k int = 0
	fmt.Scan(&s)
	if strings.Count(s, "a") > 0 {
		k++
		}
	if strings.Count(s, "b") > 0 {
		k++
		}
	if strings.Count(s, "c") > 0 {
		k++
		}
	if strings.Count(s, "d") > 0 {
		k++
		}
	if strings.Count(s, "e") > 0 {
		k++
		}
	if strings.Count(s, "f") > 0 {
		k++
		}
	if strings.Count(s, "g") > 0 {
		k++
		}
	if strings.Count(s, "h") > 0 {
		k++
		}
	if strings.Count(s, "i") > 0 {
		k++
		}
	if strings.Count(s, "j") > 0 {
		k++
		}
	if strings.Count(s, "k") > 0 {
		k++
		}
	if strings.Count(s, "l") > 0 {
		k++
		}
	if strings.Count(s, "m") > 0 {
		k++
		}
	if strings.Count(s, "n") > 0 {
		k++
		}
	if strings.Count(s, "o") > 0 {
		k++
		}
	if strings.Count(s, "p") > 0 {
		k++
		}
	if strings.Count(s, "q") > 0 {
		k++
		}
	if strings.Count(s, "r") > 0 {
		k++
		}
	if strings.Count(s, "s") > 0 {
		k++
		}
	if strings.Count(s, "t") > 0 {
		k++
		}
	if strings.Count(s, "u") > 0 {
		k++
		}
	if strings.Count(s, "v") > 0 {
		k++
		}
	if strings.Count(s, "w") > 0 {
		k++
		}
	if strings.Count(s, "x") > 0 {
		k++
		}
	if strings.Count(s, "y") > 0 {
		k++
		}
	if strings.Count(s, "z") > 0 {
		k++
		}
	if k%2 != 0 {
		fmt.Println("IGNORE HIM!")
		} else {
			fmt.Println("CHAT WITH HER!")
			}
}
