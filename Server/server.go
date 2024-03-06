package main

import (
	"fmt"
	"net"

	utils "github.com/HackJack14/SteamSync/Utils"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println(utils.ReadString(conn))
		utils.ReadFile(conn)
	}
}
