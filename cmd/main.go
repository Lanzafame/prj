package main

import (
	"fmt"

	"github.com/Lanzafame/prj"
)

func main() {
	fmt.Println("Creating prj config directory...")
	err := prj.SetupPrjDir()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Config directory successfully created.")
}
