package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	flag.Parse()
	port := flag.Arg(0)
	if port == "" {
		fmt.Println("Usage:\n\tsocketserver <port number>\n\t- <port number> is the listening port")
		return
	}

	addy := ":" + port
	l, err := net.Listen("tcp4", addy)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	fmt.Println("Listening on port ", port)

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}
