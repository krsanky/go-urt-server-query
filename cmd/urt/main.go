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
	fmt.Println(string(data))

	vars, err := urt.ServerVars(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("len vars:%d\n", len(vars))

	players, err := urt.Players(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("len players:%d\n", len(players))
	for _, p := range players {
		fmt.Println(p)
	}
}
