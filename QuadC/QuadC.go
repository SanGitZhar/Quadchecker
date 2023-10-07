package main

import "github.com/01-edu/z01"

func main() {
	QuadC(5,3)
	z01.PrintRune('\n')
	QuadC(5,1)
	z01.PrintRune('\n')
	QuadC(1,1)
	z01.PrintRune('\n')
	QuadC(1,5)
}

func QuadC(x,y int) {
	for i:=0; i<y; i++ {
		for j:=0; j<x; j++ {
			if i==0 {
				if j==0 || j==x-1 {
					z01.PrintRune('A')
				} else {
					z01.PrintRune('B')
				}
			} else if i==y-1 {
				if j==0 || j==x-1 {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}
			} else {
				if j==0 || j==x-1 {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}
