# picservice

A simple API server that draws predefined pictures on the Pi's Sense HAT.

## Prerequisites

- Sense HAT set up (out of scope here)
- Dependencies installed:

```bash
$ make get
go get github.com/gin-gonic/gin
go get github.com/nathany/bobblehat/sense/screen
go get github.com/nathany/bobblehat/sense/screen/texture
go get github.com/sirupsen/logrus
```

## Build

```bash
$ make build
go build -o picservice
```

## Run

### Non-Production

```bash
$ ./picservice
...
[GIN-debug] [WARNING] Running in "debug" mode. ...
[GIN-debug] GET    /api/v1/screen/clear      --> main.ScreenClearHandler (4 handlers)
[GIN-debug] GET    /api/v1/screen/draw/:pic  --> main.ScreenDrawHandler (4 handlers)
INFO[0000] picservice start...                          
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
```

### Production

```bash
$ GIN_MODE=release ./picservice 
frame buffer device not found
INFO[0000] picservice start...
```

## Sources

- [Pixel art on Sense HAT](https://magpi.raspberrypi.org/articles/pixel-art-on-sense-hat)
- [Learn Sense HAT with Raspberry Pi](https://magpi.raspberrypi.org/articles/learn-sense-hat-with-raspberry-pi)
- [bobbleHAT](https://github.com/nathany/bobblehat)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
