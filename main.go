package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
)

var commands = map[string]any{
	".help": func() {
		fmt.Println(helpMessage)
	},
	".exit": exitGame,
}

func main() {

	clearScreen()

	config := getConfig(configPath)
	scanner := bufio.NewScanner(os.Stdin)

	// Starting menu

	// 1 start menu
	// 2 choosing game mode
	// 3 playing adventure mode
	// 4 settings
	// 5 playing company mode
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
			if config.TotalScore <= 0 {
				fmt.Println(startingAdventureModeText)
			}
			adventureMode(config, scanner)
			switch input {

			}
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
}

func adventureMode(config Config, scanner *bufio.Scanner) {
	monster := getNextMonster(config.TotalScore)
	exp := nextExpression()
	fmt.Println(*monster.ascii, "You see a", *monster.name)
	fmt.Println("Health: ", monster.hp)
	var input int
	for input != -1 {
		fmt.Println("You attack: ", exp.first, "x", exp.second)
		input := getIntInput(scanner)
		if input == exp.result {
			monster.hp -= input
		} else {
			fmt.Println("You got ", input, " of damage!")
		}
		if monster.hp < 1 {
			fmt.Println("The monster is elliminated!")
		} else {
			fmt.Println("Health: ", monster.hp)
			exp = nextExpression()
		}
	}
}

// Returns to the starting menu
func exitGame() {

}

func getStrInput(scanner *bufio.Scanner) string {
	// Print prompt
	fmt.Print(cliName, "> ")
	for scanner.Scan() {
		input := cleanInput(scanner.Text())
		if input[0] == '.' {
			if command, exists := commands[input]; exists {
				command.(func())()
			}
		} else {
			return input
		}
		// Print prompt
		fmt.Print(cliName, "> ")
	}
	// This normally should not happen
	return ""
}

// Get next monster by considering totalScore
func getNextMonster(totalScore int) Monster {
	chosenGroup := monsters1
	if totalScore > 1000 {
		chosenGroup = monsters3
	} else if totalScore > 300 {
	}
	chosenMonster := chosenGroup[rand.Intn(len(chosenGroup))]
	resultedMonster := Monster{
		&chosenMonster.ascii,
		&chosenMonster.name,
		rand.Intn(chosenMonster.hpMax-chosenMonster.hpMin) + chosenMonster.hpMin,
		chosenMonster.damageMin,
		chosenMonster.hpMax,
	}
	return resultedMonster
}

func getIntInput(scanner *bufio.Scanner) int {
	// Print prompt
	fmt.Print(cliName, "> ")
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
		// Print prompt
		fmt.Print(cliName, "> ")
	}
	// This normally should not happen
	return 0
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

func printArbitraryAmountOfNewLines(amount int) {
	for range amount {
		fmt.Println()
	}
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

func saveConfig(cfg Config, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return toml.NewEncoder(f).Encode(cfg)
}

func readConfig(path string) (Config, error) {
	var cfg Config
	_, err := toml.DecodeFile(path, &cfg)
	return cfg, err
}
