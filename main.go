package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"unicode"

	"github.com/BurntSushi/toml"
)

func main() {

	clearScreen()

	config := getConfig(configPath)
	scanner := bufio.NewScanner(os.Stdin)

	// 0 start menu with logo
	// 1 start menu without logo
	// 2 choosing game mode
	// 3 playing adventure mode
	// 4 settings
	// 6 exiting
	// 7 tutorial
	state := 0
	var input int

	for {
		switch state {
		case 0:
			fmt.Println(logo, starDevil, welcomeMessage)
			state = 1
		case 1:
			fmt.Println(startingMenu)
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
				state = 7
			case 3:
				state = 1
			}
		case 3:
			if config.TotalScore <= 1 {
				showTutorial(&config, scanner)
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
		case 7:
			showTutorial(&config, scanner)
			state = 2
		}
	}
}

func showTutorial(config *Config, scanner *bufio.Scanner) {
	// Print tutorial text and wait until user triggers <Enter>
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
	fmt.Printf("%s\nYou see a %s.", *monster.ascii, *monster.name)
	playerHealth := setPlayerHealth(config.TotalScore)
	var input int
	var exp Expression
	for input != -1 {
		exp = nextExpression()
		fmt.Printf("\nThe monster's health: %d\nYour health: %d\nYou attack: %d x %d",
			monster.hp, playerHealth, exp.first, exp.second,
		)
		input := getInput(scanner)
		if input == exp.result {
			monster.hp -= exp.result
		} else {
			playerHealth -= exp.result
			fmt.Println(Red + "\nYou got ", exp.result, " of damage!" + Reset)
			if playerHealth < 1 {
				fmt.Printf("%s\n\n", death)
				return 1
			}
		}
		if monster.hp < 1 {
			fmt.Printf("\n\nThe monster is elliminated!")
			config.TotalScore += 1
			saveConfig(config, configPath)
			fmt.Printf(Green + "\n\nYour total score was increased and now it equals %d" + Reset, config.TotalScore)
			monster = getNextMonster(config.TotalScore)
			playerHealth = setPlayerHealth(config.TotalScore)
			fmt.Printf("%s\nYou see a %s.", *monster.ascii, *monster.name)
		}
	}
	return 0
}

func setPlayerHealth(totalScore int) int {
	// Decrease player's health
	// Kinda a rudimentary formula but whatever
	playerHealth := 50000 - totalScore * 100
	// Prevent it go below 1000
	playerHealth = max(playerHealth, 1000)
	return playerHealth
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
	fmt.Print(promptLine)
	scanner.Scan()

	input := strings.TrimSpace(scanner.Text())

	// Remove all non-digits in case user accidently
	// triggers "\", "]", etc., instead of <Enter>
	var b strings.Builder
	for _, r := range input {
		if unicode.IsDigit(r) {
			b.WriteRune(r)
		}
	}
	input = b.String()

	if number, err := strconv.Atoi(input); err == nil {
		return number
	}
	// Only digits are allowed
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

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func nextExpression() Expression {
	f := rand.Intn(89) + 10
	s := rand.Intn(89) + 10
	f, s = 100, 100
	exp := Expression{
		f, s, f * s,
	}
	return exp
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
