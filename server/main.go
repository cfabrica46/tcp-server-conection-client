package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	log.SetFlags(log.Llongfile)
	if len(os.Args) >= 3 {
		serverAddress := fmt.Sprintf("%s:%s", os.Args[1], os.Args[2])

		l, err := net.Listen("tcp", serverAddress)
		if err != nil {
			log.Fatal(err)
		}
		defer l.Close()

		fmt.Println("Listening on " + serverAddress)
		fmt.Println()

		for {
			fmt.Println("esperando una conexion...")

			conn, err := l.Accept()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("conexion recibida/aceptada")

			fmt.Println("leyendo...")

			reader := bufio.NewReader(conn)

			for {

				fmt.Println("nuevo ciclo")

				content, err := reader.ReadString('\n')

				if err != nil {
					if strings.Contains(err.Error(), "host") {
						break
					} else {
						log.Fatal(err)
					}
				}
				fmt.Println("se a leido!")
				fmt.Println()

				respuesta := fmt.Sprintf("Mensaje Recivido: %s\n", content)

				fmt.Println(respuesta)

				_, err = conn.Write([]byte(respuesta))

				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("fin")
			}
		}

	}
}
