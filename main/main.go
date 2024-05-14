package main

import (
	"log"
	"os"
	"strings"

	asciiart "asciiart/functions"
)

func main() {
	readFile := map[rune][]string{}

	if len(os.Args) != 2 && len(os.Args) != 3 {
		log.Fatalln("err: you shoud enter two or three argumrnts")
	}
	if len(os.Args) == 2 {
		readFile = asciiart.ReadFile("../sources/standard.txt")
	} else if !strings.HasSuffix(os.Args[2], ".txt") {
		readFile = asciiart.ReadFile("../sources/" + os.Args[2] + ".txt")
	} else if strings.HasSuffix(os.Args[2], ".txt") {
		readFile = asciiart.ReadFile("../sources/" + os.Args[2])
	} else {
		log.Fatalln("err: the third argument should be one of these file names (standerd),(shadow),(thinkertoy)")
	}
	// check if the input is among ascii manual
	checkTheChar := asciiart.CheckTheChar(os.Args[1])
	// go print my argument
	asciiart.FindAndPrint(checkTheChar, readFile)
}
