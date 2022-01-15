package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("=== TCP SCANNER START ===")
	/* _, err := net.Dial("tcp", "scanme.nmap.org:80")*/
	/*if err == nil {*/
	/*fmt.Println("Connection successful")*/
	/*}*/

	for i := 1; i <= 1024; i++ {
		go func(j int) {
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
}
