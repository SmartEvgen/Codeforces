package main
import (
"fmt"
)
func main() {
	var x int
	var s string
	fmt.Scan(&x)
	for j := 1; j <= x; j++ {
		fmt.Scan(&s)
		if len(s) > 10 {
			fmt.Println(string(s[0]) + fmt.Sprint(len(s)-2) + string(s[len(s)-1]))
			} else if len(s) <= 10 {
				fmt.Println(s)
				}
		}
	}
