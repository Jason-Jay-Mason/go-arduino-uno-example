# Arduino Go Example
This is an example/tutorial for using TinyGo to program Arduino boards. TinyGo is a Go compiler for small places, making it possible to use the Go programming language for microcontrollers and embedded systems. This tutorial demonstrates how to write Go code that can be compiled to run on Arduino hardware, combining the simplicity and safety features of Go with embedded systems development.

# Prerequisites

- Go (latest stable version)
- TinyGo with Avrdude
- VS Code (recommended) or your preferred IDE
- Git
- Arduino Uno board

# Help to install prerequisites
- Make sure you have Go installed:
https://go.dev/doc/install

- Install TinyGo:
https://tinygo.org/getting-started/install/
(Make sure you install Avrdude, don't skip that step)

- If you don't have an IDE, install VS Code:
https://code.visualstudio.com/

- Make sure you have git installed:
https://github.com/git-guides/install-git

# Project setup
- Clone the code down to your machine by first navigating to where you would like to save the project in your terminal, then:
```bash
git clone https://github.com/Jason-Jay-Mason/tinygo-arduinouno.git
```


# Editor config
You should use gopls for LSP support here. The problem is that gopls is configured for Go, not TinyGo. To inject the proper TinyGo packages into your editor, take a look at this documentation:
https://tinygo.org/docs/guides/ide-integration/

## VS Code
VS code is fairly strait forward in the documentation
- Install the Go extension in VS code.
- Get the tags by following this: https://tinygo.org/docs/guides/ide-integration/
- Use the tags in your VS code configuration: https://tinygo.org/docs/guides/ide-integration/vscode/

## Vim/NeoVim (for advanced people that happen upon this repo)
### Docs
https://tinygo.org/docs/guides/ide-integration/vim-neovim/

### plugin suggestion for Neovim
This is kind of a hacky solution, but it works.

Install this plugin with your nvim plugin manager of choice:
https://github.com/pcolladosoto/tinygo.nvim

Then in nvim run:
```vim
:TinyGoSetTarget arduino
```
You're all set!

# Developing your Arduino program
It's a good idea to use a simulator during development like https://wokwi.com

## Wokwi
- Build the project:
```bash
make build
```
- Go to https://wokwi.com and log in
- Create a new project for Arduino
- Delete the sketch.ino file
- Press the plus button and add an LED to the simulation window
- Connect the cathode of the LED to the GND on the Arduino 
- Connect the anode of the LED to pin 1 on the Arduino
```
GND -> cathode (smaller wire on LED)
pin 1 -> anode (longer wire on LED)

```
- Click on the diagram.json file
- Press F1 to open command prompt search
- Search for "Upload Firmware and Start Simulation"
- Select "Upload Firmware and Start Simulation"
- Select the hex file that you built in step 1 at "./your/root-dir/of/project/.bin/main.hex"

You should see a flashing LED! Happy hacking!

# Flash to Arduino
When you are happy in the simulator, flash your code to your arduino uno by:

- Plug in your arduino via usb, then run:
```bash
 tinygo ports
```
- Find the /dev/{your-device} path that maps to your arduino device.
(TIP: If you're not sure which /dev/* path maps to your arduino, try unpluging your arduino and running the command again, then use process of elimination to find the proper path.)
- Copy the port path (eg /dev/cu.usbserial-210) and paste it into the Makefile var 'PORT'
```Makefile
#add path to your usb port for the arduino in Makefile
PORT=/dev/cu.usbserial-210
```
- Run 
```bash
make flash
```

And that should do it!
