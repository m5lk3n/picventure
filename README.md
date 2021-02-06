# Private go-rpg

The RPG text adventure project from under [raspberrypi.org](https://projects.raspberrypi.org/en/projects/rpg) implemented in Go for educational purposes.

## Development environment

Developed on a [Raspberry Pi 400](https://www.raspberrypi.org/products/raspberry-pi-400/):

```bash
$ uname -a
Linux raspi400 5.8.0-1010-raspi #13-Ubuntu SMP PREEMPT Wed Dec 9 17:14:07 UTC 2020 aarch64 aarch64 aarch64 GNU/Linux
$ go version
go version go1.14.7 linux/arm64
```

## Make

```bash
$ make
usage: make <target>

  where <target> is one of the following

    get         to fetch all dependencies
    build       to compile binary for local machine architecture
    all         to run get, build

    help        to show this text
```

```bash
$ make all
go get github.com/deckarep/golang-set
go build -o rpg
```

## Play

```bash
$ ./rpg 
RPG Game
========

Get to the Garden with a key and a potion
Avoid the monsters!

Commands:
  go "direction"    (where "direction" is one of the following: north, east, south, west)
  get "item"
  exit
---------------------------
You are in the Hall
You see a key
Inventory:
---------------------------
> 
```

## Demo

**Spoiler alert!**

![Make and play](spoiler.gif)

## Originals

View [original base source](https://rpf.io/rpg-code) [locally](originals/rpg-rpg.py).

The [map](https://projects-static.raspberrypi.org/projects/rpg/31fb9012c6d897ad16f2f245fb4791b6384cda28/en/images/rpg-final-map.png):

![Final Map](originals/rpg-final-map.png "Final Map")

## To do

- introduce [enums](https://www.ribice.ba/golang-enums/)
- introduce constants
- improve `inventory.Contains` check

## Open

- tests
- `input = strings.TrimRight`?
- support Windows by using `input = strings.Trim(input, "\r\n")`

## Ideas

- use [PTerm](https://github.com/pterm/pterm)
- [lights-on](ideas/lights-on.py)
- save/load game
- 3 dimensions: up/down via stairs
- get random item like a skin
- add-on: visualization of position
