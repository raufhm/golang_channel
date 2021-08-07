package main

import (
	"fmt"
	"golang_channel/util"
)

func main() {
	fmt.Println("Pick a goroutine example: (1, 2, 3)")
	var input int
	fmt.Scanln(&input)

	if input == 1 {
		util.ChannelwithoutBlocking()
	} else if input == 2 {
		util.ChannelWithBlocking()
	} else if input == 3 {
		var messages = make(chan string)

		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			messages <- data
		}("wick")

		var message = <-messages
		fmt.Println(message)
	} else {
		fmt.Println("incorret input")
	}

}
