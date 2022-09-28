package client

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"goapi/massive_chat_rooms/common/message"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 1024)
	fmt.Println("read client send message")
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.Read() error = ", err)
		return message.Message{}, err
	}
	var pakLen uint64
	pakLen = binary.BigEndian.Uint64(buf[:4])
	n, err := conn.Read(buf[:pakLen])
	if n != int(pakLen) || err != nil {
		fmt.Println("conn.Read(buf[:pakLen]) error = ", err)
		return message.Message{}, err
	}
	err = json.Unmarshal(buf[:pakLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal(buf[:pakLen], &mes) error = ", err)
		return message.Message{}, err
	}
	return
}

func writePak(conn net.Conn, data []byte) (err error) {
	//send a length to serve
	var (
		pkgLen uint64
		buf    [4]byte
	)
	pkgLen = uint64(len(data))
	binary.BigEndian.PutUint64(buf[0:4], pkgLen)

	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf[:4]) error = ", err)
		return
	}

	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) error = ", err)
		return err
	}
	return
}
