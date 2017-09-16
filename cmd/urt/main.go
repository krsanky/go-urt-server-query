package main

import (
	"github.com/krsanky/go-urt-server-query/server_query"
)

func main() {
	//server_query.GetStatus("216.52.148.134:27961") // urtctf
	data, err := server_query.GetServersData()
	if err != nil {
		panic(err)
	}
	server_query.TransformServersData(data)
}
