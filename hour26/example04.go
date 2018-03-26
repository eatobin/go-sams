package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

var connections4 []net.Conn
var messages4 = make(chan []byte)
var addClient4 = make(chan net.Conn)
var removeClient4 = make(chan net.Conn)

func main() {

	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	go startChannels()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		addClient4 <- conn

		go handleRequest(conn)

	}

}

func startChannels() {
	for {
		select {

		case message := <-messages4:
			broadcastMessage(message)
		case newClient := <-addClient4:
			connections4 = append(connections4, newClient)
			fmt.Println(len(connections4))
		case deadClient := <-removeClient4:
			removeConn(deadClient)
			fmt.Println(len(connections4))
		}
	}
}

func handleRequest(conn net.Conn) {
	for {
		message := make([]byte, 1024)

		_, err := conn.Read(message)
		if err != nil {
			if err == io.EOF {
				removeClient4 <- conn
				conn.Close()
				return
			}
			log.Fatal(err)
		}

		messages4 <- message
	}
}

func broadcastMessage(m []byte) {
	for _, conn := range connections4 {
		_, err := conn.Write(m)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func removeConn(conn net.Conn) {
	var i int
	for i = range connections4 {
		if connections4[i] == conn {
			break
		}
	}
	connections4 = append(connections4[:i], connections4[i+1:]...)
}
