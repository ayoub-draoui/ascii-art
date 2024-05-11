package asciiart

import (
	"fmt"
)

func FindAndPrint(checkTheChar []string, readFile map[rune][]string) {
	is_printed := false
	for idx, word := range checkTheChar {
		if word != "" {
			i := 0
			for i < 8 {
				line := ""
				for _, char := range word {
					line = readFile[char][i]
					fmt.Print(line)
				}
				fmt.Println()
				i++
				is_printed = true
			}

		} else {
			if idx == len(checkTheChar)-1 && !is_printed {
				continue

			}else{

				fmt.Println()
			}
		}
	}
}
