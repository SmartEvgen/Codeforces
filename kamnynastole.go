package main

import "fmt"

func main() {
	var zero int = 0
	var num int
	var lett string
	var sh int
	fmt.Scan(&num)
	fmt.Scan(&lett)
	for i := 1; i <= num-1; i++ {
		if lett[zero] == lett[zero+1] {
			zero++
			sh++
			} else {
				zero++
				}
		}
		fmt.Println(sh)
	}


