package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 0; i <= 9; i++ {
		for j := i + 1; j <= 9; j++ {
			for k := j + 1; k <= 9; k++ {
				if i != j && i != k && i < j && j < k {
					if i != 0 || j != 1 || k != 2 {
						z01.PrintRune(' ')
					}
					z01.PrintRune(rune(48 + i))
					z01.PrintRune(rune(48 + j))
					z01.PrintRune(rune(48 + k))
					if i != 7 || j != 8 || k != 9 {
						z01.PrintRune(',')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
