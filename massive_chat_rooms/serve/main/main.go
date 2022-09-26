package main

import (
	"fmt"
	"net"
)

//处理与客户端的通讯
func process(coon net.Conn) {
	defer coon.Close()
	buf := make([]byte, 1024)
	read, err := coon.Read(buf[:4])
	fmt.Println("Read the message sent by th client---")
	if read != 4 || err != nil {
		fmt.Println("read err:", err)
		return
	}
	fmt.Println("The data read is:", buf[:4])
}
func main() {
	fmt.Println("Serve listened in 8888")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen() error=", err)
		return
	}
	for {
		coon, err := listen.Accept()
		if err != nil {
			fmt.Println("conn.Accept() error=", err)
			return
		}
		go process(coon)
	}
}
