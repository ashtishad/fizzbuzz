package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"regexp"
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
// and prints it to the screen until timeout(120 seconds)
func handle(conn net.Conn) {
	defer conn.Close()

	// setting deadline 120 seconds
	err := conn.SetDeadline(time.Now().Add(120 * time.Second))
	if err != nil {
		log.Fatalln("CONN TIMEOUT")
	}

	scanner := bufio.NewScanner(conn)

	// scans input as text, applies fizzbuzz logic and writes into connection
	for scanner.Scan() {
		sc := scanner.Text()

		// add validate function
		num, err := validate(sc)
		if err != nil {
			fmt.Fprintln(conn, err.Error())
			continue
		}

		// add fizzbuzz logic
		res, err := getFizzBuzz(num)
		if err != nil {
			fmt.Fprintln(conn, err.Error())
			continue
		}
		fmt.Fprintln(conn, res)
	}
}

// validate returns a number if input is a number, used regexp for validation
// if not, returns error
func validate(s string) (int, error) {
	// validate input number
	if s == "" {
		return 0, errors.New("empty string")
	}
	rgx := regexp.MustCompile(`[0-9]+`)
	if rgx.MatchString(s) == false {
		return 0, errors.New("input must be a number")
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("entered mixture of numbers and letters")
	}
	return num, nil
}

// getFizzBuzz returns a string with fizzbuzz logic
// number is divisible by 3, returns "Fizz" , divisible by 5, returns "Buzz"
// divisible by 3 and 5, returns "FizzBuzz"
// not divisible by 3 or 5, returns "not fizzbuzz"
func getFizzBuzz(num int) (string, error) {

	if num%3 == 0 && num%5 == 0 {
		fmt.Println("FizzBuzz")
		return "FizzBuzz", nil
	}

	if num%3 == 0 {
		fmt.Println("Fizz")
		return "Fizz", nil
	}
	if num%5 == 0 {
		fmt.Println("Buzz")
		return "Buzz", nil
	}

	return "", errors.New("not fizzbuzz")
}
