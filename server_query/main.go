package server_query

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"net"
	"time"
)

var prefix = []byte{0xff, 0xff, 0xff, 0xff}

type Server struct {
	Address net.IP
	Port    int
}

func (s Server) String() string {
	return fmt.Sprintf("<Server ip:%s port:%d>", s.Address, s.Port)
}

func GetStatus(address string) ([]byte, error) {
	return Get(address, "getstatus")
}
func Get(address string, msg string) ([]byte, error) {
	conn, err := net.Dial("udp", address)
	if err != nil {
		return nil, err
	}
	err = conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	m := append(prefix, msg...)
	fmt.Fprintf(conn, string(m))

	buf := make([]byte, 65507)
	_, err = bufio.NewReader(conn).Read(buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func GetServersData() ([][]byte, error) {
	addrs, err := net.LookupHost("master.urbanterror.info")
	if err != nil {
		return nil, err
	}
	if len(addrs) < 1 {
		return nil, errors.New("error with master server resolution")
	}
	fmt.Printf("master server:%s\n", addrs[0])

	conn, err := net.Dial("udp", fmt.Sprintf("%s:27900", addrs[0]))
	if err != nil {
		return nil, err
	}
	err = conn.SetDeadline(time.Now().Add(3 * time.Second))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	conn.Write(append(prefix, "getservers 68 empty full"...))

	var bufs [][]byte
	var n int
	for {
		buf := make([]byte, 65507)
		n, err = conn.Read(buf)
		if e, ok := err.(net.Error); ok && e.Timeout() {
			break
		} else if err != nil {
			return nil, err
		}
		bufs = append(bufs, buf[:n])
	}

	return bufs, nil
}

// borken
func responseValid(resp [][]byte) bool {
	fmt.Printf("%v == %v\n", resp[0], append(prefix, "getserversResponse"...))
	if bytes.Equal(resp[0], append(prefix, "getserversResponse"...)) {
		return true
	}
	return false
}

func GetServers(resp [][]byte) []Server {
	var servers []Server
	var data [][]byte

	for i, d := range resp {
		data = bytes.Split(d, []byte("\\"))
		data = data[1:]
		if i == len(resp)-1 {
			data = data[:len(data)-1]
		}
		for _, s := range data {
			if len(s) != 6 {
				fmt.Printf("%q %d ", s, len(s))
				continue
			}
			ip := net.IPv4(s[0], s[1], s[2], s[3])
			port := int(s[4])<<8 | int(s[5])
			servers = append(servers, Server{ip, port})
		}
	}

	return servers
}
