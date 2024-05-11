package asciiart

import (
	"fmt"
)

func FindAndPrint(checkTheChar []string, readFile map[rune][]string) {
	for _, word := range checkTheChar {
		i := 0
		for i < 8 {
			line := ""
			for _, char := range word {
				line = readFile[char][i]
				fmt.Print(line)
			}
			fmt.Println()
			i++
		}
	}
}
