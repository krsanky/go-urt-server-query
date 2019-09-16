package main

import (
	"fmt"
	"os"

	"github.com/krsanky/go-urt-server-query/urt"
)

func main() {
	if len(os.Args) >= 2 {
		switch arg1 := os.Args[1]; arg1 {
		case "urtctf":
			urtCtf()
		default:
			usage()
		}
	} else {
		usage()
	}
}

func usage() {
	fmt.Println()
	fmt.Println(`urt [urtctf]`)
	fmt.Println()
}

func urtCtf() {
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
