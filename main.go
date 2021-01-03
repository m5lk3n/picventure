package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	mapset "github.com/deckarep/golang-set"
)

func showInstructions() {
	fmt.Println(`RPG Game
========

Get to the Garden with a key and a potion
Avoid the monsters!

Commands:
  go [direction]
  get [item]`)
}

func printRoom() {
	fmt.Println("You are in the", currentRoom)
	if roomItem, ok := items[currentRoom]; ok {
		fmt.Println("You see a", roomItem)
	}
}

func printInventory() {
	fmt.Print("Inventory:")

	items := inventory.Iterator()
	for item := range items.C {
		fmt.Print(" ", item)
	}

	fmt.Println()
}

func showStatus() {
	fmt.Println("---------------------------")
	printRoom()
	printInventory()
	fmt.Println("---------------------------")
}

type direction2Room map[string]string

var inventory = mapset.NewSet()

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
	if newRoom, ok := rooms[currentRoom][direction]; ok {
		currentRoom = newRoom

		// player loses if they enter a room with a monster
		if roomItem, ok := items[currentRoom]; ok {
			if roomItem == "monster" {
				fmt.Println("A monster has got you... GAME OVER!")
				os.Exit(0)
			}
		}

		if currentRoom == "Garden" && inventory.Contains("key") && inventory.Contains("potion") {
			fmt.Println("You escaped the house... YOU WIN!")
			os.Exit(0)
		}

		return
	}

	fmt.Println("You can't go that way!")
}

func handleGet(item string) {
	if roomItem, ok := items[currentRoom]; ok {
		if item == roomItem {
			inventory.Add(item)
			delete(items, currentRoom)
			return
		}
	}

	fmt.Printf("Can't get %s!\n", item)
}

func handleExit() {
	// TODO: add confirmation
	os.Exit(0)
}

// start the player in the Hall
var currentRoom = "Hall"

func main() {
	showInstructions()
	reader := bufio.NewReader(os.Stdin)

	for {
		showStatus()

		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
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
			// not understood
		}
	}
}
