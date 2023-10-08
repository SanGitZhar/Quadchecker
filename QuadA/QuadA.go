package main

import "github.com/01-edu/z01"

func main() {
	QuadA(5,3)
	z01.PrintRune('\n')
	QuadA(5,1)
	z01.PrintRune('\n')
	QuadA(1,1)
	z01.PrintRune('\n')
	QuadA(1,5)
}

func QuadA(x,y int) {
	if x > 0 && y > 0 {
		for i:=0; i<y; i++ {
			for j:=0; j<x; j++ {
				if i==0 || i==y-1 {
					if j==0 || j==x-1 {
						z01.PrintRune('o')
					} else {
						z01.PrintRune('-')
					}
				} else {
					if j==0 || j==x-1 {
						z01.PrintRune('|')
					} else {
						z01.PrintRune(' ')
					}
				}
			}
			z01.PrintRune('\n')
		}
	}
}
