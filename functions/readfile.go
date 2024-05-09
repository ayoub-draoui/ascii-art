package asciiart

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFile(fileName string) map[rune][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("someting went wrong")
	}
	defer file.Close()
	data := make(map[rune][]string)


	scaner := bufio.NewScanner(file)
	// "./sources/standard.txt"
	for scaner.Scan() {
		line := scaner.Text()
		fmt.Println(line)
	}
}
