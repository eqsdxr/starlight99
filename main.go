package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var cliName string = "gotrain"

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

type Expression struct {
	f, s, r int
}

func main() {
	showMenu(0)
	fmt.Println()
	showMenu(1)
	return
	commands := map[string]any{
		".help": printHelp,
	}
	exp := nextExpression()
	reader := bufio.NewScanner(os.Stdin)
	printHelp()
	fmt.Println(exp.f, "x", exp.s, "= ???")
	printPrompt()
	for reader.Scan() {
		input := cleanInput(reader.Text())
		if input == "" {
			continue
		} else if strings.EqualFold(".exit", input) {
			return
		} else if number, err := strconv.Atoi(input); err == nil {
			if number != exp.r {
				fmt.Println(exp.f, "x", exp.s, "= ???")
			}
		} else if command, exists := commands[input]; exists {
			command.(func())()
		} else {
			handleCmd(input)
		}
		clearScreen()
		printArbitraryAmountOfNewLines(30)
		exp = nextExpression()
		fmt.Println(exp.f, "x", exp.s, "= ???")
		printPrompt()
	}
}

func showMenu(state int) {
	switch state {
	case 0:
		fmt.Println("Welcome to ", cliName, ", my dear wanderer! Choose an option:")
		fmt.Println("1. Play ", cliName)
		fmt.Println("2. Settings")
		fmt.Println("3. Exit ", cliName)
	case 1:
		fmt.Println("Choose the difficulty level:")
		fmt.Println()
		fmt.Println("1.  " + Green + "Easy" + Reset)
		fmt.Println("2.  " + Yellow + "Medium " + Reset)
		fmt.Println("3.  " + Red + "Hard" + Reset)
		fmt.Println("4.  " + Magenta + "Very Hard" + Reset)
		fmt.Println()
	default:
		fmt.Println(Red + "Critical error! Incorrect state number was passed to showMenu function!" + Red)
	}
}

func flushStdin() {
	var discard string
	fmt.Scanln(&discard)
}

func startAdventureMode(difficulty int) {

}

func printArbitraryAmountOfNewLines(amount int) {
	for range amount {
		fmt.Println()
	}
}

func printPrompt() {
	fmt.Print(cliName, "> ")
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}


func nextExpression() Expression {
	f := rand.Intn(89) + 10
	s := rand.Intn(89) + 10
	f = 10
	s = 10
	exp := Expression{
		f, s, f*s,
	}
	return exp
}

func printUnknown(text string) {
	fmt.Println(text, ": command not found")
}

func printHelp() {
	fmt.Println("Available commands:")
	fmt.Println("\\help - show this text")
	fmt.Println("\\exit - exit the program")
}

func handleInvalidCmd(text string) {
	printUnknown(text)
}

func handleCmd(text string) {
	handleInvalidCmd(text)
}

func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}
