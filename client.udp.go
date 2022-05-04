package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 60000,
	})
	if err != nil {
		fmt.Println("Failed to connect the server，err:", err)
		return
	}
	defer socket.Close()
	var reply [1024]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Message: ")
		msg, _ := reader.ReadString('\n')
		socket.Write([]byte(msg))
		n, addr, err := socket.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("read reply msg failed, err: ", err)
			return
		}
		fmt.Println(addr, ": ", string(reply[:n]))
	}
}
