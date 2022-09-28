package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"goapi/massive_chat_rooms/common/message"
	"io"
	"net"
)

//处理与客户端的通讯
func process(coon net.Conn) {
	defer coon.Close()
	for {
		pak, err := readPak(coon)
		if err != nil {
			if err == io.EOF {
				fmt.Println("client and serve exit")
				return
			} else {
				fmt.Println("readPak(coon) error = ", err)
			}
		}
		err = serveProcessMessage(coon, &pak)
		if err != nil {
			fmt.Println("serveProcessMessage(coon,&pak) error=", err)
			return
		}
		fmt.Println("readPak message = ", pak)
	}

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

//
// readPak
//  #Summary: read client data
//  #Description:
//
func readPak(coon net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 1024)
	_, err = coon.Read(buf[:4])
	if err != nil {
		fmt.Println("read err:", err)
		return
	}
	fmt.Println("Read the message sent by th client---")
	var readPak uint64
	readPak = binary.BigEndian.Uint64(buf[0:4])
	n, err := coon.Read(buf[:readPak])
	if n != int(readPak) || err != nil {
		fmt.Println("read err:", err)
		return
	}
	err = json.Unmarshal(buf[:readPak], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal(buf[:readPak], &mes) error = ", err)
		return
	}
	fmt.Println("The data read is:", buf[:readPak])
	return
}
func serveProcessMessage(coon net.Conn, mes *message.Message) (err error) {
	switch mes.MessageType {
	case message.LoginMesType:
		err = serveLogin(coon, mes)
		if err != nil {
			return err
		}
	case message.RegisterMesType:
	default:
		fmt.Println("message type not existence")
	}
	return err
}
func serveLogin(coon net.Conn, mes *message.Message) (err error) {
	var (
		loginMes    message.Login
		resMes      message.Message
		loginResMes message.ResMessage
	)
	err = json.Unmarshal([]byte(mes.MessageData), &loginMes)
	resMes.MessageType = message.LoginResMesType
	if err != nil {
		fmt.Println("json.Unmarshal() fail err = ", err)
		return
	}
	if loginMes.UserId == 100 && loginMes.UserPassword == "123456" {
		loginResMes.ResCode = 200
	} else {
		loginResMes.ResCode = 500
		loginResMes.MessageData = "user not exit!"
	}
	_, err = json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal() fail", err)
		return
	}
	return
}
