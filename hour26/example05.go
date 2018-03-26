package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

type Message struct {
	conn    net.Conn
	message []byte
}

var connections5 []net.Conn
var messages = make(chan Message)
var addClient5 = make(chan net.Conn)
var removeClient5 = make(chan net.Conn)

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
		addClient5 <- conn

		go handleRequest(conn)

	}

}

func startChannels() {
	for {
		select {

		case message := <-messages:
			broadcastMessage(&message)
		case newClient := <-addClient5:
			connections5 = append(connections5, newClient)
			fmt.Println(len(connections5))
		case deadClient := <-removeClient5:
			removeConn(deadClient)
			fmt.Println(len(connections5))
		}
	}
}

func handleRequest(conn net.Conn) {
	for {

		message := make([]byte, 1024)

		_, err := conn.Read(message)
		if err != nil {
			if err == io.EOF {
				removeClient5 <- conn
				conn.Close()
				return
			}
			log.Fatal(err)
		}

		m := Message{
			conn:    conn,
			message: message,
		}

		messages <- m
	}

}

func broadcastMessage(m *Message) {
	for _, conn := range connections5 {
		if m.conn != conn {
			_, err := conn.Write(m.message)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func removeConn(conn net.Conn) {
	var i int
	for i = range connections5 {
		if connections5[i] == conn {
			break
		}
	}
	connections5 = append(connections5[:i], connections5[i+1:]...)
}
