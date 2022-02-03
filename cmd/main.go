package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

// handle scans input from "telnet localhost 8080"
// and prints it to the screen until timeout(100 seconds)
func handle(conn net.Conn) {
	defer conn.Close()

	// setting deadline 100 seconds
	err := conn.SetDeadline(time.Now().Add(100 * time.Second))
	if err != nil {
		log.Fatalln("CONN TIMEOUT")
	}

	scanner := bufio.NewScanner(conn)

	// scans input as text, applies fizzbuzz logic and writes into connection
	for scanner.Scan() {
		ln := scanner.Text()

		num, _ := strconv.Atoi(ln)
		if num%3 == 0 && num%5 == 0 {
			fmt.Fprintln(conn, "FizzBuzz")
			continue
		}

		if num%3 == 0 {
			fmt.Fprintln(conn, "Fizz")
			continue
		}
		if num%5 == 0 {
			fmt.Fprintln(conn, "Buzz")
			continue
		}
	}
}
