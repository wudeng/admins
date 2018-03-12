package data

import (
	"testing"
	"fmt"
	"basic/ssdb/gossdb"
)

func Test_normal_record(t *testing.T) {
	gossdb.Connect("192.168.50.10", 8888, 1)
	//InitServers()
	var s Server
	s.Id = 2
	s.Get()
	ret := s.Mail("1")
	fmt.Println(ret)
}
