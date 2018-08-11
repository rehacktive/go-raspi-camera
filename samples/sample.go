package main

import (
	"github.com/rehacktive/go-raspi-camera"
	"fmt"
	"os"
)

func main() {
	c := camera.New("/ramfs")
	s, err := c.Capture()
	if err != nil {
		fmt.Println("Error ", err)
		os.Exit(1)
	}	
	fmt.Println(s)
}	
