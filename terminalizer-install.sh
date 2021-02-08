#!/bin/bash

# https://github.com/faressoft/terminalizer

sudo apt install npm
npm install terminalizer

# https://github.com/faressoft/terminalizer/issues/150#issuecomment-775487720
echo ""
echo "Patch node_modules/terminalizer/render/index.js to include the following after the var block:"
echo ""
echo 'app.disableHardwareAcceleration();'
echo 'app.commandLine.appendSwitch("disable-software-rasterizer");'