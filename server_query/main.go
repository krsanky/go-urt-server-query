package server_query

import (
	"bufio"
	"fmt"
	"net"
)

var magic = []byte{0xff, 0xff, 0xff, 0xff}

func GetStatus(hostport string) {
	conn, err := net.Dial("udp", hostport)
	if err != nil {
		panic(fmt.Sprintf("err:%v\n", err))
	}

	m := append(magic, []byte("getstatus")...)
	fmt.Fprintf(conn, string(m))

	buf := make([]byte, 65507)
	len, err := bufio.NewReader(conn).Read(buf)
	if err != nil {
		panic(fmt.Sprintf("err2:%v\n", err))
	}
	fmt.Printf("len:%d\n", len)
	fmt.Printf("buf:%s\n", string(buf))

}
