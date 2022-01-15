package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("=== TCP SCANNER START ===")

	for i := 1; i <= 1024; i++ { //create loop of ports to scan
		go func(j int) { //crea una nueva rutina por cada iteración del bucle
			address := fmt.Sprintf("scanme.nmap.org:%d", j) // creas la direccion
			conn, err := net.Dial("tcp", address)           //conectas a la direccion
			if err != nil {                                 //si hay error pasas al siguiente valor
				return
			}
			conn.Close()               //cierras la conexion
			fmt.Printf("%d open\n", j) //imprimes el valor del puerto
		}(i) //pasas el valor a la funcion de cada hilo
	}
}
