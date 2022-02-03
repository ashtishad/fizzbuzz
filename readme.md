## FizzBuzz Tcp Server

#### What is it?

A simple tcp server that interacts with a lower level tcp client(telnet for instance).
It first validates user input and then prints out the input to the client in a fizzbuzz style.

* if number is divisible by 3, returns "Fizz" , 
* if number is divisible by 5, returns "Buzz" ,
* if number is divisible by 3 and 5, returns "FizzBuzz"
* if not divisible by 3 or 5, returns "not fizzbuzz"

#### How to use it?

* Run the server with the following command:
   `go run main.go`
* Connect to the server with telnet
   `telnet localhost:8080`
* Type in a number and press enter
* The server will respond with the fizzbuzz style response.

#### Few Gotchas

* This server can be used by multiple client, press control+] to exit from telnet client.
* Used regex for server side validation.
* Sample inputs: 15,10,6,fhjsdh, jh3545djb, 2784479234.
