package util

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func ChannelWithBlocking() {
	//GOMAXPROCS digunakan untuk menentukan jumlah core yang diaktifkan untuk eksekusi program.
	runtime.GOMAXPROCS(2)

	go print(5, "halo")
	print(5, "apa kabar")

	var input string

	//Scanln mengakibatkan proses jalannya aplikasi berhenti di baris itu (blocking) hingga user menekan tombol enter
	fmt.Scanln(&input)
}

func ChannelwithoutBlocking() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data
	}

	go sayHelloTo("john wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("jason bourne")

	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)
}
