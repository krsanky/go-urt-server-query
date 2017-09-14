package server_query

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	bs, err := Get("216.52.148.134:27961", "getstatus")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(bs))
}

func TestGetStatus(t *testing.T) {
	bs, err := GetStatus("216.52.148.134:27961")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(bs))
}

func TestGetServers(t *testing.T) {
	bs, err := GetServers()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(bs))
}
