package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

// message represents a message object
type message struct {
	text      string
	iteration int
}

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
	iterateMessage(times, os.Args[2])
}

// iterageMessage fires printMessage in a goroutine.
// It then sends n messages immediately on the msgChan.
// It waits for the waitgroup to be done before exiting.
func iterateMessage(times int, messageStr string) {
	var wg sync.WaitGroup
	msgChan := make(chan message, times)
	wg.Add(times)
	go printMessage(msgChan, &wg)
	for i := 0; i < times; i++ {
		msgChan <- message{text: messageStr, iteration: i}
	}
	fmt.Println("All messages sent.")
	wg.Wait()
}

// printMessage ranges over the msgChan printing messages
// 100ms apart and calling Done() on the waitgroup.
func printMessage(msgChan chan message, wg *sync.WaitGroup) {
	for m := range msgChan {
		fmt.Printf("%d: %s\n", m.iteration, m.text)
		time.Sleep(time.Millisecond * 100)
		wg.Done()
	}
}
