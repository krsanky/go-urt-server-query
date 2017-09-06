package server_query

import (
	"bufio"
	"fmt"
	"net"
)

var magic = []byte{0xff, 0xff, 0xff, 0xff}
var msg = []byte("getstatus")

func Test1() {
	fmt.Println("Test1...")
	conn, err := net.Dial("udp", "216.52.148.134:27961")
	if err != nil {
		panic(fmt.Sprintf("err:%v\n", err))
	}
	m := append(magic, msg...)
	fmt.Fprintf(conn, string(m))
	res, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		panic(fmt.Sprintf("err2:%v\n", err))
	}
	fmt.Printf("res:%v\n", res)

}

//(send-msg socket "getstatus" host port)
//(let [buffer (byte-array 65507)
//(get-status "216.52.148.134" 27961))

//(def magic (byte-array (map unchecked-byte (repeat 4 0xff))))
//(defn send-msg [socket msg host port]
//	(let [msg' (->> msg
//				 (.getBytes)
//				 (concat magic)
//				 (byte-array))
//          host' (InetAddress/getByName host)
//          packet (new DatagramPacket msg' (count msg') host' port)]
//		(.send socket packet)))
