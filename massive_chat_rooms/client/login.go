package client

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"goapi/massive_chat_rooms/common/message"
	"net"
)

func Login(userId int, userPwd string) (err error) {
	fmt.Println("Login connection ready")
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("net.Dial() error=", err)
		return err
	}
	defer conn.Close()
	var (
		mes   message.Message
		login message.Login
	)
	mes.MessageType = message.LoginMesType
	login.UserId = userId
	login.UserPassword = userPwd
	data, err := json.Marshal(login)
	if err != nil {
		fmt.Println("json.Marshal() error = ", err)
		return err
	}
	mes.MessageData = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(data) error = ", err)
		return err
	}
	var (
		pkgLen uint64
		buf    [4]byte
	)
	pkgLen = uint64(len(data))
	binary.BigEndian.PutUint64(buf[0:4], pkgLen)
	n, err := conn.Write(data)
	//serve request data
	//time.Sleep(10 * time.Second)

	pkg, err := readPkg(conn)
	if err != nil {
		fmt.Println("readPkg(conn) error = ", err)
		return
	}
	//request loginMessage
	var loginMes message.ResMessage
	err = json.Unmarshal([]byte(pkg.MessageData), &loginMes)
	if err != nil {
		return
	}
	if loginMes.ResCode == 200 {
		fmt.Println("login success!")
	} else if loginMes.ResCode == 500 {
		fmt.Println(loginMes.ResCode)
	}
	if n != 4 || err != nil {
		fmt.Println("conn.Write(data) error = ", err)
		return
	}
	fmt.Printf("client send message length = %d , messge = %s", len(data), string(data))
	return
}
