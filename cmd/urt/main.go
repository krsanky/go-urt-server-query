package main

import (
	"fmt"

	"github.com/krsanky/go-urt-server-query/urt"
)

func main() {
	//urt.GetStatus("216.52.148.134:27961") // urtctf
	servers, err := urt.GetServers()
	if err != nil {
		panic(err)
	}
	for _, s := range servers {
		fmt.Println(s, s.Address())
	}
	fmt.Println()
}
