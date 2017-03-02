package main

import (
	"fmt"

	zmq "github.com/alecthomas/gozmq"
)

func main() {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.REP)
	defer context.Close()
	defer socket.Close()
	socket.Bind("tcp://*:5555")

	// Wait for messages
	for {
		msg, _ := socket.Recv(0)
		println("Received ", string(msg))

		res, _ := calculate(string(msg))

		// send reply back to client
		reply := fmt.Sprintf("%f", res)
		socket.Send([]byte(reply), 0)
	}
}
