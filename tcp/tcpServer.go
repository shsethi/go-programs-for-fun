package main

import (
	"bufio"
	"log"
	"net"
	"strings"
	"time"
)

// Connect to this using `telnet 0.0.0.0 1234`

func main() {

	PORT:= ":1234"

	l, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Println(err)
		return
	}
	defer l.Close()


	connection, err := l.Accept()
	if err != nil {
		log.Println(err)
		return
	}


	for {
		netData, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		if strings.TrimSpace(string(netData)) == "STOP" {
			log.Println("Exiting TCP server!")
			return
		}

		log.Print("got message -> ", string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		connection.Write([]byte(myTime))
	}

}
