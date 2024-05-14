package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestMainFunction(t *testing.T) {
	x := map[string]string{
		"hello": `
 _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
`,
	}
	for input, expectOutput := range x {
		cmd := exec.Command("go", "run", "main.go", input)
		outputbyte, _ := cmd.CombinedOutput()
		output := string(outputbyte)
		if strings.TrimSpace(output) == strings.TrimSpace(expectOutput) {
			fmt.Printf(input)
		} else {
			fmt.Println("failed")
		}
	}
}
