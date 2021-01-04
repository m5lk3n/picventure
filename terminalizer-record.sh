#!/bin/bash

# https://github.com/faressoft/terminalizer

read -r -p "Prompt reset and ready to record? [y/N] " response
case "$response" in
    [yY])
        clear
        BIN=./node_modules/.bin/terminalizer
        $BIN record spoiler --config terminalizer-config.yml
        $BIN render spoiler --output spoiler.gif
        ;;
esac