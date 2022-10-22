package main
import (
"fmt"
)
func main() {
	var x int
	var m1 int
	var m2 int
	var m3 int
	var k int
	fmt.Scan(&x)
	for i := 1; i <= x; i++ {
		fmt.Scan(&m1, &m2, &m3)
		if m1+m2+m3 >= 2 {
			k += 1
			}
		}
	fmt.Println(k)
}
