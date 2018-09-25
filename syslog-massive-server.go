package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func readUDPMsg(connection *net.UDPConn, ch chan string) {
	buffer := make([]byte, 1024)
	n, _, err := connection.ReadFromUDP(buffer)
	if err != nil {
		panic(err)
	}
	ch <- string(buffer[:n])

}

func listenUDP(syslogPort int, syslogProto string) {
	addr := net.UDPAddr{
		Port: syslogPort,
		IP:   net.IP{0, 0, 0, 0},
	}
	connection, err := net.ListenUDP("udp", &addr)
	defer connection.Close()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Listening on %s:%d/%s\n", addr.IP, syslogPort, syslogProto)
	ch := make(chan string)
	cnt := 0

	for err == nil {
		go readUDPMsg(connection, ch)
		cnt++
		fmt.Printf("%d - %s", cnt, <-ch)
	}

}

func main() {

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Missing parameters: <udp|tcp> <port>")
		os.Exit(1)
	}

	syslogPort, _ := strconv.Atoi(args[1])
	syslogProto := args[0]

	if syslogProto == "udp" {
		listenUDP(syslogPort, syslogProto)
	}
}

