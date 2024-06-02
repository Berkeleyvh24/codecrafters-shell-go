package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		res := strings.TrimSpace(input)
		commands := strings.Split(res, " ")
		if commands[0] == "exit" {
			os.Exit(0)

		} else if commands[0] == "type" {
			// cmd := exec.Command("type", strings.Join(commands[1:], " "))
			// output, err := cmd.CombinedOutput()
			// if err != nil {
			// 	newString := string(output[:])
			// 	stringArray := strings.Split(newString, " ")
			// 	finalString := strings.Join(stringArray[4:], " ")
			// 	fmt.Fprintf(os.Stdout, "%s", strings.Replace(finalString, ":", "", -1))
			// } else {
			// 	fmt.Printf("%s", output)
			// }
			switch commands[1] {
			case "exit", "echo", "type":
				fmt.Printf("%s is a shell builtin\n", commands[1])
			default:
				fmt.Printf("%s not found\n", commands[1])
			}
		} else if commands[0] == "echo" {
			fmt.Fprintf(os.Stdout, "%s\n", strings.Join(commands[1:], " "))
		} else {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", res)
		}
	}

}
