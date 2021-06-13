package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	start()
}

// start starts our shell.
func start() {
	scanner := bufio.NewScanner(os.Stdin)
	showPrompt()

	for scanner.Scan() {
		// text: cat args1 args2
		text := scanner.Text()
		textSlice := strings.Split(text, " ")

		fmt.Println(textSlice)

		cmd := exec.Command(textSlice[0], textSlice[1:]...)

		// Execute the command and return its output.
		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
			log.Fatal(err)
		}
		fmt.Print(string(stdout))
		showPrompt()
	}
}
func showPrompt() {
	fmt.Print("> ")
}
