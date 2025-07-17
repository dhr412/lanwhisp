package main

import (
	"fmt"
	"time"
)

func main() {
	name, port, passphrase := ParseArgs()

	salt := "lan-chat-static-salt"

	key, err := DeriveKey(passphrase, salt)
	if err != nil {
		fmt.Println("Error deriving key:", err)
		return
	}

	go StartServer(port, key, name)

	time.Sleep(500 * time.Millisecond)

	fmt.Println("Encrypted LAN chat started.")
	fmt.Println("Enter IP:PORT to send messages.")

	InputLoop(key, name)
}
