package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(c net.Conn) {
	remote := c.RemoteAddr().String()
	fmt.Printf("Serving %s\n", remote)
	defer c.Close()

	reader := bufio.NewReader(c)

	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			if data != "" {
				fmt.Print(remote, " - ", data, "\n")
				fmt.Print(remote, err, "\n")
			} else {
				fmt.Print(remote, " - ", err, "\n")
			}
			return
		}
		fmt.Print(remote, " - ", data)
	}
}
