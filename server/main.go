package main

import (
	"fmt"
	"io"
	"log"
	"matinvpn/auth"
	"matinvpn/crypto"
	"matinvpn/obfuscation"
	"net"
)

var secretKey = []byte("12345678901234567890123456789012") // 32 bytes AES

func handleConnection(conn net.Conn) {
	defer conn.Close()

	err := obfuscation.FakeHTTP2Handshake(conn)
	if err != nil {
		log.Println("Obfuscation failed:", err)
		return
	}

	crypt := crypto.NewAESCrypto(secretKey)

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("Error reading:", err)
		return
	}

	decrypted, err := crypt.Decrypt(buf[:n])
	if err != nil {
		log.Println("Decrypt error:", err)
		return
	}

	var username, password string
	fmt.Sscanf(string(decrypted), "%s|%s", &username, &password)

	if err := auth.Authenticate(username, password); err != nil {
		io.WriteString(conn, "AUTH_FAILED")
		return
	}

	io.WriteString(conn, "AUTH_OK")
	log.Println("User authenticated:", username)
}

func main() {
	listener, err := net.Listen("tcp", ":443")
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	defer listener.Close()

	fmt.Println("MatinVpn server listening on port 443...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
