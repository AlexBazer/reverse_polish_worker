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
	fmt.Println("Start reverse polish worker")
	fmt.Println("...to stop worker hit ctrl-c")
	var reply string

	// Wait for messages
	for {
		msg, _ := socket.Recv(0)
		println("Received ", string(msg))

		res, err := calculate(string(msg))
		if err != nil {
			reply = fmt.Sprintf("ERROR: %v", err)
		} else {
			reply = fmt.Sprintf("%f", res)
		}

		socket.Send([]byte(reply), 0)
		fmt.Println("Sent ", reply)
	}
}
