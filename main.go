package main

import (
	"log"
	"os"

	asciiart "asciiart/functions"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("err: you shoud enter two argumrnts")
	}
	// check if the input is among ascii manual
	checkTheChar := asciiart.CheckTheChar(os.Args[1])
	// go read the file standard
	readFile := asciiart.ReadFile("./sources/standard.txt")
	// go print my argument  
	asciiart.FindAndPrint(checkTheChar, readFile)
}
