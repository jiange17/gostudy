package main

import (
	"os"
	"fmt"
	"net"
	"bytes"
	"strings"
)

var (
	nameMap = make(map[string]bool)
)

func main() {
	fmt.Println("Starting the server...")
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}
		index := bytes.IndexAny(buf, "says")
		name := string(buf[:index])
		value := string(buf[index+6:len])
		if _, exist := nameMap[name]; !exist {
			nameMap[name] = true
		}
		value = strings.Trim(value, " ")
		fmt.Printf("value: %v, value==WHO? %v\n", value, value=="WHO")
		if value == "WHO" {
			for name := range nameMap {
				fmt.Println("name: ", name)
			}	
			fmt.Println("End")
		} else {
			if value == "SH" {
				os.Exit(0)
			} else {
				fmt.Printf("Received data: %v\n", value)
			}
		}
	}
}