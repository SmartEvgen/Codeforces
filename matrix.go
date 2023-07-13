package main

import (
    "fmt"
)

func main() {
	m1 := make([]int, 5)  
	m2 := make([]int, 5) 
	m3 := make([]int, 5)  
	m4 := make([]int, 5) 
	m5 := make([]int, 5) 
	fmt.Scan(&m1)
	fmt.Scan(&m2)
	fmt.Scan(&m3)
	fmt.Scan(&m4)
	fmt.Scan(&m5)
	var t int
	var v int
	var a int
	var itog1 int
	var itog2 int
	b := 0
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		if a == 1 {
			t = 1
			v = i+1
			}
		m1[b] = a
		b++
		}
	b = 0
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		if a == 1 {
			t = 2
			v = i+1
			}
		m2[b] = a
		b++
		}
	b = 0
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		if a == 1 {
			t = 3
			v = i+1
			}
		m3[b] = a
		b++
		}
	b = 0
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		if a == 1 {
			t = 4
			v = i+1
			}
		m4[b] = a
		b++
		}
	b = 0
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		if a == 1 {
			t = 5
			v = i+1
			}
		m5[b] = a
		b++
		}
	

	if v == 3 && t == 3 {
		fmt.Println(0)
		} else if v != 3 || t != 3 {
			itog1 = v - 3
			itog2 = t - 3
			if itog1 < 0 {
				itog1 = itog1*(-1)
				}
			if itog2 < 0 {
				itog2 = itog2*(-1)
				}
			fmt.Println(itog1 + itog2)
			}
}
