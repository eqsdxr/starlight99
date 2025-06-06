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
			fmt.Print(logo, starDevil, welcomeMessage)
			state = 1
		case 1:
			fmt.Print(startingMenu)
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
			fmt.Print(gameModeMenu)
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
				fmt.Print("\nWould you like to choose a name for the character?",
					"(you can change it later in settings)\n1.Yes\n2.No")
				input := getInput(scanner)
				if input == 1 {
					changeCharacterName(&config, scanner)
				}
			}
			res := playGame(&config, scanner)
			switch res {
			case 1: // Death
				state = 1
			case 0: // A player has left
				state = 1
			}
		case 4:
			fmt.Print(settingsMenu)
			input = getInput(scanner)
			switch input {
			case 1:
				changeCharacterName(&config, scanner)
			case 3:
				fmt.Printf("\n%v", config)
			case 5:
				state = 1
			}
		case 6:
			fmt.Print(reallyWannaExit)
			input = getInput(scanner)
			switch input {
			case 1:
				fmt.Print(exitMessage)
				os.Exit(0)
			case 2:
				state = 1
			}
		case 7:
			showTutorial(&config, scanner)
			state = 2
		default:
			panic("\nInvalid state 1")
		}

	}
}

func changeCharacterName(config *Config, scanner *bufio.Scanner) {
	fmt.Print("Write the name (numbers only):")
	input := getInput(scanner)
	config.PlayerName = input
	saveConfig(config, configPath)
	fmt.Print("\nNow he's called ", config.PlayerName)
	return
}

// Print tutorial text and wait until user triggers <Enter>
func showTutorial(config *Config, scanner *bufio.Scanner) {
	fmt.Print(tutorial1(config.PlayerName))
	getInput(scanner)
	fmt.Print(tutorial2)
	getInput(scanner)
	fmt.Print(tutorial3)
	getInput(scanner)
	fmt.Print(tutorial4)
	getInput(scanner)
	return
}

func playGame(config *Config, scanner *bufio.Scanner) int {
	monster := getNextMonster(config.TotalScore)
	fmt.Printf("%s\nYou see a %s.", *monster.ASCII, *monster.Name)
	playerHealth := setPlayerHealth(config.TotalScore)
	var input int
	var exp Expression
	var event Event
	for {
		exp = generateNextExpression()
		fmt.Printf("\nThe monster's health: %d\nYour health: %d\nYou attack: %d x %d",
			monster.HP, playerHealth, exp.First, exp.Second,
		)
		input = getInput(scanner)
		event.newEvent(config.TotalScore)
		if input == -1 {
			break
		} else if input == exp.Result {
			exp.Result = calibrateDamage(exp.Result, config.TotalScore)
			switch event.Type {
			case empty:
				monster.HP -= exp.Result
			case miss:
				// Check if event has already happened in current fight
				if event.OptValue == 0 {
					fmt.Print(event.Text) // How to say that the input was correct?
					// Set flag that event has happened
					event.OptValue = 1
				} else {
					monster.HP -= exp.Result
				}
			case additionalDamage:
				fmt.Print(event.Text)
				monster.HP -= exp.Result * event.OptValue
			case totalScoreIncreased:
				config.TotalScore += 1
				saveConfig(config, configPath)
				fmt.Print(event.Text)
			case accidentalMonsterDeath:
				fmt.Print(event.Text)
				monster.HP -= 99999
			case accidentalPlayerDeath:
				fmt.Print(event.Text)
				playerHealth -= 99999
				if playerHealth < 1 {
					fmt.Print(death)
					return 1
				}
			}
		} else {
			playerHealth -= exp.Result
			fmt.Print(Red+"\nYou got ", exp.Result, " of damage!"+Reset)
			if playerHealth < 1 {
				fmt.Print(death)
				return 1
			}
		}
		if monster.HP < 1 {
			fmt.Printf("\nThe monster is elliminated!")
			config.TotalScore++
			saveConfig(config, configPath)
			fmt.Printf(Green+"\nYour total score was increased and now it equals %d"+Reset, config.TotalScore)
			if config.TotalScore == 100 {
				fmt.Print(Magenta + "\nYou have become much stronger" + Reset)
			}
			monster = getNextMonster(config.TotalScore)
			playerHealth = setPlayerHealth(config.TotalScore)
			fmt.Printf("%s\nYou see a %s.", *monster.ASCII, *monster.Name)
		}
	}
	fmt.Print(Blue + "\nYou left..." + Reset)
	return 0
}

func calibrateDamage(damage, totalScore int) {

}

// Generate a battle game event (or no event) with some probability
func (event *Event) newEvent(totalScore int) {
	// The eventList is sorted in *decreasing order* so that events
	// with lower probability but which are devisible by events
	// with higher probability are able to occur and not be
	// overshadowed with the ones with higer probability
	for _, e := range eventList {
		// Rudimentary probability calculating
		if totalScore%e.Probability == 0 {
			// Prevent overwriting
			if event.Type != e.Type {
				*event = e
				return // Return to avoid overwriting by emptyEvent
			}
		}
	}
	// Assign empty event type if it's not already
	// to reset an occured event to prevent it from
	// repeating
	if event.Type != empty {
		*event = emptyEvent
	}
}

// Calculate and return starting player's health
// depending on totalScore
func setPlayerHealth(totalScore int) int {
	// Decrease player's health
	// This is kinda a rudimentary formula but whatever
	playerHealth := 50000 - totalScore*100
	// Prevent it go below 1000
	playerHealth = max(playerHealth, 1000)
	return playerHealth
}

// Get next monster by examining totalScore
func getNextMonster(totalScore int) Monster {
	chosenGroup := monsters1
	if totalScore >= 500 {
		chosenGroup = monsters3
	} else if totalScore >= 100 {
		chosenGroup = monsters2
	} else {
		chosenGroup = monsters1
	}
	chosenMonster := chosenGroup[rand.Intn(len(chosenGroup))]
	resultedMonster := Monster{
		&chosenMonster.ASCII,
		&chosenMonster.Name,
		rand.Intn(chosenMonster.HPMax-chosenMonster.HPMin) + chosenMonster.HPMin,
		chosenMonster.DamageMin,
		chosenMonster.HPMax,
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
		if unicode.IsDigit(r) || r == '-' {
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

// Get config from saved file or create new. It panics
// if config cannot be saved and read for some reason
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
			panic("\nCannot save config file")
		}
		// Read again to make sure that everything is OK
		config, err = readConfig(path)
		if err != nil {
			panic("\nCannot read config file")
		}
	}
	return config
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func generateNextExpression() Expression {
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
