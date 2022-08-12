package main

import (
	"fmt"
	"gabiparser/cmd"
)

func main() {
	fmt.Println("Start Processing...")
	err := cmd.Execute()
	if err != nil {
		return
	}
}
