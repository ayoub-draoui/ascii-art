package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	asciiart "asciiart/functions"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("err: you shoud enter two argumrnts")
	}
	// check if the input is among ascii manual
	changed := asciiart.CheckTheChar(os.Args[1])
	// go read the file standard
	redFile := asciiart.ReadFile("./sources/standard.txt")

}
