package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Missing parameters: <udp|tcp> <port>")
		os.Exit(1)
	}

	syslogPort := args[1]
	syslogProto := args[0]


	if syslogProto == "udp" {
		l, err := net.ListenPacket(syslogProto, "0.0.0.0" + ":" + syslogPort)

		if err != nil {
			fmt.Println("Error listening:", err.Error())
			os.Exit(1)
		}
		// Close the listener when the application closes.
		defer l.Close()
		fmt.Printf("Listening for messages on 0.0.0.0 port %s/%s:\n", syslogProto, syslogPort)
		readUDP(l)
	} else {
		l, err := net.Listen("tcp4", "0.0.0.0" + ":" + syslogPort)
		if err != nil {
			fmt.Println("Error listening:", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Listening for messages on 0.0.0.0 port %s/%s:\n", syslogProto, syslogPort)
		conn, _ := l.Accept()
		defer l.Close()
		readTCP(conn)
	}
}

func readTCP(conn net.Conn){
	count := 0
	reader :=  bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		//buffer := make([]byte, 2048)
		//conn.Read(buffer)
		count++
		switch err {
		case nil:
			break
		case io.EOF:
		default:
			fmt.Println("ERROR", err)
		}
		fmt.Printf("%d\t%s", count, message)
	}
}

func readUDP(conn net.PacketConn){
	count := 0

	for {
		buffer := make([]byte, 2048)
		conn.ReadFrom(buffer)
		count++
		fmt.Printf("%d\t%s", count, string(buffer))
	}
}
