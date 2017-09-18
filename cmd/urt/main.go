package main

import (
	"fmt"

	"github.com/krsanky/go-urt-server-query/urt"
)

func main() {
	data, err := urt.GetRawStatus("216.52.148.134:27961") // urtctf
	if err != nil {
		panic(err)
	}
	vars, err := urt.ServerVars(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("len vars:%d\n", len(vars))
	for k, v := range vars {
		fmt.Printf("k:%s v:%s\n", k, v)
	}

	//servers, err := urt.GetServers()
	//if err != nil {
	//	panic(err)
	//}
	//for _, s := range servers {
	//	fmt.Println(s, s.Address())
	//}
	//fmt.Println()
}
