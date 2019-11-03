package main

import (
	"fmt"
	"time"

	"github.com/krsanky/go-urt-server-query/urt"
)

func main() {
	urtRfa()
}

func urtRfa() {
	data, err := urt.GetRawStatus("74.91.112.64:27960") // RFA
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(data))

	vars, err := urt.ServerVars(data)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("len vars:%d\n", len(vars))

	players, err := urt.Players(data)
	if err != nil {
		panic(err)
	}

	now := time.Now()
	fmt.Printf("%s %s %d\n", now.Format(time.RFC3339), vars["mapname"], len(players))
}
