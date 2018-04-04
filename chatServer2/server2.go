package main

import (
	"flag"
	"fmt"
	"net"
	"syscall"
)

const maxRead = 25

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		panic("usage: host port")
	}
	hostandPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	listener := initServer(hostandPort)
	for {
		conn, err := listener.Accept()
		checkError(err, "Accept: ")
		go connectionHandler(conn)
	}
}
func initServer(hostandPort string) net.Listener {
	serverAddr, err := net.ResolveTCPAddr("tcp", hostandPort)
	checkError(err, "Resolving address:port failed: '"+hostandPort+"'")
	listener, err := net.ListenTCP("tcp", serverAddr)
	checkError(err, "ListenTCP: ")
	println("Listening to: ", listener.Addr().String())
	return listener
}
func connectionHandler(conn net.Conn) {
	connFrom := conn.RemoteAddr().String()
	println("Connection from: ", connFrom)
	sayHello(conn)
	for {
		var ibuf = make([]byte, maxRead+1)
		length, err := conn.Read(ibuf[0:maxRead])
		ibuf[maxRead] = 0
		switch err {
		case nil:
			handleMsg(length, err, ibuf)
		case syscall.EAGAIN:
			continue
		default:
			goto DISCONNECT
		}
	}
DISCONNECT:
	err := conn.Close()
	println("Closed connection: ", connFrom)
	checkError(err, "Close: ")
}
func sayHello(to net.Conn) {
	obuf := []byte{'L', 'e', 't', '\'', 's', ' ', 'G', 'O', '!', '\n'}
	wrote, err := to.Write(obuf)
	checkError(err, "Write: wrote "+string(wrote)+" bytes.")
}
func handleMsg(length int, err error, msg []byte) {
	if length > 0 {
		var misLength int
		if length > maxRead {
			misLength = maxRead
		} else {
			misLength = length
		}
		fmt.Println(string(msg[:misLength]))
	}
}
func checkError(error error, info string) {
	if error != nil {
		panic("ERROR: " + info + " " + error.Error())
	}
}