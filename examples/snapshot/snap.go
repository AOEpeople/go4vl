package main

import (
	"context"
	"log"
	"os"

	"github.com/vladimirvivien/go4vl/device"
)

var callback = func(error) {
	return
}

func main() {
	dev, err := device.Open("/dev/video0", device.WithBufferSize(1))
	if err != nil {
		log.Fatal(err)
	}
	defer dev.Close()

	if err := dev.Start(context.TODO(), callback); err != nil {
		log.Fatal(err)
	}

	frame := <-dev.GetOutput()

	file, err := os.Create("pic.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := file.Write(frame); err != nil {
		log.Fatal(err)
	}
}
