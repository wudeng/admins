package data

import (
	"basic/ssdb/gossdb"
	"encoding/json"
	"net/http"
	"strconv"
)

type Response struct {
	Sl []Server      `json:"sl"`
	My []interface{} `json:"my"`
}

type Server struct {
	Id int64 `json:"id"`
	Nm string `json:"nm"`
	Ip string `json:"ip"`
	Pt uint16 `json:"pt"`
	St uint8  `json:"st"`
	Re uint8  `json:"re"`
}

func InitServers() {
	url := "http://192.168.4.168:8080/server_list"
	res, _ := http.Get(url)
	var server_list Response
	json.NewDecoder(res.Body).Decode(&server_list)
	for _, v := range server_list.Sl {
		v.Save()
	}
}

func (s *Server) Get() error {
	value, err := gossdb.C().Hget(KEY_SERVER_LIST, strconv.FormatInt(s.Id, 10))
	if err == nil {
		err = value.As(s)
	}
	return err
}

func (s *Server) Save() error {
	err := gossdb.C().Hset(KEY_SERVER_LIST, strconv.FormatInt(s.Id, 10), s)
	return err
}
