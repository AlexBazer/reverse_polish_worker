package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq2"
)

func main() {
	//  Socket to talk to clients
	responder, _ := zmq.NewSocket(zmq.REP)
	defer responder.Close()
	responder.Bind("tcp://*:5555")
	fmt.Println("Start reverse polish worker")
	fmt.Println("...to stop worker hit ctrl-c")

	for {
		//  Wait for next request from client
		msg, _ := responder.Recv(0)
		fmt.Println("Received ", msg)

		//  Make calculations
		res, _ := calculate(string(msg))

		//  Send reply back to client
		reply := fmt.Sprintf("%f", res)
		responder.Send(reply, 0)
		fmt.Println("Sent ", reply)
	}
}
