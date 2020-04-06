package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage:\ncli <int> <string>")
		os.Exit(1)
	}
	times, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	printMessage(times, os.Args[2])
	if len(os.Args) > 3 {
		fmt.Printf("extra args: %v", os.Args[3:])
	}
}

func printMessage(times int, message string) {
	for i := 0; i < times; i++ {
		fmt.Printf("%d: %s\n", i, message)
	}
}
