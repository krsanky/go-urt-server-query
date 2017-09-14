package server_query

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"time"
)

var magic = []byte{0xff, 0xff, 0xff, 0xff}

func Get(address string, msg string) ([]byte, error) {
	conn, err := net.Dial("udp", address)
	if err != nil {
		return nil, err
	}
	err = conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return nil, err
	}

	m := append(magic, msg...)
	fmt.Fprintf(conn, string(m))

	buf := make([]byte, 65507)
	len, err := bufio.NewReader(conn).Read(buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("len:%d\n", len)
	fmt.Printf("buf:%s\n", string(buf))

	return buf, nil
}

func GetStatus(address string) ([]byte, error) {
	return Get(address, "getstatus")
}

//Requesting servers from the master...
//Resolving master.urbanterror.info
//master.urbanterror.info resolved to 176.31.100.80:27900
// protocol 68
//function GetServers($master_server, $port, $protocol, $keywords = "empty full", $timeout = 1) {
//	if($master_server != "" && $port != 0 && $protocol != 0 && $socket = fsockopen('udp://'.$master_server, $port))
//		fwrite($socket, str_repeat(chr(255),4).'getservers '.$protocol.' '.$keywords.''."\n");
func GetServers() ([]byte, error) {
	fmt.Printf("GetServers...\n")
	addrs, err := net.LookupHost("master.urbanterror.info")
	if err != nil {
		return nil, err
	}

	if len(addrs) > 0 {
		fmt.Printf("master server:%s\n", addrs[0])
		return Get(fmt.Sprintf("%s:27900", addrs[0]), "getservers 68 empty full")
		//return nil, errors.New("err...")
	}
	return nil, errors.New("error with master server resolution")
}
