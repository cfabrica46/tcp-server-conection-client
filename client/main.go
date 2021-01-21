package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	if len(os.Args) >= 3 {

		serverAddress := fmt.Sprintf("%s:%s", os.Args[1], os.Args[2])

		conn, err := net.Dial("tcp", serverAddress)

		if err != nil {
			log.Fatal(err)
		}

		r := bufio.NewReader(conn)

		for {
			message, err := reader.ReadString('\n')

			if err != nil {
				log.Fatal(err)
			}

			_, err = conn.Write([]byte(message))

			if err != nil {
				log.Fatal(err)
			}

			respuesta, err := r.ReadString('\n')

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(respuesta)

		}
	}

}
