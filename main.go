package main

import (
	"bufio"
	"flag"
	"os"
	"strings"
	"time"

	set "github.com/deckarep/golang-set"
	"github.com/pterm/pterm"
	log "github.com/sirupsen/logrus"
)

func init() {
	debugFlag := flag.Bool("debug", false, "turn debug output on/off")

	flag.Parse()

	if *debugFlag {
		log.SetLevel(log.DebugLevel)
	}
}

func showInstructions() {
	pterm.DefaultHeader.Println("RPG Game")
	pterm.FgLightCyan.Println(`Get to the Garden with a key and a potion
Avoid the monsters!

Commands:
  go "direction"    (where "direction" is one of the following: north, east, south, west)
  get "item"
  exit`)
}

func printRoom() {
	pterm.FgLightWhite.Println("You are in the ", currentRoom)
	if roomItem, ok := items[currentRoom]; ok {
		pterm.Println("You see a ", roomItem)
	}
}

func printInventory() {
	pterm.FgLightWhite.Print("Inventory:")

	items := inventory.Iterator()
	for item := range items.C {
		pterm.FgLightWhite.Print(" ", item)
	}

	pterm.Println()
}

func showStatus() {
	pterm.FgLightYellow.Println("---------------------------")
	printRoom()
	printInventory()
	pterm.FgLightYellow.Println("---------------------------")
}

func confirmExit() bool {
	pterm.Print("Do you really want to exit the Game? [y/n] ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.ToLower(strings.TrimSpace(input))

	return input == "y" || input == "yes"
}

type direction2Room map[string]string

var inventory = set.NewSet()

var rooms = map[string]direction2Room{
	"Hall":        direction2Room{"south": "Kitchen", "east": "Dining Room"},
	"Kitchen":     direction2Room{"north": "Hall"},
	"Dining Room": direction2Room{"west": "Hall", "south": "Garden"},
	"Garden":      direction2Room{"north": "Dining Room"},
}

var items = map[string]string{
	"Hall":        "key",
	"Kitchen":     "monster",
	"Dining Room": "potion",
}

func handleGo(direction string) {
	spinner, err := pterm.DefaultSpinner.Start("Going...")
	if err != nil {
		log.Fatalln(err)
	}
	time.Sleep(time.Second * 1) // simulate 1 second of activity...

	if newRoom, ok := rooms[currentRoom][direction]; ok {

		currentRoom = newRoom

		// player loses if they enter a room with a monster
		if roomItem, ok := items[currentRoom]; ok {
			if roomItem == "monster" {
				pterm.Error.Prefix = pterm.Prefix{
					Text:  "GAME OVER!",
					Style: pterm.NewStyle(pterm.BgRed, pterm.FgWhite),
				}
				spinner.Fail("A monster has got you...")
				os.Exit(0)
			}
		}

		if currentRoom == "Garden" {
			if inventory.Contains("key") && inventory.Contains("potion") {
				pterm.Success.Prefix = pterm.Prefix{
					Text:  "YOU WIN!",
					Style: pterm.NewStyle(pterm.BgGreen, pterm.FgWhite),
				}
				spinner.Success("You escaped the house...")
				os.Exit(0)
			}
		}

		spinner.Success(" ")
		return
	}

	pterm.Error.Prefix = pterm.Prefix{
		Text:  "NO LUCK",
		Style: pterm.NewStyle(pterm.BgRed, pterm.FgBlack),
	}
	spinner.Fail("You can't go that way!")
}

func handleGet(item string) {
	if roomItem, ok := items[currentRoom]; ok {
		if item == roomItem {
			pterm.FgLightGreen.Println("Got the item!")
			inventory.Add(item)
			delete(items, currentRoom)
			return
		}
	}

	pterm.FgLightRed.Printf("Can't get %s!\n", item)
}

func handleExit() {
	if confirmExit() {
		os.Exit(0)
	}
}

// start the player in the Hall
var currentRoom = "Hall"

func main() {
	showInstructions()
	reader := bufio.NewReader(os.Stdin)

	for {
		showStatus()

		pterm.FgLightGreen.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		input = strings.Trim(input, " \n")
		input = strings.ToLower(input)
		move := strings.Split(input, " ")

		verb := ""
		noun := ""

		if len(move) > 0 {
			verb = move[0]
			if len(move) > 1 {
				noun = move[1]
			}
		}

		switch verb {
		case "go":
			handleGo(noun)
		case "get":
			handleGet(noun)
		case "exit":
			handleExit()
		default:
			log.Debugln("command verb not understood")
		}
	}
}
