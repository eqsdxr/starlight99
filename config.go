package main

type Config struct {
	PlayerName string
	TotalScore int
}

type rawMonster struct {
	ascii, name                        string
	hpMin, hpMax, damageMin, damageMax int
}

type Monster struct {
	ascii, name              *string
	hp, damageMin, damageMax int
}

type Expression struct {
	first, second, result int
}

var (

	configPath string = "/home/ivan/.local/share/starlight99.toml"

	cliName string = "starlight99"
	initialName = "AnonymousPlayer"

	Reset = "\033[0m"
	Red = "\033[31m"
	Green = "\033[32m"
	Yellow = "\033[33m"
	Blue = "\033[34m"
	Magenta = "\033[35m"
	Cyan = "\033[36m"
	Gray = "\033[37m"
	White = "\033[97m"

	startingMenu string = `
Welcome to Starlight99 my dear wanderer! Choose an option:
1. Play
2. Settings
3. Exit
`
	reallyWannaExit string = `
Do you want to exit the game?
1. Yes
2. No
	`
	exitMessage string = `
Exiting...
	`
	gameModeMenu string = `
Choose the game mode:
1. Adventure
2. Company
3. Go back
	`
	changingCharacterName string = `
Write new name for your character:
	`
	settingsMenu string = `
1. Change the character's name
2. Change config file location
3. View stats
4. About
5. Go back
	`
	helpMessage string = `
Available commands:
\\help - show this text
\\exit - exit the program
	`
	startingAdventureModeText string = `
You're starting an adventure mode.
	`

	monsters1 = []rawMonster{bearMonster, batMonster, scorpionMonster, spiderMonster, ravenMonster}
	monsters2 = []rawMonster{centaurMonster, gryphonMonster, grimReaperMonster, unicornMonster, phoenixMonster, devilMonster}
	monsters3 = []rawMonster{starDevilMonster, skeletonMonster, dragonMonster, foxMonster}

	allMonsters = func() []rawMonster {
		all := append(monsters1, monsters2...)
		return append(all, monsters3...)
	}

	starDevilMonster = rawMonster{starDevil, "Star Devil", 100000, 100000, 100000, 100000}
	skeletonMonster = rawMonster{skeleton, "Skeleton", 40000, 70000, 15000, 30000}
	dragonMonster = rawMonster{dragon, "Dragon", 10000, 30000, 5000, 7000}
	foxMonster = rawMonster{fox, "Fox", 5000, 10000, 1000, 5000}

	devilMonster = rawMonster{devil, "Devil", 40000, 70000, 15000, 30000}
	phoenixMonster = rawMonster{phoenix, "Phoenix", 40000, 70000, 15000, 30000}
	unicornMonster = rawMonster{unicorn, "Unicorn", 40000, 70000, 15000, 30000}
	grimReaperMonster = rawMonster{grimReaper, "Grim Reaper", 40000, 70000, 15000, 30000}
	gryphonMonster = rawMonster{gryphon, "Gryphon", 40000, 70000, 15000, 30000}
	centaurMonster = rawMonster{centaur, "Centaur", 40000, 70000, 15000, 30000}

	spiderMonster = rawMonster{spider, "Spider", 500, 1500, 70, 120}
	bearMonster = rawMonster{bear, "Bear", 2000, 5000, 200, 400}
	scorpionMonster = rawMonster{scorpion, "Scorpion", 150, 350, 30, 70}
	ravenMonster = rawMonster{raven, "Raven", 40000, 70000, 15000, 30000}
	batMonster = rawMonster{bat, "Bat", 50, 150, 5, 15}

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
	                                   VV
	
	`
	
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
	                                             |||||  -nabis
	
	`
	
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
	                  (vvv(VVV)(VVV)vvv)
	                                                -???
	`
	
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
	       """'          ""'             '"" mh
	
	`

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
	'''  '   '''
	
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
	     \/ ||    ||
	
	`

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
	jgs     \)           (/'
	
	`
	
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
	(((____.--(((____/mrf
	
	`

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
		            '---'
	
	`

	// https://www.asciiart.eu/animals/bats
	bat string = `
	
		       (_    ,_,    _)
		       / ''--) (--'' \
		      /  _,-'\_/'-,_  \
		jgs  /.-'     "     '-.\
		
	`
	
	// https://www.asciiart.eu/mythology/phoenix
	raven string = `
	
	   _,="( _  )"=,_
	_,'    \_>\_/    ',_
	.7,     {  }     ,\.
	 '/:,  .m  m.  ,:\'
	   ')",(/  \),"('
	      '{'!!'}'      -???
	
	`
)
