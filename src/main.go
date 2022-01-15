package main

import (
	"fmt"
	"net"
	"sort"
)

//#######################
//## Valores a cambiar ##
//#######################
// Dirección: Strin antes del %d variable address
// Numero Puertos: Variable puertos

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	/* puertos := 1024*/

	/*var wg sync.WaitGroup //añade un objeto de espera para que de tiempo a completar las conexiones*/
	/*fmt.Println("=== TCP SCANNER START ===")*/
	/*for i := 1; i <= puertos; i++ { //create loop of ports to scan*/
	/*wg.Add(1)*/

	/*go func(j int) { //crea una nueva rutina por cada iteración del bucle*/
	/*defer wg.Done()*/

	/*address := fmt.Sprintf("scanme.nmap.org:%d", j) // creas la direccion*/
	/*conn, err := net.Dial("tcp", address)           //conectas a la direccion*/
	/*if err != nil {                                 //si hay error pasas al siguiente valor*/
	/*return*/
	/*}*/

	/*conn.Close()               //cierras la conexion*/
	/*fmt.Printf("%d open\n", j) //imprimes el valor del puerto*/

	/*}(i) //pasas el valor a la funcion de cada hilo*/
	/*}*/
	/*wg.Wait()*/

	ports := make(chan int, 100)
	results := make(chan int)
	puertos := 1024
	var openports []int

	fmt.Println("===== TCP SCANNER START =====")

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= puertos; i++ {
			ports <- i
		}
	}()

	for i := 0; i < puertos; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(openports)

	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
