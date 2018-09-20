package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"strconv"
	"time"
)

func main() {

	args := os.Args[1:]

	if len(args) != 3 {
		fmt.Println("Missing parameters: <nr-of-messages> <udp|tcp> <host>:<port>")
		os.Exit(1)
	}

	msgQty, _ := strconv.Atoi(args[0])
	syslogProto := args[1]
	syslogSrv := args[2]

	// setup syslog endpoint
	sysLog, err := syslog.Dial(syslogProto, syslogSrv,
		syslog.LOG_WARNING|syslog.LOG_DAEMON, "demotag")

	if err != nil {
		log.Fatal(err)
	}

	// cycling based on message quantity
	for i := 0; i < msgQty; i++ {
		fmt.Fprintf(sysLog, "This is a daemon warning with demotag.")
		time.Sleep(time.Nanosecond*200000)
	}

	// print out comment
	fmt.Printf("%d messages sent to %s through %s proto\n", msgQty, syslogSrv, syslogProto)

}