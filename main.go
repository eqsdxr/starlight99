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

func main() {

	clearScreen()

	config := getConfig(configPath)
	scanner := bufio.NewScanner(os.Stdin)

	// Starting menu

	// 1 start menu
	// 2 choosing game mode
	// 3 playing adventure mode
	// 4 settings
	// 6 exiting
	state := 1
	var input int

	for {
		switch state {
		case 1:
			fmt.Println(logo, starDevil, startingMenu)
			input = getInput(scanner)
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
			input = getInput(scanner)
			switch input {
			case 1:
				state = 3
			case 2:
			case 3:
				state = 1
			}
		case 3:
			if config.TotalScore <= 1 {
				showTutorial(config, scanner)
			}
			res := playGame(&config, scanner)
			switch res {
			case 1:
				state = 1
			}
		case 4:
			fmt.Println(settingsMenu)
			input = getInput(scanner)
			switch input {
			case 3:
				fmt.Printf("\n%v\n", config)
			case 5:
				state = 1
			}
		case 6:
			fmt.Println(reallyWannaExit)
			input = getInput(scanner)
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

func showTutorial(config Config, scanner *bufio.Scanner) {
	fmt.Println(tutorial1(config.PlayerName))
	getInput(scanner)
	fmt.Println(tutorial2)
	getInput(scanner)
	fmt.Println(tutorial3)
	getInput(scanner)
	fmt.Println(tutorial4)
	getInput(scanner)
	return
}

func playGame(config *Config, scanner *bufio.Scanner) int {
	monster := getNextMonster(config.TotalScore)
	exp := nextExpression()
	playerHealth := config.TotalScore * 10
	fmt.Println(*monster.ascii, "\nYou see a", *monster.name, "\nMonster's health: ",
		monster.hp, "\nYour health: ", playerHealth,
	)
	var input int
	for input != -1 {
		fmt.Println("You attack: ", exp.first, "x", exp.second)
		input := getInput(scanner)
		if input == exp.result {
			monster.hp -= exp.result
		} else {
			fmt.Println("You got ", input, " of damage!")
			playerHealth -= input
			if playerHealth < 1 {
				fmt.Println("You got killed!")
				return 1
			}
		}
		if monster.hp < 1 {
			fmt.Println("The monster is elliminated!")
			config.TotalScore += 1
			saveConfig(config, configPath)
			fmt.Println("Your total score was increased and now it's ", config.TotalScore)
			monster := getNextMonster(config.TotalScore)
			exp = nextExpression()
			playerHealth = config.TotalScore * 10
			fmt.Println(*monster.ascii, "\nYou see a", *monster.name, "\nMonster's health: ",
				monster.hp, "\nYour health: ", playerHealth,
			)
		} else {
			fmt.Println("Monster's health: ", monster.hp)
			fmt.Println("Your health: ", playerHealth)
			exp = nextExpression()
		}
	}
	return 0
}

// Returns to the starting menu
func exitGame() {

}

// func handleEvent() Event {}

// Get next monster by examining totalScore
func getNextMonster(totalScore int) Monster {
	chosenGroup := monsters1
	if totalScore > 500 {
		chosenGroup = monsters3
	} else if totalScore > 100 {
		chosenGroup = monsters2
	} else {
		chosenGroup = monsters1
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

// Only for int (it's done on purpose)
func getInput(scanner *bufio.Scanner) int {
	// Print prompt
	fmt.Print(promptLine)
	scanner.Scan()
	input := cleanInput(scanner.Text())
	if number, err := strconv.Atoi(input); err == nil {
		return number
	}
	return 0
}

func getConfig(path string) Config {
	config, err := readConfig(path)
	if err != nil {
		// User runs game without config
		initialConfig := Config{
			PlayerName: initialName,
			TotalScore: 1,
		}
		// Try to save
		err = saveConfig(&initialConfig, path)
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

func saveConfig(cfg *Config, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return toml.NewEncoder(f).Encode(*cfg)
}

func readConfig(path string) (Config, error) {
	var cfg Config
	_, err := toml.DecodeFile(path, &cfg)
	return cfg, err
}
