package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	argLength := len(os.Args[1:])
	fmt.Printf("Arg length is %d", argLength)
	fmt.Printf("other args: %+v\n", flag.Args())
	service := os.Args[1]
	fmt.Printf("Args are", service)
	//parseIp, parseIpPort, err := net.SplitHostPort(service)
	//fmt.Printf("ParseIp, parseIpPort, err are", parseIp, parseIpPort, err)
	//joinedPort := parseIp + `:` + parseIpPort
	//fmt.Printf("Joined port is", joinedPort)
	//ipAddr, err := net.ResolveIPAddr("tcp4", joinedPort)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	//ipAddr, err := net.ResolveIPAddr("ip", service)
	fmt.Printf("Formatted address is", tcpAddr)
	//fmt.Printf("Formatted address is", ipAddr)
	checkError(err)
	//conn, err := net.Dial("tcp", service)
	//conn, err := net.DialIP("ip", nil, ipAddr)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	//conn, err := net.DialIP("tcp", nil, tcpAddr)
	checkError(err)
	//_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	checkError(err)
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
