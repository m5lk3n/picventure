# default target
.PHONY: help
help:
	@echo "usage: make <target>"
	@echo
	@echo "  where <target> is one of the following"
	@echo
	@echo "    get         to fetch all package dependencies"
	@echo "    build       to compile binary for local machine architecture"
	@echo "    all         to run all targets"
	@echo
	@echo "    help        to show this text"

.PHONY: get
get:
	go get github.com/gin-gonic/gin
	go get github.com/nathany/bobblehat/sense/screen
	go get github.com/nathany/bobblehat/sense/screen/texture
	go get github.com/sirupsen/logrus

.PHONY: build
build:
	go build -o picservice

.PHONY: all
all: get build
