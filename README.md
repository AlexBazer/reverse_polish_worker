##Reverse polish notation calculator worker

Calculator runs on ZeroMQ server and listens to tcp://localhost:5555 socket


### Build
ZeroMQ binaries should be installed into the system.
For **Ubuntu** ```sudo apt-get install libzmq-dev```

Golang **gozmq** lib should be installed

```
go get github.com/alecthomas/gozmq
go install github.com/alecthomas/gozmq
```
