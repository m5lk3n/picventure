# RPG

The RPG text adventure from https://projects.raspberrypi.org/en/projects/rpg implemented in go.

View [original base source](https://rpf.io/rpg-code) [locally](originals/rpg-rpg.py).

![Final Map](originals/rpg-final-map.png "Final Map")

[source](https://projects-static.raspberrypi.org/projects/rpg/31fb9012c6d897ad16f2f245fb4791b6384cda28/en/images/rpg-final-map.png)

## prereqs

`$ go get github.com/deckarep/golang-set`

## to do

- makefile
- introduce enums
- introduce constants
- improve `inventory.Contains` check

## open

- `input = strings.TrimRight`?
- support Windows by using `input = strings.Trim(input, "\r\n")`

## ideas

- get random item like a skin
- add-on: visualization of position