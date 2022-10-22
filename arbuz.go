package main

import "fmt"

func main() {
	var w int
	fmt.Scan(&w)
	if w%2 == 0 && w != 2 {
		fmt.Println("YES")
		} else if w%2 != 0 {
			fmt.Println("NO")
			} else if w == 2 {
			fmt.Println("NO")
			}
	}


