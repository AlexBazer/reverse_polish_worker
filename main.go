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
	var reply string
	for {
		//  Wait for next request from client
		msg, _ := responder.Recv(0)
		fmt.Println("Received ", msg)

		//  Make calculations
		res, err := calculate(string(msg))
		if err != nil {
			reply = fmt.Sprintf("ERROR: %v", err)
		} else {
			reply = fmt.Sprintf("%f", res)
		}
		//  Send reply back to client
		responder.Send(reply, 0)
		fmt.Println("Sent ", reply)
	}
}
