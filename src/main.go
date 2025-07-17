package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	cfg := ParseArgs()
	portStr := strconv.Itoa(cfg.Port)

	salt := cfg.Name + portStr + cfg.Passphrase
	key := DeriveKey(cfg.Passphrase, salt)
	// if err != nil {
	//	fmt.Println("Error deriving key:", err)
	//	return
	// }

	go StartServer(portStr, key, cfg.Name)

	time.Sleep(512 * time.Millisecond)

	fmt.Println("Encrypted LAN chat started.")
	fmt.Println("Enter IP:PORT to send messages.")

	InputLoop(key, cfg.Name)
}
