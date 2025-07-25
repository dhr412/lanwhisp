package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type BroadcastPayload struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
}

func StartServer(port string, key []byte, name string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer ln.Close()
	fmt.Println("Listening on port", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go handleConnection(conn, key)
	}
}

func handleConnection(conn net.Conn, key []byte) {
	defer conn.Close()
	buffer := make([]byte, 4096)

	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	decrypted, err := Decrypt(buffer[:n], key)
	if err != nil {
		fmt.Println("Decryption error:", err)
		return
	}

	splitIndex := strings.Index(decrypted, ": ")
	if splitIndex == -1 {
		fmt.Println("\n[Message Received] →", decrypted)
		return
	}

	meta := decrypted[:splitIndex]
	msg := decrypted[splitIndex+2:]
	fmt.Printf("\n%s: %s\n", meta, msg)
}

func SendMessage(targetIP string, targetPort string, message string, key []byte, senderName string) error {
	conn, err := net.Dial("tcp", targetIP+":"+targetPort)
	if err != nil {
		return fmt.Errorf("connection error: %w", err)
	}
	defer conn.Close()

	localIP := GetLocalIP()
	meta := fmt.Sprintf("%s@%s", senderName, localIP)
	fullMessage := fmt.Sprintf("%s: %s", meta, message)

	encrypted, err := Encrypt(fullMessage, key)
	if err != nil {
		return fmt.Errorf("encryption error: %w", err)
	}

	_, err = conn.Write(encrypted)
	if err != nil {
		return fmt.Errorf("send error: %w", err)
	}
	return nil
}

func InputLoop(key []byte, name string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\nTarget IP:Port → ")
		addr, _ := reader.ReadString('\n')
		addr = strings.TrimSpace(addr)
		parts := strings.Split(addr, ":")
		if len(parts) != 2 {
			fmt.Println("Invalid address format. Use IP:Port")
			continue
		}

		fmt.Print("Message → ")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)

		err := SendMessage(parts[0], parts[1], msg, key, name)
		if err != nil {
			fmt.Println("Error sending message:", err)
		}
	}
}

func GetLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "unknown"
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}
