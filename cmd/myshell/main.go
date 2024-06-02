package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
			switch commands[1] {
			case "exit", "echo", "type":
				fmt.Printf("%s is a shell builtin\n", commands[1])
			default:
				DoType(commands[1])
			}
		} else if commands[0] == "echo" {
			fmt.Fprintf(os.Stdout, "%s\n", strings.Join(commands[1:], " "))
		} else {
			cmd := exec.Command(commands[0], commands[1:]...)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			err := cmd.Run()
			if err != nil {
				fmt.Printf("%s: command not found\n", commands[0])
			}
		}
	}

}

func DoType(param string) {
	env := os.Getenv("PATH")
	paths := strings.Split(env, ":")
	for _, path := range paths {
		fileDir := path + "/" + param
		if _, err := os.Stat(fileDir); err == nil {
			fmt.Fprintf(os.Stdout, "%s is %s\n", param, fileDir)
			return
		}
	}
	fmt.Fprintf(os.Stdout, "%s: not found\n", param)
}
