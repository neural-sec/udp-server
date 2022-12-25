package main

import (
	"fmt"
	"net"
)

func main() {
	// cria endereço
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8888")
	if err != nil {
		fmt.Println(err)
		return
	}

	// new socket udp
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// envia mensagem
	_, err = conn.Write([]byte("Hello, Server!"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// faz leitura da resposta
	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Received response:", string(buffer[:n]))
}
