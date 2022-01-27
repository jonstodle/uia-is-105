package main

import (
	"log"
	"net"
)

func main() {
	strEcho := "Hallo"
	servAddr := "localhost:5001"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		log.Fatalf("ResolveTCPAddr failed: %v", err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}

	_, err = conn.Write([]byte(strEcho))
	if err != nil {
		log.Fatalf("Write to server failed: %v", err)
	}

	log.Printf("write to server = %v", strEcho)

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)
	if err != nil {
		log.Fatalf("Write to server failed: %v", err)
	}

	log.Printf("reply from server=%v", string(reply))

	conn.Close()
}
