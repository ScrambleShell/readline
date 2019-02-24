package main

import (
	"fmt"
	"strings"

	"github.com/lmorg/murex/utils/readline"
)

func main() {
	rl := readline.NewInstance()
	rl.TabCompleter = Tab

	for {
		s, err := rl.Readline()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Readline: '" + s + "'")
	}
}

// Tab is the tab-completion handler for this readline example program
func Tab(line []rune, pos int) (string, []string, map[string]string, readline.TabDisplayType) {
	items := []string{
		"aa",
		"abaya",
		"abomasum",
		"absquatulate",
		"adscititious",
		"afreet",
		"Albertopolis",
		"alcazar",
		"amphibology",
		"amphisbaena",
		"anfractuous",
		"anguilliform",
		"apoptosis",
		"apple-knocker",
		"argle-bargle",
		"Argus-eyed",
		"argute",
		"ariel",
		"aristotle",
		"aspergillum",
		"astrobleme",
		"Attic",
		"autotomy",
		"badmash",
		"bandoline",
		"bardolatry",
		"Barmecide",
		"barn",
		"bashment",
		"bawbee",
		"benthos",
		"bergschrund",
		"bezoar",
		"bibliopole",
		"bichon",
		"bilboes",
		"bindlestiff",
		"bingle",
		"blatherskite",
		"bleeding",
		"blind",
		"bobsy-die",
		"boffola",
		"boilover",
		"borborygmus",
		"breatharian",
		"Brobdingnagian",
		"bruxism",
		"bumbo",
	}

	var suggestions []string

	for i := range items {
		if strings.HasPrefix(items[i], string(line)) {
			suggestions = append(suggestions, items[i][pos:])
		}
	}

	return string(line[:pos]), suggestions, nil, readline.TabDisplayGrid
}
