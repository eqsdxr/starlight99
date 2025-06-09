package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type GameState int

const (
	GameStartLogo GameState = iota
	GameStartNoLogo
	ChoosingGameMode
	PlayingAdventureMode
	Settings
	Exiting
	Tutorial
)

func main() {

	config := getConfig(configPath)
	scanner := bufio.NewScanner(os.Stdin)

	State := GameStartLogo
	var input int

	for {
		switch State {
		case GameStartLogo:
			fmt.Print(logo, starDevil, welcomeMessage)
			State = GameStartNoLogo
		case GameStartNoLogo:
			fmt.Print(startingMenu)
			input = getInput(scanner)
			switch input {
			case 1:
				State = ChoosingGameMode
			case 2:
				State = Settings
			case 3:
				State = Exiting
			}
		case ChoosingGameMode:
			fmt.Print(gameModeMenu)
			input = getInput(scanner)
			switch input {
			case 1:
				State = PlayingAdventureMode
			case 2:
				State = Tutorial
			case 3:
				State = GameStartNoLogo
			}
		case PlayingAdventureMode:
			if config.TotalScore <= 1 {
				showTutorial(config, scanner)
				fmt.Print("\nWould you like to choose a name for the character?",
					"(you can change it later in settings)\n1.Yes\n2.No")
				input := getInput(scanner)
				if input == 1 {
					changeCharacterName(config, scanner)
				}
			}
			res := playGame(config, scanner)
			switch res {
			case 0 | 1: // Death or a player has left
				State = GameStartLogo
			}
		case Settings:
			fmt.Print(settingsMenu)
			input = getInput(scanner)
			switch input {
			case 1:
				changeCharacterName(config, scanner)
			case 2:
				printStats(config, scanner)
			case 3:
				printAbout(scanner)
			case 4:
				State = GameStartNoLogo
			}
		case Exiting:
			fmt.Print(reallyWannaExit)
			input = getInput(scanner)
			switch input {
			case 1:
				fmt.Print(exitMessage)
				os.Exit(0)
			case 2:
				State = GameStartNoLogo
			}
		case Tutorial:
			showTutorial(config, scanner)
			State = ChoosingGameMode
		}
	}
}

func printAbout(scanner *bufio.Scanner) {
	fmt.Print(about)
	getInput(scanner)
}

func processPlayerDeath(config *Config, configPath string) {
	fmt.Print(death)
	config.Deaths++
	saveConfig(config, configPath)
}

func printStats(config *Config, scanner *bufio.Scanner) {
	fmt.Printf("\n\n\n\nPlayer's name: %d\nTotal score: %d\nDeaths: %d\n",
		config.PlayerName, config.TotalScore, config.Deaths)
	getInput(scanner)
	return
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
	monster := newMonster(config.TotalScore)
	playerHealth := setPlayerHealth(config.TotalScore)
	var input int
	var exp Expression
	var event Event

	event.newEvent(config.TotalScore)

	for {
		exp = generateNextExpression(config.TotalScore)
		printStatus(*monster, playerHealth, exp)
		input = getInput(scanner)

		if input == -2 {
			input = 999999
			exp.Result = 999999
			exp.Damage = 999999
		}

		if input == -1 {
			break
		} else if input == exp.Result {
			handleCorrectInput(monster, &playerHealth, &event, exp, config)
		} else {
			handleIncorrectInput(monster, &playerHealth, &event, exp, config)
		}

		if playerHealth < 1 {
			processPlayerDeath(config, configPath)
			return 1
		}

		if monster.HP < 1 {
			event.newEvent(config.TotalScore)
			fmt.Printf("\nThe monster is eliminated!")
			config.TotalScore++
			saveConfig(config, configPath)
			fmt.Printf(Green+"\nYour total score was increased and now it equals %d"+Reset, config.TotalScore)
			if config.TotalScore == 100 {
				fmt.Print(Magenta + "\nYou have become much stronger" + Reset)
			}
			monster = newMonster(config.TotalScore)
			playerHealth = setPlayerHealth(config.TotalScore)
			getInput(scanner)
		}
	}

	fmt.Print(Blue + "\nYou left..." + Reset)
	return 0
}

func newMonster(totalScore int) *Monster {
	monster := getNextMonster(totalScore)
	fmt.Printf("%s\nYou see a %s.", monster.ASCII, monster.Name)
	return monster
}

func printStatus(monster Monster, playerHealth int, exp Expression) {
	fmt.Printf("\nThe monster's health: %d\nYour health: %d\nYou attack: %d x %d",
		monster.HP, playerHealth, exp.First, exp.Second)
}

func handleCorrectInput(monster *Monster, playerHealth *int, event *Event, exp Expression, config *Config) {
	switch event.Type {
	case empty:
		monster.HP -= exp.Damage
	case miss:
		fmt.Print(event.Text)
	case additionalDamage:
		fmt.Print(event.Text)
		monster.HP -= exp.Damage * 3
	case totalScoreIncreased:
		config.TotalScore++
		saveConfig(config, configPath)
		fmt.Print(event.Text)
	case accidentalMonsterDeath:
		fmt.Print(event.Text)
		monster.HP = 0
	case accidentalPlayerDeath:
		fmt.Print(event.Text)
		*playerHealth -= 99999
	}

	if event.Type != empty {
		*event = emptyEvent
	}
}

func handleIncorrectInput(monster *Monster, playerHealth *int, event *Event, exp Expression, config *Config) {
	switch event.Type {
	case totalScoreIncreased:
		config.TotalScore++
		saveConfig(config, configPath)
		fmt.Print(event.Text)
		*playerHealth -= exp.Result
		fmt.Print(Red+"\nYou got ", exp.Damage, " of damage!"+Reset)
	case accidentalMonsterDeath:
		fmt.Print(event.Text)
		monster.HP = 0
	case accidentalPlayerDeath:
		fmt.Print(event.Text)
	default:
		*playerHealth -= exp.Damage
		fmt.Print(Red+"\nYou got ", exp.Damage, " of damage!"+Reset)
	}

	if event.Type != empty {
		*event = emptyEvent
	}
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
	// Decrease player's health and prevent it go below the threshold
	// This is kinda a rudimentary formula but whatever
	return int(max(10000-math.Log(float64(totalScore))*1500, 500))
}

// Get next monster by examining totalScore
func getNextMonster(totalScore int) *Monster {
	var chosenGroup []Monster
	if totalScore >= 500 {
		chosenGroup = monsters3
	} else if totalScore >= 100 {
		chosenGroup = monsters2
	} else {
		chosenGroup = monsters1
	}
	chosenMonster := chosenGroup[rand.Intn(len(chosenGroup))]
	chosenMonster.HP = rand.Intn(chosenMonster.HPMax-chosenMonster.HPMin) + chosenMonster.HPMin
	return &chosenMonster
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
func getConfig(path string) *Config {
	config, err := readConfig(path)
	if err != nil {
		// User runs game without config
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

func generateNextExpression(totalScore int) Expression {
	f := rand.Intn(89) + 10
	s := rand.Intn(89) + 10
	damage := f * s
	// It's too big to be counted as damage
	if totalScore < 100 {
		damage /= 10
	}
	return Expression{
		f, s, f * s, damage,
	}
}

func saveConfig(cfg *Config, path string) error {
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		return err
	}
	jsonStr, err := json.MarshalIndent(*cfg, "", "\t")
	if err != nil {
		return err
	}
	f.Write(jsonStr)
	return nil
}

func readConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = json.Unmarshal(data, &cfg)
	return &cfg, err
}
