package main

import (
	"fmt"
	"os"

	camera "github.com/rehacktive/go-raspi-camera"
)

func main() {
	c := camera.New("/ramfs", 800, 600)
	s, err := c.Capture()
	if err != nil {
		fmt.Println("Error ", err)
		os.Exit(1)
	}
	fmt.Println(s)
}
