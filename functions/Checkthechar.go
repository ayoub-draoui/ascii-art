package asciiart

import (
	"fmt"
	"strings"
)


func CheckTheChar(text string) []string {
	for _, char := range text{
		if rune (char) < 32 || rune (char) > 126{
			fmt.Println("please enter a valid character")
		}

	}
	line := strings.Split(text, "\\n")
	return line
} 