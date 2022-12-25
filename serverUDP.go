package main

import (
	"fmt"
	"net"
)

func main() {
	// cria servidor na porta 8888
	serverAddr, err := net.ResolveUDPAddr("udp", ":8888")
	if err != nil {
		fmt.Println(err)
		return
	}

	// cria socket
	conn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// faz leitura das mensagens client
	for {
		// aloca um buffer e armazena mensagem
		buffer := make([]byte, 1024)

		// lê a mensagem
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		// mensagem exibida no client
		fmt.Println("Received message:", string(buffer[:n]), "from", addr)
	}
}
