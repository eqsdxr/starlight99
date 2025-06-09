package main

import "fmt"

type Config struct {
	PlayerName int `json:"playername"`
	TotalScore int `json:"totalscore"`
	Deaths     int `json:"deaths"`
}

type Monster struct {
	ASCII, Name                            string
	HP, HPMin, HPMax, DamageMin, DamageMax int
}

type Expression struct {
	First, Second, Result, Damage int
}

type Event struct {
	Text                        string
	Type, OptValue, Probability int
}

var (
	initialConfig = Config{
		PlayerName: initialName,
		TotalScore: 1,
	}

	configPath string = "/home/ivan/.local/share/starlight99.json"

	cliName     string = "starlight99"
	initialName        = 99

	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"

	empty                  = 0
	miss                   = 1
	additionalDamage       = 2
	unexpectedHealing      = 3
	totalScoreIncreased    = 4
	accidentalMonsterDeath = 5
	accidentalPlayerDeath  = 6

	emptyEvent                  = Event{"", 0, empty, 0}
	missEvent                   = Event{Gray + "You missed!" + Reset, miss, 0, 10}
	criticalDamageEvent         = Event{Yellow + "Additional damage to the opponent!" + Reset, additionalDamage, 3, 20}
	totalScoreIncreasedEvent    = Event{Green + "Your total score accidently increased!" + Reset, totalScoreIncreased, 0, 99}
	accidentalMonsterDeathEvent = Event{"The monster suddenly died!", accidentalMonsterDeath, 0, 100}
	accidentalPlayerDeathEvent  = Event{Red + "The character suddenly died!" + Reset, accidentalPlayerDeath, 0, 10000}

	// Don't add emptyEvent to eventList since there's no point in it
	// It is sorted in a *decreasing* order
	eventList = []Event{accidentalPlayerDeathEvent, accidentalMonsterDeathEvent, totalScoreIncreasedEvent,
		criticalDamageEvent, missEvent}

	promptLine string = fmt.Sprintf("\n\n%s> ", cliName)

	welcomeMessage string = `

 Welcome to Starlight99 dear wanderer!`
	startingMenu string = `

 Choose an option:
 1. Play
 2. Settings
 3. Exit`
	reallyWannaExit string = `

 Do you really want to exit?
 1. Yes
 2. No`
	exitMessage string = `

 Exiting...`
	gameModeMenu string = `

 Choose the game mode:
 1. Adventure Mode
 2. Show tutorial
 3. Go back`

	tutorial1 = func(playerName int) string {
		return fmt.Sprintf(
			"\n\n"+` Hello %d, Welcome to %s!`+
				"\n\n"+` This is a short tutorial before you start.`+
				"\n\n"+` To continue, press <Enter>`, playerName, cliName,
		)
	}

	tutorial2 string = `

 On your journey, you have fought countless monsters. You're
 really strong so they shouldn't be a big problem in a normal
 situation (except for a few weird ones). But it's not quite the case for you.
 You're actually a bit mental and every time you need to make a hit, you force
 yourself to count some random numbers in your head, you already tried
 hundreds of methods to stop that and the only result you got was that
 numbers are not float anymore (!).`

	tutorial3 string = `

 Once you discovered and started coping with this special side of yours, you have found
 methods how to do it faster and easier. For example, you fight a regular bat
 and when you're going to make a punch, a combination 59*71 suddenly appears in you head.
 To count the result, you can use the following algorithm:
 
 5 * 7 = 35
 5*1 + 9*7 = 68
 9 * 1 = 9
 
 +___9
 +_68
 +35
 =4189
 
 Which makes 4189 damage to a poor bat! But be careful. If you're unlucky enough to make a careless mistake
 it will be you who gets these 4189 damage (don't ask why, I don't know). By the way, if you want to escape a battle,
 use -1.`

	tutorial4 string = `

 So this is the end. Good luck, and have fun!`

	changingCharacterName string = `

 Write new name for your character (numbers only):`

	settingsMenu string = `

 1. Change the character's name
 2. View stats
 3. About
 4. Go back`

	about string = `
 For all additional information refer to the source code:
 
 https://github.com/eqsdxr/starlight99
`

	monsters1 = []Monster{bearMonster, batMonster, scorpionMonster, spiderMonster, ravenMonster}
	monsters2 = []Monster{centaurMonster, gryphonMonster, grimReaperMonster, unicornMonster, phoenixMonster, devilMonster}
	monsters3 = []Monster{starDevilMonster, skeletonMonster, dragonMonster, foxMonster}

	allMonsters = func() []Monster {
		all := append(monsters1, monsters2...)
		return append(all, monsters3...)
	}

	starDevilMonster = Monster{starDevil, "Star Devil", 0, 10000, 12000, 500, 1000}
	skeletonMonster  = Monster{skeleton, "Skeleton", 0, 8000, 10000, 400, 800}
	dragonMonster    = Monster{dragon, "Dragon", 0, 7000, 9000, 350, 700}
	foxMonster       = Monster{fox, "Fox", 0, 6000, 8000, 300, 600}

	devilMonster      = Monster{devil, "Devil", 0, 3000, 5000, 200, 400}
	phoenixMonster    = Monster{phoenix, "Phoenix", 0, 3000, 5000, 200, 400}
	unicornMonster    = Monster{unicorn, "Unicorn", 0, 3000, 5000, 200, 400}
	grimReaperMonster = Monster{grimReaper, "Grim Reaper", 0, 3200, 5200, 220, 450}
	gryphonMonster    = Monster{gryphon, "Gryphon", 0, 3400, 5400, 250, 470}
	centaurMonster    = Monster{centaur, "Centaur", 0, 3500, 5500, 250, 500}

	bearMonster     = Monster{bear, "Bear", 0, 600, 900, 40, 80}
	spiderMonster   = Monster{spider, "Spider", 0, 300, 500, 20, 40}
	scorpionMonster = Monster{scorpion, "Scorpion", 0, 200, 400, 15, 30}
	ravenMonster    = Monster{raven, "Raven", 0, 150, 300, 10, 20}
	batMonster      = Monster{bat, "Bat", 0, 100, 200, 5, 10}

	// https://patorjk.com/software/taag/#p=display&f=ANSI%20Regular&t=
	logo string = Red + `

   ██████ ▄▄▄█████▓ ▄▄▄       ██▀███   ██▓     ██▓  ▄████  ██░ ██ ▄▄▄█████▓  ████████     ████████ 
 ▒██    ▒ ▓  ██▒ ▓▒▒████▄    ▓██ ▒ ██▒▓██▒    ▓██▒ ██▒ ▀█▒▓██░ ██▒▓  ██▒ ▓▒ ██    █▒░░   ██    █▒░░
 ░ ▓██▄   ▒ ▓██░ ▒░▒██  ▀█▄  ▓██ ░▄█ ▒▒██░    ▒██▒▒██░▄▄▄░▒██▀▀██░▒ ▓██░ ▒░  ▓██▄ ██░     ▓██▄ ██░
   ▒   ██▒░ ▓██▓ ░ ░██▄▄▄▄██ ▒██▀▀█▄  ▒██░    ░██░░▓█  ██▓░▓█ ░██ ░ ▓██▓ ░  ░       ██░  ░      ████░
 ▒██████▒▒  ▒██▒ ░  ▓█   ▓██▒░██▓ ▒██▒░██████▒░██░░▒▓███▀▒░▓█▒░██▓  ▒██▒ ░  ░ ▒▒██████▒▒ ░ ▒▒████▒▒
 ▒ ▒▓▒ ▒ ░  ▒ ░░    ▒▒   ▓▒█░░ ▒▓ ░▒▓░░ ▒░▓  ░░▓   ░▒   ▒  ▒ ░░▒░▒  ▒ ░░     ▒  ▒ ▒▓▒ ▒   ▒  ▒ ▒▓▒ ▒
 ░ ░▒  ░ ░    ░      ▒   ▒▒ ░  ░▒ ░ ▒░░ ░ ▒  ░ ▒ ░  ░   ░  ▒ ░▒░ ░    ░       ▒ ░░░▒  ░   ▒ ░░░▒  ░
 ░  ░  ░    ░        ░   ▒     ░░   ░   ░ ░    ▒ ░░ ░   ░  ░  ░░ ░  ░        ▒  ░  ░     ▒  ░  ░
       ░                 ░  ░   ░         ░  ░ ░        ░  ░  ░  ░
` + Reset

	// https://www.asciiart.eu/mythology/skeletons
	death string = `
 
                            ,--.
                           {    }
                           K,   }
                          /  'Y'
                     _   /   /
                    {_'-K.__/
                      '/-.__L._
                      /  ' /'\_}
                     /  ' /     -ART BY ZEUS-
             ____   /  ' /
      ,-'~~~~    ~~/  ' /_
    ,'             ''~~~%%',
   (                     %  Y
  {                      %% I
 {      -                 %  '.
 |       ',                %  )
 |        |   ,..__      __. Y
 |    .,_./  Y ' / ^Y   J   )|
 \           |' /   |   |   ||    You got killed!
  \          L_/    . _ (_,.'(
   \,   ,      ^^""' / |      )
     \_  \          /,L]     /
       '-_'-,       ' '   ./'
          '-(_            )
              ^^\..___,.--'`
 
	// https://www.asciiart.eu/mythology/devils
	starDevil string = `

             ._                                            ,
              (')..                                    ,.-')
               (',.)-..                            ,.-(..')
                (,.' ,.)-..                    ,.-(. '.. )
                 (,.' ..' .)-..            ,.-( '.. '.. )
                  (,.' ,.'  ..')-.     ,.-( '. '.. '.. )
                   (,.'  ,.' ,.'  )-.-('   '. '.. '.. )
                    ( ,.' ,.'    _== ==_     '.. '.. )
                     ( ,.'   _==' ~  ~  '==_    '.. )
                      \  _=='   ----..----  '==_   )
                   ,.-:    ,----___.  .___----.    -..
               ,.-'   (   _--====_  \/  _====--_   )  '-..
           ,.-'   .__.''.  '-_I0_-'    '-_0I_-'  .''.__.  '-..
       ,.-'.'   .'      (          |  |          )      '.   '.-..
   ,.-'    :    '___--- ''.__.    / __ \    .__.' '---___'    :   '-..
 -'_________'-____________'__ \  (O)  (O)  / __'____________-'________'-
                             \ . _  __  _ . /
                              \ 'V-'  '-V' |
                               | \ \ | /  /
                                V \ ~| ~/V
                                 |  \  /|
                                  \~ | V             - JGG
                                   \  |
                                    VV`
 
	// https://www.asciiart.eu/mythology/skeletons
	skeleton string = `

                               _.--""-._
   .                         ."         ".
  / \    ,^.         /(     Y             |      )\
 /   '---. |--'\    (  \__..'--   -   -- -'""-.-'  )
 |        :|    '>   '.     l_..-------.._l      .'
 |      __l;__ .'      "-.__.||_.-'v'-._||'"----"
  \  .-' | |  '              l._       _.'
   \/    | |                   l'^^'^^'j
         | |                _   \_____/     _
         j |               l '--__)-'(__.--' |
         | |               | /'---''-----'"1 |  ,-----.
         | |               )/  '--' '---'   \'-'  ___  '-.
         | |              //  '-'  ''----'  /  ,-'   I'.  \
       _ L |_            //  '-.-.''-----' /  /  |   |  '. \
      '._' / \         _/(   '/   )- ---' ;  /__.J   L.__.\ :
       '._;/7(-.......'  /        ) (     |  |            | |
       '._;l _'--------_/        )-'/     :  |___.    _._./ ;
         | |                 .__ )-'\  __  \  \  I   1   / /
         '-'                /   '-\-(-'   \ \  '.|   | ,' /
                            \__  '-'    __/  '-. '---'',-'
                               )-._.-- (        '-----'
                              )(  l\ o ('..-.
                        _..--' _'-' '--'.-. |
                 __,,-'' _,,-''            \ \
                f'. _,,-'                   \ \
               ()--  |                       \ \
                 \.  |                       /  \
                   \ \                      |._  |
                    \ \                     |  ()|
                     \ \                     \  /
                      ) '-.                   | |
                     // .__)                  | |
                  _.//7'                      | |
                '---'                         j_| '
                                             (| |
                                              |  \
                                              |lllj
                                              |||||  -nabis`

	// https://www.asciiart.eu/mythology/dragons
	dragon string = `

                  ___====-_  _-====___
            _--^^^#####//      \\#####^^^--_
         _-^##########// (    ) \\##########^-_
        -############//  |\^^/|  \\############-
      _/############//   (@::@)   \\############\_
     /#############((     \\//     ))#############\
    -###############\\    (oo)    //###############-
   -#################\\  / VV \  //#################-
  -###################\\/      \//###################-
 _#/|##########/\######(   /\   )######/\##########|\#_
 |/ |#/\#/\#/\/  \#/\##\  |  |  /##/\#/  \/\#/\#/\#| \|
 '  |/  V  V  '   V  \#\| |  | |/#/  V   '  V  V  \|  '
    '   '  '      '   / | |  | | \   '      '  '   '
                     (  | |  | |  )
                    __\ | |  | | /__
                   (vvv(VVV)(VVV)vvv)           -???`

	// https://emojicombos.com/kawaii-anime-girl-ascii-art
	fox string = `

                         ⠉⠙⠻⢿⣛⣋⠉⠉⠉⠉⠉⠒⠲⠤⣤⢾⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
 ⠀⠀⠀⠀⠀⠀⠀⢀⡤⠴⠒⠒⠒⠒⠒⠦⢤⣀⠀⠀⠀⠀⠀⠀⢀⣠⠤⠒⠉⠁⠀⠀⠀⠀⠀⠀⠀⡴⡱⣿⠦⡀⠀⢀⠤⠆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
 ⠀⠀⠀⠀⠀⢀⡜⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠙⠦⡉⠒⣺⠭⠅⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡜⡑⠀⡿⠀⠈⠞⠁⡘⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
 ⠀⠀⠀⠀⠀⣸⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡼⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⠎⡰⠁⠸⡇⠀⠀⡀⡰⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
 ⠀⠀⠀⠀⠀⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣰⠋⠀⠀⣀⣠⠄⠒⠈⠀⠀⠀⠀⠀⠀⢠⠏⠀⠁⠉⠀⠈⠛⢥⣠⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
 ⠀⠀⠀⠀⠀⢸⣀⣀⣀⣀⠀⠀⠀⠀⠀⠀⠀⡰⠁⡠⠖⠋⢠⠃⠀⠀⠀⠀⠀⠀⠀⢀⢔⡏⠀⠀⠀⠀⠸⣷⠠⡀⠙⣆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
 ⠀⠀⣠⠔⠋⠉⠀⠀⠀⠈⠉⠙⠒⠦⣄⡀⢠⢣⠊⠀⠀⢀⠆⠀⠀⠀⢀⠊⠀⢠⢁⡎⡸⠀⠀⠀⠀⠀⠀⠈⠉⠁⠀⠘⢦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
 ⠀⡼⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⡿⣅⠀⠀⠀⡼⠀⠀⢀⡔⠁⠀⣠⡇⡞⢸⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠩⣒⠄⠀⠀⠀⠀⠀⠀⠀⠀
 ⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠳⡀⠀⡇⠀⢠⡞⠀⢀⠔⢹⡎⠀⡞⠀⠀⠀⠀⠀⠀⢠⡒⠒⠷⠿⠷⠒⠚⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀
 ⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⣄⠇⢀⢻⢃⣴⠁⠀⢨⠇⣸⠁⠀⠀⠀⠀⠀⠀⢸⠉⠲⢤⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
 ⠀⠹⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⡶⡈⠸⠃⡇⢰⠀⡜⢀⠇⠀⠀⠀⠀⠀⠀⠀⢸⠀⠀⠀⠙⠦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
 ⠀⠀⠹⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⠀⠑⢄⠀⠀⢳⠀⠀⣸⢁⢆⡼⡡⡞⠀⠀⠀⠀⠀⠀⠀⠀⠈⣇⠀⠀⠀⠀⠈⢣⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀
 ⠀⠀⣠⠼⣆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⡄⠀⠀⢣⠀⣼⠰⢠⣇⠮⠞⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⡄⠀⠀⠀⠀⠀⠱⡄⠀⠀⠀⠀⠀⠀⠀⠀
 ⠀⡞⠁⠀⠈⢣⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣗⠲⢤⡀⣃⣷⣡⢷⡖⠂⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡀⢷⢄⠀⠀⠀⠀⠀⠰⡀⠀⠀⠀⠀⠀⠀⠀
 ⠰⡇⠀⠀⠀⠀⠙⢦⠀⠀⠀⠀⠀⠀⠀⠀⠸⣀⡴⠋⠁⠀⢠⠎⠀⠀⠀⠀⠀⢀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠁⠸⠀⠁⠢⡀⠀⠀⠀⠘⢄⠀⠀⠀⠀⠀⠀
 ⠀⢧⠀⠀⠀⠀⠀⠀⠑⢦⡀⠀⠀⠀⠀⠀⠀⢻⡀⠀⠀⠀⢸⡠⡶⢀⡜⡰⣡⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⠇⠀⠀⠀⠈⠲⣄⡀⠀⠈⠳⢄⠀⠀⠀⠀
 ⠀⠘⣆⠀⠀⠀⠀⠀⠀⠀⠙⠢⣄⠀⠀⠀⠀⠀⠳⡄⠀⠀⠘⠁⠗⠃⢿⡟⠙⣿⡈⡄⠀⠀⠀⠀⡄⡀⣰⡟⠀⠀⠀⠀⠀⠀⠈⢯⡑⠢⠄⣀⠑⠢⢀⣀
 ⠀⠀⠈⢦⡀⠀⠀⠀⠀⠀⠀⠀⠈⢿⠦⣀⠀⠀⠀⠙⢦⠀⠀⠀⠀⠀⠀⡇⠀⠐⢣⢹⡀⠀⡆⣼⣴⠗⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⢳⡀⠀⠀⠀⠀⠀⠀
 ⠀⠀⠀⡜⠻⣄⠀⠀⠀⠀⠀⠀⠀⠸⡴⠚⡟⠢⢄⣀⠀⠳⣄⠀⡟⢲⠒⡇⠀⠀⠀⢑⡇⡼⠛⠉⠁⡼⢡⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢳⡀⠀⠀⠀⠀⠀
 ⠀⠀⢸⠁⠀⠈⠳⣄⠀⠀⠀⠀⠀⠀⢳⡀⢷⠀⠀⠈⠙⠒⠬⣿⣇⣸⡀⡇⠀⠀⢀⠏⢹⠀⠀⠀⣰⠃⠀⠄⠀⠀⠀⠀⠀⠀⠀⠀⠀⡀⣇⠀⠀⠀⠀⠀
 ⠀⠀⢸⠀⠀⠀⠀⠈⠳⣄⠀⠀⠀⠀⠀⠳⣸⠀⠀⠀⠀⠀⣠⠃⠀⢨⣧⡇⠀⠀⡼⠀⡟⠀⠀⢠⠇⠀⠀⠘⡀⠀⠀⠀⠀⠀⠀⠀⠀⢀⢸⠀⠀⠀⠀⠀
 ⠀⠀⠈⢧⠀⠀⠀⠀⠀⠈⢳⢦⡀⠀⠀⠀⠙⢆⠀⠀⣀⠔⠁⢀⡠⢛⣿⠀⠀⢰⠃⢀⡇⠀⠀⡞⠀⠀⠀⠀⠹⡢⢀⠀⠀⠀⠀⠀⠀⠘⣼⠀⠀⠀⠀⠀
 ⠀⠀⠀⠈⢳⡀⠀⠀⠀⠀⠈⢧⠉⣳⣤⣀⠀⠈⠳⣮⡥⠴⠚⠉⠀⢸⡏⠀⢀⡟⡀⢸⠁⠀⢸⠁⠀⠀⠀⠀⠀⠱⡀⠑⢤⡀⠀⠀⠀⠀⠹⡄⠀⠀⠀⠀
 ⠀⠀⠀⠀⠀⠙⠢⣄⠀⠀⠀⠈⢿⡁⢀⡏⠑⢢⢤⣈⠳⢄⡀⠀⠀⢸⠇⠀⡼⠀⠑⢿⠀⠀⣼⠀⠀⠀⠀⠀⠀⠀⠙⣄⠀⠈⠓⠦⣄⡀⠀⠑⢄⠀⠀⠀
 ⠀⠀⠀⠀⠀⠀⠀⠀⠙⠲⠤⣀⣀⠙⢼⡀⠀⣿⠀⠈⠉⠓⠻⠶⣄⣸⠀⢠⣧⡀⠀⠈⡇⠀⢸⠈⢲⢄⡀⠀⠀⠀⠀⠈⠳⣄⠀⠀⠀⠈⠉⠐⠒⠓⠄⠀
 ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠉⡿⢷⡤⣿⢤⠀⠀⠀⠀⢀⡤⣧⣄⣼⣽⣽⠀⠀⢳⠀⠀⢳⡢⠧⠌⠒⠤⢄⣀⡀⠀⠈⠑⠤⢀⣀⣀⠠⠀⠀⠀⠀
 ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠧⡼⠧⠼⠼⠀⠀⠀⠀⢿⣰⡇⢸⡆⣹⠀⠀⠀⠸⡄⠀⢄⢹⠓⡄⠀⠀⠀⠀⠈⠉⠉⠉⠉⠉⠁⠀⠀⠀⠀⠀⠀
 ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠉⠉⠉⠁⠀⠀⠀⠀⠉⠙⠛⠛⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀-???
`

	// https://www.asciiart.eu/mythology/devils
	devil string = `

    ,    ,    /\   /\
   /( /\ )\  _\ \_/ /_
   |\_||_/| < \_   _/ >
   \______/  \|0   0|/
     _\/_   _(_  ^  _)_
    ( () ) /'\|V"""V|/'\
      {}   \  \_____/  /
      ()   /\   )=(   /\
 jgs  {}  /  \_/\=/\_/  \
 `

	// https://www.asciiart.eu/mythology/phoenix
	phoenix string = `

                 (                           )
           ) )( (                           ( ) )( (
        ( ( ( )  ) )                     ( (   (  ) )(
       ) )     ,,\\\                     ///,,       ) (
    (  ((    (\\\\//                     \\////)      )
     ) )    (-(__//                       \\__)-)     (
    (((   ((-(__||                         ||__)-))    ) )
   ) )   ((-(-(_||           '''\__        ||_)-)-))   ((
   ((   ((-(-(/(/\\        ''; 9.- '      //\)\)-)-))    )
    )   (-(-(/(/(/\\      '';;;;-\~      //\)\)\)-)-)   (   )
 (  (   ((-(-(/(/(/\======,:;:;:;:,======/\)\)\)-)-))   )
     )  '(((-(/(/(/(//////:%%%%%%%:\\\\\\)\)\)\)-)))'  ( (
    ((   '((-(/(/(/('uuuu:WWWWWWWWW:uuuu')\)\)\)-))'    )
      ))  '((-(/(/(/('|||:wwwwwwwww:|||')\)\)\)-))'    ((
   (   ((   '((((/(/('uuu:WWWWWWWWW:uuu')\)\))))'     ))
         ))   '':::UUUUUU:wwwwwwwww:UUUUUU:::''     ((   )
           ((      '''''''\uuuuuuuu/''''''         ))
            ))            'JJJJJJJJJ'           ((
              ((            LLLLLLLLLLL         ))
                ))         ///|||||||\\\       ((
                  ))      (/(/(/(^)\)\)\)       ((
                   ((                           ))
                     ((                       ((
                       ( )( ))( ( ( ) )( ) (()      -???
	`

	// https://www.asciiart.eu/mythology/unicorns
	unicorn string = `

 \.
  \\      .
   \\ _,.+;)_
   .\\;~%:88%%.
  (( a   ')9,8;%.
  /'   _) ' '9%%%?
 (' .-' j    '8%%'
  '"+   |    .88%)+._____..,,_   ,+%$%.
        :.   d%9'             '-%*'"'~%$.
     ___(   (%C                 '.   68%%9
   ."        \7                  ;  C8%%)'
   : ."-.__,'.____________..,'   L.  \86' ,
   : L    : :            '  .'\.   '.  %$9%)
   ;  -.  : |             \  \  "-._ '. '~"
    '. !  : |              )  >     ". ?
      ''  : |            .' .'       : |
          ; !          .' .'         : |
         ,' ;         ' .'           ; (
        .  (         j  (            '  \
        """'          ""'             '"" mh`
 
 // https://www.asciiart.eu/mythology/grim-reapers
	grimReaper string = `
 
              ___
             /   \\
        /\\ | . . \\
      ////\\|     ||
    ////   \\ ___//\
   ///      \\      \
  ///       |\\      |
 //         | \\  \   \
 /          |  \\  \   \
            |   \\ /   /
            |    \/   /
            |     \\/|
            |      \\|
            |       \\
            |        |
            |_________\
      from Dustin Slater
 `
 
 // https://www.asciiart.eu/mythology/gryphon
	gryphon string = `
 
                         ______
              ______,---'__,---'
          _,-'---_---__,---'
   /_    (,  ---____',
  /  ',,   ', ,-'
 ;/)   ,',,_/,'
 | /\   ,.'//\
 '-' \ ,,'    '.
      '',   ,-- '.
      '/ / |      ',         _
      //'',.\_    .\\      ,{==>-
   __//   __;_'-  \ ';.__,;'
 ((,--,) (((,------;  '--' jv
 '''  '   ''',
 `
 
 // https://www.asciiart.eu/mythology/centaurs
	centaur string = `
 
   <=======]}======
     --.   /|
    _\"/_.'/
  .'._._,.'
  :/ \{}/
 (L  /--',----._
     |          \\
    : /-\ .'-'\ / |
 snd \\, ||    \|
      \/ ||    ||`

	// https://www.asciiart.eu/animals/spiders
	spider string = `

               (
                )
               (
         /\  .-"""-.  /\
        //\\/  ,,,  \//\\
        |/\| ,;;;;;, |/\|
        //\\\;-"""-;///\\
       //  \/   .   \/  \\
      (| ,-_| \ | / |_-, |)
        //'__\.-.-./__'\\
       // /.-(() ())-.\ \\
      (\ |)   '---'   (| /)
       ' (|           |) '
 jgs     \)           (/'`

	// https://www.asciiart.eu/animals/bears
	bear string = `

      (()__(()
      /       \
     ( /    \  \
      \ o o    /
      (_()_)__/ \
     / _,==.____ \
    (   |--|      )
    /\_.|__|'-.__/\_
   / (        /     \
   \  \      (      /
    )  '._____)    /
 (((____.--(((____/mrf`

	// https://www.asciiart.eu/animals/scorpions
	scorpion string = `

      ( _<    >_ )
      //        \\
      \\___..___//
       '-(    )-'
         _|__|_
        /_|__|_\
        /_|__|_\
        /_\__/_\
         \ || /  _)
           ||   ( )
      Max  \\___//
            '---'`

	// https://www.asciiart.eu/animals/bats
	bat string = `

        (_    ,_,    _)
        / ''--) (--'' \
       /  _,-'\_/'-,_  \
 jgs  /.-'     "     '-.\`

	// https://www.asciiart.eu/mythology/phoenix
	raven string = `

    _,="( _  )"=,_
 _,'    \_>\_/    ',_
 .7,     {  }     ,\.
  '/:,  .m  m.  ,:\'
    ')",(/  \),"('
       '{'!!'}'      -???`
)
