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

func printInventory() {
	it := inventory.Iterator()

	for elem := range it.C {
		fmt.Print(elem)
	}
}
func showStatus() {
	// print the player's current status
	fmt.Println("---------------------------")
	fmt.Println("You are in the " + currentRoom)
	fmt.Println("Inventory: ")
	printInventory()
}

type dir2Room struct {
	direction string
	room      string
}

var inventory = mapset.NewSet()

var rooms = map[string]dir2Room{
	"Hall":    dir2Room{"south", "Kitchen"},
	"Kitchen": dir2Room{"north", "Hall"},
}

func handleGo(dir string) {
	fmt.Println("debug")
	if d2R, ok := rooms[currentRoom]; ok {
		newRoom := d2R
		fmt.Println("new ", newRoom)
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
