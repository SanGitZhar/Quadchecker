package main

import "github.com/01-edu/z01"

func main() {
	QuadD(5,3)
	z01.PrintRune('\n')
	QuadD(5,1)
	z01.PrintRune('\n')
	QuadD(1,1)
	z01.PrintRune('\n')
	QuadD(1,5)
}

func QuadD(x,y int) {
	if x > 0 && y > 0 {
		for i:=0; i<y; i++ {
			for j:=0; j<x; j++ {
				if i==0 || i==y-1 {
					if j==0 {
						z01.PrintRune('A')
					} else if j==x-1 {
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
}
