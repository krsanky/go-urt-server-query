package urt

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"
)

var prefix = []byte{0xff, 0xff, 0xff, 0xff}

type Server struct {
	Ip   net.IP
	Port int
}

func (s Server) String() string {
	return fmt.Sprintf("<Server ip:%s port:%d>", s.Ip, s.Port)
}

func (s Server) Address() string {
	return fmt.Sprintf("%s:%d", s.Ip, s.Port)
}

type Player struct {
	Name  string
	Ping  int
	Score int
}

func (p Player) String() string {
	return fmt.Sprintf("<Player %s ping:%d score:%d>", p.Name, p.Ping, p.Score)
}

func GetRawStatus(address string) ([]byte, error) {
	return Get(address, "getstatus")
}
func Get(address string, msg string) ([]byte, error) {
	conn, err := net.Dial("udp", address)
	if err != nil {
		return nil, err
	}
	//err = conn.SetDeadline(time.Now().Add(5 * time.Second))
	//if err != nil {
	//	return nil, err
	//}
	defer conn.Close()

	fmt.Fprintf(conn, string(append(prefix, msg...)))

	buf := make([]byte, 65507)
	n, err := bufio.NewReader(conn).Read(buf)
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
}

func ServerVars(data []byte) (map[string]string, error) {
	split := bytes.Split(data, []byte("\n"))
	if len(split) < 2 {
		return nil, errors.New("problem with server status data")
	}
	split = bytes.Split(split[1], []byte("\\"))
	split = split[1:]
	if len(split)%2 != 0 {
		return nil, errors.New("problem with server status data 2")
	}
	var vars = make(map[string]string)
	for i := 0; i < len(split); i += 2 {
		vars[string(split[i])] = string(split[i+1])
	}
	return vars, nil
}

func getServersData() ([][]byte, error) {
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

func GetServers() ([]Server, error) {
	resp, err := getServersData()
	if err != nil {
		return nil, err
	}

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

	return servers, nil
}

func Players(data []byte) ([]Player, error) {
	split := bytes.Split(data, []byte("\n"))
	if len(split) < 2 {
		return nil, errors.New("problem with server status data")
	}
	split = split[2:]

	var players []Player
	for _, p := range split {
		if bytes.Equal([]byte(""), p) {
			continue
		}
		split = bytes.SplitN(p, []byte(" "), 3)
		if len(split) != 3 {
			continue
		}
		ping, err := strconv.Atoi(string(split[1]))
		if err != nil {
			continue
		}
		score, err := strconv.Atoi(string(split[0]))
		if err != nil {
			continue
		}
		players = append(players, Player{string(split[2]), ping, score})

	}
	return players, nil
}
