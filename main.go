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

var cliName string = "starlight99"

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

var initialName = "AnonymousPlayer"

type Expression struct {
	first, second, result int
}

type Enemy struct {
	ascii string
	hp, damage int
}

var commands = map[string]any{
	".help": func() {
		fmt.Println(helpMessage)
	},
	".exit": exitGame,
}

func main() {

	config := getConfig(configPath)
	scanner := bufio.NewScanner(os.Stdin)

	// Starting menu

	// 1 start menu
	// 2 choosing game mode
	// 3 playing adventure mode
	// 5 playing company mode
	// 4 settings
	// 6 exiting
	state := 1
	var input int

	for {
		switch state {
		case 1:
			fmt.Println(logo, starDevil, startingMenu)
			input = getIntInput(scanner)
			switch input {
			case 1:
				state = 2
			case 2:
				state = 4
			case 3:
				state = 6
			}
		case 2:
			fmt.Println(gameModeMenu)
			input = getIntInput(scanner)
			switch input {
			case 1:
				state = 3
			case 2:
			case 3:
				state = 1
			}
		case 3:
			fmt.Println()
			input = getIntInput(scanner)
		case 4:
			fmt.Println(settingsMenu)
			input = getIntInput(scanner)
			switch input {
			case 3:
				fmt.Printf("\n%v\n", config)
			case 5:
				state = 1
			}
		case 6:
			fmt.Println(reallyWannaExit)
			input = getIntInput(scanner)
			switch input {
			case 1:
				fmt.Println(exitMessage)
				os.Exit(0)
			case 2:
				state = 1
			}
		}
	}

	// fmt.Println(exp.first, "x", exp.second, "= ???")
}

func getStrInput(scanner *bufio.Scanner) string {
	printPrompt()
	for scanner.Scan() {
		input := cleanInput(scanner.Text())
		if input[0] == '.' {
			if command, exists := commands[input]; exists {
				command.(func())()
			}
		} else {
			return input
		}
		printPrompt()
	}
	// This normally should not happen
	return ""
}

func nextMonster() 

func getIntInput(scanner *bufio.Scanner) int {
	printPrompt()
	for scanner.Scan() {
		input := cleanInput(scanner.Text())
		if number, err := strconv.Atoi(input); err == nil {
			return number
		} else if input[0] == '.' {
			if command, exists := commands[input]; exists {
				command.(func())()
			}
		} else {
			fmt.Println("???")
		}
		printPrompt()
	}
	// This normally should not happen
	return 0
}

func exitGame() {

}

func Settings() {

}

func initGame(config Config) {

}

func runGame(config Config) {
	// exp := nextExpression()
}

func getConfig(path string) Config {
	config, err := readConfig(path)
	if err != nil {
		// User runs game without config
		initialConfig := Config{
			PlayerName: initialName,
			TotalScore: 0,
		}
		// Try to save
		err = saveConfig(initialConfig, path)
		if err != nil {
			panic("Cannot save config file")
		}
		// Read again to make sure that everything is OK
		config, err = readConfig(path)
		if err != nil {
			panic("Cannot read config file")
		}
	}
	return config
}

func getRandomAsciiArt() {
	return
}

func flushStdin() {
	var discard string
	fmt.Scanln(&discard)
}

func startAdventureMode() {

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
		f, s, f * s,
	}
	return exp
}

func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}
