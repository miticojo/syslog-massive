package main

import (
	"fmt"
	"log/syslog"
	"net"
	"os"
	"strconv"
	"time"
)

type Priority syslog.Priority

type Formatter func(p Priority, hostname, tag, content string) string
func DefaultFormatter(p Priority, hostname, tag, content string) string {
	timestamp := time.Now().Format(time.RFC3339)
	msg := fmt.Sprintf("<%d> %s %s %s[%d]: %s\n",
		p, timestamp, hostname, tag, os.Getpid(), content)
	return msg
}
func RFC3164Formatter(p Priority, hostname, tag, content string) string {
	timestamp := time.Now().Format(time.Stamp)
	msg := fmt.Sprintf("<%d> %s %s %s[%d]: %s\n",
		p, timestamp, hostname, tag, os.Getpid(), content)
	return msg
}

func sendUDPMsg(proto string, remote string, hostname string){
	conn, _ := net.Dial(proto, remote)
	conn.Write([]byte(DefaultFormatter(Priority(28), hostname, "syslog-client", "hello from client" )))
	conn.Close()
}

func main() {

	args := os.Args[1:]

	if len(args) != 3 {
		fmt.Println("Missing parameters: <nr-of-messages> <udp|tcp> <host>:<port>")
		os.Exit(1)
	}

	msgQty, _ := strconv.Atoi(args[0])
	syslogProto := args[1]
	syslogSrv := args[2]
	hostname, _ := os.Hostname()

	// cycling based on message quantity
	for i := 0; i < msgQty; i++ {
		sendUDPMsg(syslogProto, syslogSrv, hostname)
		time.Sleep(time.Nanosecond*200000)
	}

	// print out comment
	fmt.Printf("%d messages sent to %s through %s proto\n", msgQty, syslogSrv, syslogProto)
}