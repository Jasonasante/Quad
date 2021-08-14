package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	if x >= 1 {
		z01.PrintRune('o')
	}
	for a := 0; a < x-1; a++ {
		z01.PrintRune('-')
	}
	z01.PrintRune('o')
	z01.PrintRune('\n')
	for b := 0; b < y-1; b++ {
		if y >= 1 {
			z01.PrintRune('|')
		}
		for c := 0; c < x-1; c++ {
			z01.PrintRune(' ')
		}
		z01.PrintRune('|')
		z01.PrintRune('\n')
	}
	if x >= 1 {
		z01.PrintRune('o')
	}
	for a := 0; a < x-1; a++ {
		z01.PrintRune('-')
	}
	z01.PrintRune('o')
	z01.PrintRune('\n')
}
