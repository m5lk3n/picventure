# Picventure

The RPG text adventure project from under [raspberrypi.org](https://projects.raspberrypi.org/en/projects/rpg) implemented in Go for educational purposes. It features colors and pictures (on a Sense HAT).

## Environments

Developed on a [Raspberry Pi 400](https://www.raspberrypi.org/products/raspberry-pi-400/):

```bash
$ uname -a
Linux raspi400 5.8.0-1010-raspi #13-Ubuntu SMP PREEMPT Wed Dec 9 17:14:07 UTC 2020 aarch64 aarch64 aarch64 GNU/Linux
$ go version
go version go1.14.7 linux/arm64
```

Also tested on a Lenovo ThinkPad X250:

```bash
$ uname -a
Linux x250 4.15.0-128-generic #131-Ubuntu SMP Wed Dec 9 06:57:35 UTC 2020 x86_64 x86_64 x86_64 GNU/Linux
$ go version
go version go1.15.6 linux/amd64
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
go get github.com/pterm/pterm
go get github.com/sirupsen/logrus
go build -o rpg
```

## Play

```bash
$ ./rpg 
Picventure - A colorful RPG Adventure with Pictures
===================================================

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

## At work

- call picservice in rpg
- renew screenshots

## To do

- introduce [enums](https://www.ribice.ba/golang-enums/)
- introduce constants
- improve `inventory.Contains` check

## Open

- tests
- `input = strings.TrimRight`?
- support Windows by using `input = strings.Trim(input, "\r\n")`

## Ideas

- save/load game
- 3 dimensions: up/down via stairs
- get random item like a skin
- add-on: visualization of position
