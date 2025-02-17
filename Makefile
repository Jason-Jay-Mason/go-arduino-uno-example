#add path to your usb port for the arduino (example: /dev/cu.usbserial-210)
PORT=add port here

#Build to hex
build:
	tinygo build -o ./bin/main.hex -target=arduino main.go
flash: 
	tinygo flash -target=arduino -port=$(PORT)
