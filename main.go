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
Commands:
  go [direction]
  get [item]`)
}

func printRoom() {
	fmt.Println("You are in the " + currentRoom)
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
	"Hall":    direction2Room{"south": "Kitchen"},
	"Kitchen": direction2Room{"north": "Hall"},
}

func handleGo(direction string) {
	if newRoom, ok := rooms[currentRoom][direction]; ok {
		currentRoom = newRoom
	} else {
		fmt.Println("You can't go that way!")
	}
}

func handleGet(item string) {
	fmt.Println(item) // TODO: implement
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

		// input
		fmt.Print(">")
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
		case "go": // TODO: introduce constant
			handleGo(noun)
		case "exit":
			handleExit()
		default:
			// not understood
		}

	}
}
