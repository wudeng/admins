package data

import (
	"basic/ssdb/gossdb"
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"
	"net/url"
	"github.com/golang/glog"
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
	defer res.Body.Close()
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

func (s *Server) Exist() bool {
	boolean, _ := gossdb.C().Hexists(KEY_SERVER_LIST, strconv.FormatInt(s.Id, 10))
	return boolean
}

func (s *Server) Save() error {
	err := gossdb.C().Hset(KEY_SERVER_LIST, strconv.FormatInt(s.Id, 10), s)
	return err
}

func (s *Server) Del() error {
	err := gossdb.C().Hdel(KEY_SERVER_LIST, strconv.FormatInt(s.Id, 10))
	return err
}

func (s *Server) Mail(mailId string) error {
	str := fmt.Sprintf("http://%s:%d/send_mail", s.Ip, s.Pt+10000)
	res, err := http.PostForm(str, url.Values{"mail_id":{mailId}})
	if nil != err {
		glog.Error(err)
		return err
	}
	defer res.Body.Close()
	var ret string
	json.NewDecoder(res.Body).Decode(&ret)
	return nil
}

func (s *Server) Online() int {
	url := fmt.Sprintf("http://%s:%d/online", s.Ip, s.Pt+10000)
	res, _ := http.Get(url)
	defer res.Body.Close()
	var online int
	json.NewDecoder(res.Body).Decode(&online)
	return online
}

func ListServer() []*Server {
	list := make([]*Server, 0)
	value, err := gossdb.C().Hscan(KEY_SERVER_LIST, "", "", 50)
	if err == nil {
		for _, v := range value {
			data := &Server{}
			if err := v.As(data); err == nil {
				list = append(list, data)
			}
		}
	}
	return list
}
