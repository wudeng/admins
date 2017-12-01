package data

import (
	"basic/ssdb/gossdb"
	"testing"
	"fmt"
)

func Test_normal_record(t *testing.T) {
	gossdb.Connect("192.168.50.10", 8888, 1)
	//InitServers()
	var s Server
	s.Id = 1
	s.Get()
	fmt.Println(s)
}
