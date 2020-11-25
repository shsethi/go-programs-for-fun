package main

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {

	PORT := "0.0.0.0:1235"

	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println("Failed to connect "+err.Error())
		return
	}

	defer connection.Close()
	buffer := make([]byte, 3)
	rand.Seed(time.Now().Unix())

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		reqData := string(buffer[0:n-1])
		fmt.Printf("-> sever recieved data: %s\n", string(reqData))


		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting UDP server!")
			return
		}

		_, err = connection.WriteToUDP( []byte("ok"), addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
