package piscine

import "github.com/01-edu/z01"

func QuadB(x, y int) {
	if x >= 1 {
		z01.PrintRune('/')
	}
	for a := 1; a < x-1; a++ {
		z01.PrintRune('*')
	}
	if x > 1 {
		z01.PrintRune(92)
	}
	z01.PrintRune('\n')
	for b := 1; b < y-1; b++ {
		if y > 1 {
			z01.PrintRune('*')
		}
		for c := 1; c < x-1; c++ {
			z01.PrintRune(' ')
		}
		if x > 1 {
			z01.PrintRune('*')
		}
		z01.PrintRune('\n')
	}
	if x > 1 && y > 1 {
		z01.PrintRune(92)
	}
	for a := 1; a < x-1 && y > 1; a++ {
		z01.PrintRune('*')
	}
	if y > 1 {
		z01.PrintRune('/')
		z01.PrintRune('\n')
	}
}
