package main

import (
	"fmt"
	"math/rand"
	"time"
)

var goodPattern = "%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%+"
var badPattern = "%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%-"

func main() {
	fmt.Println("Launching Master Industries 4.0 Antivirus Service....")
	if goodPattern[:128] == badPattern[:128] {
		good()
	} else {
		bad()
	}
}

func good() {
	time.Sleep(1 * time.Second)
	fmt.Println("Antivirus service OK. ALL GOOD.")
}

func bad() {
	time.Sleep(1 * time.Second)

	for i := 1; i < 100; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("%c", byte(rand.Intn(255)))
	}

	fmt.Println("")
	fmt.Print("Encrypting main disk.")
	for i := 1; i < 5; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Print("\n")
	fmt.Println("Don't reboot, your computer is now encrypted. Pay 0.3 bitcoins to 1PC9aZC4hNX2rmmrt7uHTfYAS3hRbph4UN for decryption key.")

}
